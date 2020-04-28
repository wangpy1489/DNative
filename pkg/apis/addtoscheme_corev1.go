package apis

import (
	"k8s.io/kubernetes/pkg/apis/core"
)

func init() {
	AddToSchemes = append(AddToSchemes, core.SchemeBuilder.AddToScheme)

}