package resource

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	sparkv1beta2 "github.com/GoogleCloudPlatform/spark-on-k8s-operator/pkg/apis/sparkoperator.k8s.io/v1beta2"
	batchv1beta1 "github.com/wangpy1489/DNative/pkg/apis/batch/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/client"

)


func SubmitBatchJob (kubeclient client.Client, job *batchv1beta1.BatchJob) (*sparkv1beta2.SparkApplication, error) {
	newApp := &sparkv1beta2.SparkApplication {
		TypeMeta: metav1.TypeMeta{
			APIVersion: "sparkoperator.k8s.io/v1beta2",
			Kind: "Application",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: job.Name,
			Namespace: job.Namespace,
		},
		Spec: sparkv1beta2.SparkApplicationSpec{
			Type: "Scala",
			Mode: "cluster",
			Image: &job.Spec.Image,
			MainClass: new(string),
			MainApplicationFile: new(string),
			Driver: sparkv1beta2.DriverSpec{
				SparkPodSpec: sparkv1beta2.SparkPodSpec{
					VolumeMounts: []corev1.VolumeMount{
						corev1.VolumeMount{
							Name: "test-volume",
							MountPath: "/tmp",
						},
					},
				},
				ServiceAccount: new(string),
			},
			Executor: sparkv1beta2.ExecutorSpec{
				SparkPodSpec: sparkv1beta2.SparkPodSpec{
					VolumeMounts: []corev1.VolumeMount{
						corev1.VolumeMount{
							Name: "test-volume",
							MountPath: "/tmp",
						},
					},
				},
				Instances: new(int32),
			},
			
		},
	}
	*newApp.Spec.MainClass =  "org.apache.spark.examples.SparkPi"
	*newApp.Spec.MainApplicationFile = "local:///opt/spark/examples/jars/spark-examples_2.11-2.4.4.jar"
	*newApp.Spec.Executor.Instances = 1
	*newApp.Spec.Driver.ServiceAccount = "sparkoperator-spark"
	err := kubeclient.Create(context.TODO(),newApp)
	if err != nil {
		return nil, err
	}
	// newApp, err := r.sparkClient.Sparkoperator().SparkApplications(job.Namespace).Create(newApp)
	
	return newApp, nil
}