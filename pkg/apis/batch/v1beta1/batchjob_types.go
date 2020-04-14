package v1beta1

import (
	sparkv1 "github.com/GoogleCloudPlatform/spark-on-k8s-operator/pkg/apis/sparkoperator.k8s.io/v1beta2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	corev1 "k8s.io/kubernetes/pkg/apis/core"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// BatchJobSpec defines the desired state of BatchJob
// +k8s:openapi-gen=true
type BatchJobSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
	Type     BatchJobType `json:"type"`
	Templete JobTemplete  `json:"template"`
}

type BatchJobType string

const (
	Batch BatchJobType = "batch"
	Spark BatchJobType = "spark"
)

type JobTemplete struct {
	sparkv1.SparkApplicationSpec
	corev1.PodTemplateSpec
}

type BatchJobState string

const (
	NewState          BatchJobState = ""
	SubmittedState    BatchJobState = "SUBMITTED"
	RunningState      BatchJobState = "RUNNING"
	CompletedState    BatchJobState = "COMPLETED"
	SubmitFailedState BatchJobState = "SUBMIT_FAILED"
	FailedState       BatchJobState = "FAILED"
	RetryState        BatchJobState = "RETRY"
)

// BatchJobStatus defines the observed state of BatchJob
// +k8s:openapi-gen=true
type BatchJobStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
	JobState BatchJobState `json:"jobState,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BatchJob is the Schema for the batchjobs API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=batchjobs,scope=Namespaced
type BatchJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BatchJobSpec   `json:"spec,omitempty"`
	Status BatchJobStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BatchJobList contains a list of BatchJob
type BatchJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BatchJob `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BatchJob{}, &BatchJobList{})
}
