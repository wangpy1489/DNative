package batchjob

import (
	"context"
	// "encoding/json"
	// berrors "errors"

	sparkv1beta2 "github.com/GoogleCloudPlatform/spark-on-k8s-operator/pkg/apis/sparkoperator.k8s.io/v1beta2"
	sparkclient "github.com/GoogleCloudPlatform/spark-on-k8s-operator/pkg/client/clientset/versioned"
	batchv1beta1 "github.com/wangpy1489/DNative/pkg/apis/batch/v1beta1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

var log = logf.Log.WithName("controller_batchjob")

/**
* USER ACTION REQUIRED: This is a scaffold file intended for the user to modify with their own Controller
* business logic.  Delete these comments after modifying this file.*
 */

// Add creates a new BatchJob Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	sClient, err := sparkclient.NewForConfig(mgr.GetConfig())
	if err != nil {
		log.Error(err, "failed to new sparkclient")
	}
	return &ReconcileBatchJob{client: mgr.GetClient(), scheme: mgr.GetScheme(), sparkClient: sClient}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("batchjob-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource BatchJob
	err = c.Watch(&source.Kind{Type: &batchv1beta1.BatchJob{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	// TODO(user): Modify this to be the types you create that are owned by the primary resource
	// Watch for changes to secondary resource Pods and requeue the owner BatchJob
	err = c.Watch(&source.Kind{Type: &corev1.Pod{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &batchv1beta1.BatchJob{},
	})
	if err != nil {
		return err
	}

	return nil
}

// blank assignment to verify that ReconcileBatchJob implements reconcile.Reconciler
var _ reconcile.Reconciler = &ReconcileBatchJob{}

// ReconcileBatchJob reconciles a BatchJob object
type ReconcileBatchJob struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client      client.Client
	scheme      *runtime.Scheme
	sparkClient *sparkclient.Clientset
}

// Reconcile reads that state of the cluster for a BatchJob object and makes changes based on the state read
// and what is in the BatchJob.Spec
// TODO(user): Modify this Reconcile function to implement your Controller logic.  This example creates
// a Pod as an example
// Note:
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (r *ReconcileBatchJob) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	reqLogger.Info("Reconciling BatchJob")

	// Fetch the BatchJob instance
	instance := &batchv1beta1.BatchJob{}
	err := r.client.Get(context.TODO(), request.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			reqLogger.Info("Into cleanup logic")
			r.cleanupApps(request)
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}

	instanceUpdate := instance.DeepCopy()
	jobManager := JobManagerFactory(instanceUpdate.Spec.Type, r.client)
	switch instanceUpdate.Status.JobState {
	case batchv1beta1.NewState:
		jobApp, err := jobManager.submitJob(instanceUpdate)
		if err != nil {
			instanceUpdate.Status.JobState = batchv1beta1.SubmitFailedState
			return reconcile.Result{}, err
		}
		instanceUpdate.Status = batchv1beta1.BatchJobStatus{
			JobState: batchv1beta1.SubmittedState,
		}
		if err := controllerutil.SetControllerReference(instanceUpdate, jobApp, r.scheme); err != nil {
			reqLogger.Error(err, "failed to set controller")
			return reconcile.Result{}, err
		}
		instanceUpdate.Status.JobState = batchv1beta1.SubmittedState
	case batchv1beta1.RetryState:
		// TODO: How to retry in this section
		return reconcile.Result{}, nil
	case batchv1beta1.SubmittedState, batchv1beta1.RunningState:
		if err := jobManager.followBatchApplicationState(instanceUpdate); err != nil {
			return reconcile.Result{}, err
		}
		// return reconcile.Result{}, nil
	case batchv1beta1.SubmitFailedState:
		if r.isRetry(instanceUpdate) {
			instanceUpdate.Status.JobState = batchv1beta1.RetryState
		} else {
			instanceUpdate.Status.JobState = batchv1beta1.FailedState
		}
		return reconcile.Result{}, nil
	case batchv1beta1.CompletedState:
		return reconcile.Result{}, nil
	case batchv1beta1.FailedState:
		namespacedName := types.NamespacedName{Namespace: instanceUpdate.GetNamespace(), Name: instanceUpdate.GetName()}
		if err := jobManager.deleteJob(namespacedName); err != nil {
			return reconcile.Result{}, err
		}
		return reconcile.Result{}, nil
	}

	// Set BatchJob instance as the owner and controller
	// if err := controllerutil.SetControllerReference(instance, pod, r.scheme); err != nil {
	// 	return reconcile.Result{}, err
	// }

	// Check if this Pod already exists
	if instanceUpdate != nil {
		reqLogger.Info("Update BatchJob Status")
		err := r.updateBatchJobStatus(instance, instanceUpdate)
		if err != nil {
			reqLogger.Info("Update BatchJob Status Failed")
			return reconcile.Result{}, err
		}
	}

	// Pod already exists - don't requeue
	// reqLogger.Info("Skip reconcile: Pod already exists", "Pod.Namespace", found.Namespace, "Pod.Name", found.Name)
	return reconcile.Result{}, nil
}

// newPodForCR returns a busybox pod with the same name/namespace as the cr
func newPodForCR(cr *batchv1beta1.BatchJob) *corev1.Pod {
	labels := map[string]string{
		"app": cr.Name,
	}
	return &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      cr.Name + "-pod",
			Namespace: cr.Namespace,
			Labels:    labels,
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:    "busybox",
					Image:   "busybox",
					Command: []string{"sleep", "3600"},
				},
			},
		},
	}
}

func (r *ReconcileBatchJob) cleanupApps(request reconcile.Request) error {
	sparkapp := &sparkv1beta2.SparkApplication{}
	err := r.client.Get(context.TODO(), request.NamespacedName, sparkapp)
	// log.Info("no get sparkAPP: ",sparkapp)
	if err != nil {
		if errors.IsNotFound(err) {
			return nil
		}
		return err
	}
	err = r.client.Delete(context.TODO(), sparkapp)
	if err != nil {
		log.Info("unable to delete: ", err)
		return err
	}
	return nil
}

func (r *ReconcileBatchJob) updateBatchJobStatus(oldapp, newapp *batchv1beta1.BatchJob) error {
	if equality.Semantic.DeepEqual(oldapp, newapp) {
		// log.Info("exactly Equal newAPP:")
		return nil
	}
	err := r.client.Update(context.Background(), newapp)
	err = r.client.Status().Update(context.Background(), newapp)
	return err
}

func (r *ReconcileBatchJob) isRetry(job *batchv1beta1.BatchJob) bool {
	return true
}
