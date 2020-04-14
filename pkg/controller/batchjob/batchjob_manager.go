package batchjob

import (
	"context"

	batchv1beta1 "github.com/wangpy1489/DNative/pkg/apis/batch/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	corev1 "k8s.io/kubernetes/pkg/apis/core"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type JobManager interface {
	submitJob(job *batchv1beta1.BatchJob) (metav1.Object, error)
	deleteJob(namespacedName types.NamespacedName) error
	followBatchApplicationState(job *batchv1beta1.BatchJob) error
}

func JobManagerFactory(jobType batchv1beta1.BatchJobType, kubeclient client.Client) JobManager {
	switch jobType {
	case batchv1beta1.Batch:
		return &BatchManger{kubeclient}

	case batchv1beta1.Spark:
		return &SparkManger{kubeclient}
	}
	return &SparkManger{kubeclient}
}

type BatchManger struct {
	client client.Client
}

func (m *BatchManger) submitJob(job *batchv1beta1.BatchJob) (metav1.Object, error) {
	newApp := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      job.Name,
			Namespace: job.Namespace,
		},
		Spec: job.Spec.Templete.Spec,
	}
	err := m.client.Create(context.TODO(), newApp)
	// Set Memcached instance as the owner and controller
	// controllerutil.SetControllerReference(m, dep, r.scheme)
	return newApp, err
}
func (m *BatchManger) deleteJob(namespacedName types.NamespacedName) error {
	pod := &corev1.Pod{}
	err := m.client.Get(context.TODO(), namespacedName, pod)
	if err != nil {
		if errors.IsNotFound(err) {
			return nil
		}
		return err
	}
	err = m.client.Delete(context.TODO(), pod)
	if err != nil {
		log.Info("unable to delete: ", err)
		return err
	}
	return nil
}
func (m *BatchManger) followBatchApplicationState(job *batchv1beta1.BatchJob) error {
	pod := &corev1.Pod{}
	err := m.client.Get(context.TODO(), types.NamespacedName{Namespace: job.Namespace, Name: job.Name}, pod)
	if err != nil {
		if errors.IsNotFound(err) {
			job.Status.JobState = batchv1beta1.FailedState
			return nil
		}
		return err
	}
	switch pod.Status.Phase {
	case corev1.PodRunning:
		job.Status.JobState = batchv1beta1.RunningState
	case corev1.PodFailed:
		job.Status.JobState = batchv1beta1.FailedState
	case corev1.PodPending:
		job.Status.JobState = batchv1beta1.SubmittedState
	case corev1.PodSucceeded:
		job.Status.JobState = batchv1beta1.CompletedState
	}
	return nil
}
