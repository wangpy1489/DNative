package batchjob

import (
	"context"

	sparkv1beta2 "github.com/GoogleCloudPlatform/spark-on-k8s-operator/pkg/apis/sparkoperator.k8s.io/v1beta2"
	batchv1beta1 "github.com/wangpy1489/DNative/pkg/apis/batch/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type SparkManger struct {
	client client.Client
}

func (m *SparkManger) submitJob(job *batchv1beta1.BatchJob) (metav1.Object, error) {
	newApp := &sparkv1beta2.SparkApplication{
		ObjectMeta: metav1.ObjectMeta{
			Name:      job.Name,
			Namespace: job.Namespace,
		},
		Spec: job.Spec.Templete.SparkApplicationSpec,
	}
	err := m.client.Create(context.TODO(), newApp)
	return newApp, err
}
func (m *SparkManger) deleteJob(namespacedName types.NamespacedName) error {
	sparkapp := &sparkv1beta2.SparkApplication{}
	err := m.client.Get(context.TODO(), namespacedName, sparkapp)
	if err != nil {
		if errors.IsNotFound(err) {
			return nil
		}
		return err
	}
	err = m.client.Delete(context.TODO(), sparkapp)
	if err != nil {
		log.Info("unable to delete: ", err)
		return err
	}
	return nil
}
func (m *SparkManger) followBatchApplicationState(job *batchv1beta1.BatchJob) error {
	sparkapp := &sparkv1beta2.SparkApplication{}
	err := m.client.Get(context.TODO(), types.NamespacedName{Namespace: job.Namespace, Name: job.Name}, sparkapp)
	// log.Info("Now Get", "SparkAPP", sparkapp.Kind)
	if err != nil {
		if errors.IsNotFound(err) {
			job.Status.JobState = batchv1beta1.FailedState
			return nil
		}
		// log.Error(err, "Problem in Get sparkAPP")
		return err
	}
	switch sparkapp.Status.AppState.State {
	case sparkv1beta2.SubmittedState:
		job.Status.JobState = batchv1beta1.SubmittedState
	case sparkv1beta2.RunningState:
		job.Status.JobState = batchv1beta1.RunningState
	case sparkv1beta2.CompletedState:
		job.Status.JobState = batchv1beta1.CompletedState
	case sparkv1beta2.FailedState:
		job.Status.JobState = batchv1beta1.FailedState
	case sparkv1beta2.FailedSubmissionState:
		job.Status.JobState = batchv1beta1.SubmitFailedState

	}
	return nil
}
