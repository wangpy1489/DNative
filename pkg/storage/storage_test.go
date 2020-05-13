package storage

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/wangpy1489/DNative/pkg/apis"
	batchv1beta1 "github.com/wangpy1489/DNative/pkg/apis/batch/v1beta1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes/scheme"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestStorageManager(t *testing.T) {
	storageSource := &batchv1beta1.StorageSource{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "ss-nfs",
			Namespace: "default",
		},
		Spec: batchv1beta1.StorageSourceSpec{
			Type: "nfs",
			Source: corev1.PersistentVolumeSource{
				NFS: &corev1.NFSVolumeSource{
					Server: "localhost",
					Path:   "/data",
				},
			},
		},
	}
	bt := &batchv1beta1.BatchTemplate{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "example",
			Namespace: "default",
		},
		Spec: batchv1beta1.BatchTemplateSpec{
			Type: batchv1beta1.Batch,
			Template: batchv1beta1.JobTemplate{
				Batch: corev1.PodTemplateSpec{
					Spec: corev1.PodSpec{
						Containers: []corev1.Container{
							{
								Name:  "test",
								Image: "busybox",
							},
						},
					},
				},
			},
		},
	}

	apis.AddToScheme(scheme.Scheme)
	objs := []runtime.Object{storageSource, bt}
	cl := fake.NewFakeClient(objs...)

	sc := StoreCore{cl}
	appname := "example01"
	volume, err := sc.VolumeBuilder(*bt, appname)
	if err != nil {
		t.Error(err)
	}
	data, err := json.Marshal(*volume)
	t.Logf("%s\n", data)
	pv := &corev1.PersistentVolume{}
	err = cl.Get(context.TODO(), types.NamespacedName{Namespace: "default", Name: appname + "-pv"}, pv)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v\n", *pv)
	pvc := &corev1.PersistentVolumeClaim{}
	err = cl.Get(context.TODO(), types.NamespacedName{Namespace: "default", Name: appname + "-pvc"}, pvc)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v\n", *pvc)
}
