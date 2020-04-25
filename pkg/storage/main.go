package storage

import (
	"context"

	batchv1beta1 "github.com/wangpy1489/DNative/pkg/apis/batch/v1beta1"
	"k8s.io/apimachinery/pkg/types"
	corev1 "k8s.io/kubernetes/pkg/apis/core"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// 记录Stroe与APP对应情况
type StoreCore struct {
	client  client.Client
	manager StorageManager
}

func (s *StoreCore) VolumeBuilder(bt batchv1beta1.BatchTemplate, appname string) (*corev1.Volume, error) {
	ss := &batchv1beta1.StorageSource{}
	err := s.client.Get(context.TODO(), types.NamespacedName{Namespace: bt.Namespace, Name: bt.Spec.StorageName}, ss)
	if err != nil {
		return nil, err
	}
	pv, err := s.manager.createPV(ss, appname)
	if err != nil {
		return nil, err
	}
	pvc, err := s.manager.createPVC(pv, appname)
	if err != nil {
		return nil, err
	}
	volume := &corev1.Volume{
		Name: appname + "-volume",
		VolumeSource: corev1.VolumeSource{
			PersistentVolumeClaim: &corev1.PersistentVolumeClaimVolumeSource{
				ClaimName: pvc.Name,
			},
		},
	}
	return volume, nil
}

// 获取状态 生成具体储存下的结构体
type StorageManager interface {
	createPV(ss *batchv1beta1.StorageSource, appname string) (*corev1.PersistentVolume, error)

	createPVC(pv *corev1.PersistentVolume, appname string) (*corev1.PersistentVolumeClaim, error)
}
