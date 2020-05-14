package storage

import (
	"context"

	batchv1beta1 "github.com/wangpy1489/DNative/pkg/apis/batch/v1beta1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// 记录Stroe与APP对应情况
type StoreCore struct {
	client client.Client
}

func MakeSotreCore(client client.Client) *StoreCore {
	return &StoreCore{
		client: client,
	}
}

func (s *StoreCore) VolumeBuilder(bt batchv1beta1.BatchTemplate, appname string) (*corev1.Volume, error) {
	ss := &batchv1beta1.StorageSource{}
	err := s.client.Get(context.TODO(), types.NamespacedName{Namespace: bt.Namespace, Name: bt.Spec.StroageInfo.StorageName}, ss)
	if err != nil {
		return nil, err
	}
	pv, err := s.createPV(ss, appname)
	if err != nil {
		// fmt.Println("Heer>?")
		return nil, err
	}
	pvc, err := s.createPVC(pv, appname)
	if err != nil {
		// fmt.Println("Hee!>?")
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

func (s *StoreCore) createPV(ss *batchv1beta1.StorageSource, appname string) (*corev1.PersistentVolume, error) {
	pv := &corev1.PersistentVolume{
		ObjectMeta: metav1.ObjectMeta{
			Name:      appname + "-pv",
			Namespace: ss.Namespace,
		},
		Spec: corev1.PersistentVolumeSpec{
			Capacity:               ss.Spec.Capacity,
			PersistentVolumeSource: ss.Spec.Source,
			AccessModes:            []corev1.PersistentVolumeAccessMode{corev1.ReadWriteOnce},
		},
	}
	err := s.client.Create(context.TODO(), pv)
	if err != nil {
		return nil, err
	}
	return pv, nil
}

func (s *StoreCore) createPVC(pv *corev1.PersistentVolume, appname string) (*corev1.PersistentVolumeClaim, error) {
	pvc := &corev1.PersistentVolumeClaim{
		ObjectMeta: metav1.ObjectMeta{
			Name:      appname + "-pvc",
			Namespace: pv.Namespace,
		},
		Spec: corev1.PersistentVolumeClaimSpec{
			Resources:   corev1.ResourceRequirements{Requests: pv.Spec.Capacity},
			AccessModes: []corev1.PersistentVolumeAccessMode{corev1.ReadWriteOnce},
			VolumeName:  pv.Name,
		},
	}
	err := s.client.Create(context.TODO(), pvc)
	if err != nil {
		return nil, err
	}
	return pvc, nil
}

// 获取状态 生成具体储存下的结构体
type StorageManager interface {
	VolumeBuilder(bt batchv1beta1.BatchTemplate, appname string) (*corev1.Volume, error)

	createPV(ss *batchv1beta1.StorageSource, appname string) (*corev1.PersistentVolume, error)

	createPVC(pv *corev1.PersistentVolume, appname string) (*corev1.PersistentVolumeClaim, error)
}
