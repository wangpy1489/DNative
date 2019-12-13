// NOTE: Boilerplate only.  Ignore this file.

// Package v1beta1 contains API Schema definitions for the batch v1beta1 API group
// +k8s:deepcopy-gen=package,register
// +groupName=batch.k8s.io
package v1beta1

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/scheme"
)

var (
	// SchemeGroupVersion is group version used to register these objects
	SchemeGroupVersion = schema.GroupVersion{Group: "batch.k8s.io", Version: "v1beta1"}

	// SchemeBuilder is used to add go types to the GroupVersionKind scheme
	SchemeBuilder = &scheme.Builder{GroupVersion: SchemeGroupVersion}
)
