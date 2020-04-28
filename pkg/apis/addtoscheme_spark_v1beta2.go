package apis

import (
	"github.com/GoogleCloudPlatform/spark-on-k8s-operator/pkg/apis/sparkoperator.k8s.io/v1beta2"
)

func init() {
	AddToSchemes = append(AddToSchemes, v1beta2.AddToScheme)

}
