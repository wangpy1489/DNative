// +build !ignore_autogenerated

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1beta1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/wangpy1489/DNative/pkg/apis/batch/v1beta1.BatchJob":            schema_pkg_apis_batch_v1beta1_BatchJob(ref),
		"github.com/wangpy1489/DNative/pkg/apis/batch/v1beta1.BatchJobSpec":        schema_pkg_apis_batch_v1beta1_BatchJobSpec(ref),
		"github.com/wangpy1489/DNative/pkg/apis/batch/v1beta1.BatchJobStatus":      schema_pkg_apis_batch_v1beta1_BatchJobStatus(ref),
		"github.com/wangpy1489/DNative/pkg/apis/batch/v1beta1.BatchTemplate":       schema_pkg_apis_batch_v1beta1_BatchTemplate(ref),
		"github.com/wangpy1489/DNative/pkg/apis/batch/v1beta1.BatchTemplateSpec":   schema_pkg_apis_batch_v1beta1_BatchTemplateSpec(ref),
		"github.com/wangpy1489/DNative/pkg/apis/batch/v1beta1.BatchTemplateStatus": schema_pkg_apis_batch_v1beta1_BatchTemplateStatus(ref),
		"github.com/wangpy1489/DNative/pkg/apis/batch/v1beta1.HttpTrigger":         schema_pkg_apis_batch_v1beta1_HttpTrigger(ref),
		"github.com/wangpy1489/DNative/pkg/apis/batch/v1beta1.HttpTriggerSpec":     schema_pkg_apis_batch_v1beta1_HttpTriggerSpec(ref),
		"github.com/wangpy1489/DNative/pkg/apis/batch/v1beta1.HttpTriggerStatus":   schema_pkg_apis_batch_v1beta1_HttpTriggerStatus(ref),
		"github.com/wangpy1489/DNative/pkg/apis/batch/v1beta1.StorageSource":       schema_pkg_apis_batch_v1beta1_StorageSource(ref),
		"github.com/wangpy1489/DNative/pkg/apis/batch/v1beta1.StorageSourceSpec":   schema_pkg_apis_batch_v1beta1_StorageSourceSpec(ref),
		"github.com/wangpy1489/DNative/pkg/apis/batch/v1beta1.StorageSourceStatus": schema_pkg_apis_batch_v1beta1_StorageSourceStatus(ref),
		"github.com/wangpy1489/DNative/pkg/apis/batch/v1beta1.TimerTrigger":        schema_pkg_apis_batch_v1beta1_TimerTrigger(ref),
		"github.com/wangpy1489/DNative/pkg/apis/batch/v1beta1.TimerTriggerSpec":    schema_pkg_apis_batch_v1beta1_TimerTriggerSpec(ref),
		"github.com/wangpy1489/DNative/pkg/apis/batch/v1beta1.TimerTriggerStatus":  schema_pkg_apis_batch_v1beta1_TimerTriggerStatus(ref),
	}
}

func schema_pkg_apis_batch_v1beta1_BatchJob(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "BatchJob is the Schema for the batchjobs API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/wangpy1489/DNative/pkg/apis/batch/v1beta1.BatchJobSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/wangpy1489/DNative/pkg/apis/batch/v1beta1.BatchJobStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/wangpy1489/DNative/pkg/apis/batch/v1beta1.BatchJobSpec", "github.com/wangpy1489/DNative/pkg/apis/batch/v1beta1.BatchJobStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_batch_v1beta1_BatchJobSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "BatchJobSpec defines the desired state of BatchJob",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"type": {
						SchemaProps: spec.SchemaProps{
							Description: "INSERT ADDITIONAL SPEC FIELDS - desired state of cluster Important: Run \"operator-sdk generate k8s\" to regenerate code after modifying this file Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"template": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/wangpy1489/DNative/pkg/apis/batch/v1beta1.JobTemplete"),
						},
					},
				},
				Required: []string{"type", "template"},
			},
		},
		Dependencies: []string{
			"github.com/wangpy1489/DNative/pkg/apis/batch/v1beta1.JobTemplete"},
	}
}

func schema_pkg_apis_batch_v1beta1_BatchJobStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "BatchJobStatus defines the observed state of BatchJob",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"jobState": {
						SchemaProps: spec.SchemaProps{
							Description: "INSERT ADDITIONAL STATUS FIELD - define observed state of cluster Important: Run \"operator-sdk generate k8s\" to regenerate code after modifying this file Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
			},
		},
	}
}

func schema_pkg_apis_batch_v1beta1_BatchTemplate(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "BatchTemplate is the Schema for the batchtemplates API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/wangpy1489/DNative/pkg/apis/batch/v1beta1.BatchTemplateSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/wangpy1489/DNative/pkg/apis/batch/v1beta1.BatchTemplateStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/wangpy1489/DNative/pkg/apis/batch/v1beta1.BatchTemplateSpec", "github.com/wangpy1489/DNative/pkg/apis/batch/v1beta1.BatchTemplateStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_batch_v1beta1_BatchTemplateSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "BatchTemplateSpec defines the desired state of BatchTemplate",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"type": {
						SchemaProps: spec.SchemaProps{
							Description: "INSERT ADDITIONAL SPEC FIELDS - desired state of cluster Important: Run \"operator-sdk generate k8s\" to regenerate code after modifying this file Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"template": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/wangpy1489/DNative/pkg/apis/batch/v1beta1.JobTemplete"),
						},
					},
					"storageName": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
				Required: []string{"type", "template", "storageName"},
			},
		},
		Dependencies: []string{
			"github.com/wangpy1489/DNative/pkg/apis/batch/v1beta1.JobTemplete"},
	}
}

func schema_pkg_apis_batch_v1beta1_BatchTemplateStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "BatchTemplateStatus defines the observed state of BatchTemplate",
				Type:        []string{"object"},
			},
		},
	}
}

func schema_pkg_apis_batch_v1beta1_HttpTrigger(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "HttpTrigger is the Schema for the httptriggers API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/wangpy1489/DNative/pkg/apis/batch/v1beta1.HttpTriggerSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/wangpy1489/DNative/pkg/apis/batch/v1beta1.HttpTriggerStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/wangpy1489/DNative/pkg/apis/batch/v1beta1.HttpTriggerSpec", "github.com/wangpy1489/DNative/pkg/apis/batch/v1beta1.HttpTriggerStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_batch_v1beta1_HttpTriggerSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "HttpTriggerSpec defines the desired state of HttpTrigger",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"jobref": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/wangpy1489/DNative/pkg/apis/batch/v1beta1.JobReference"),
						},
					},
					"relativeurl": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
				Required: []string{"jobref", "relativeurl"},
			},
		},
		Dependencies: []string{
			"github.com/wangpy1489/DNative/pkg/apis/batch/v1beta1.JobReference"},
	}
}

func schema_pkg_apis_batch_v1beta1_HttpTriggerStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "HttpTriggerStatus defines the observed state of HttpTrigger",
				Type:        []string{"object"},
			},
		},
	}
}

func schema_pkg_apis_batch_v1beta1_StorageSource(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "StorageSource is the Schema for the storagesources API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/wangpy1489/DNative/pkg/apis/batch/v1beta1.StorageSourceSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/wangpy1489/DNative/pkg/apis/batch/v1beta1.StorageSourceStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/wangpy1489/DNative/pkg/apis/batch/v1beta1.StorageSourceSpec", "github.com/wangpy1489/DNative/pkg/apis/batch/v1beta1.StorageSourceStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_batch_v1beta1_StorageSourceSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "StorageSourceSpec defines the desired state of StorageSource",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"type": {
						SchemaProps: spec.SchemaProps{
							Description: "INSERT ADDITIONAL SPEC FIELDS - desired state of cluster Important: Run \"operator-sdk generate k8s\" to regenerate code after modifying this file Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"source": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/kubernetes/pkg/apis/core.PersistentVolumeSource"),
						},
					},
				},
				Required: []string{"type", "source"},
			},
		},
		Dependencies: []string{
			"k8s.io/kubernetes/pkg/apis/core.PersistentVolumeSource"},
	}
}

func schema_pkg_apis_batch_v1beta1_StorageSourceStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "StorageSourceStatus defines the observed state of StorageSource",
				Type:        []string{"object"},
			},
		},
	}
}

func schema_pkg_apis_batch_v1beta1_TimerTrigger(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "TimerTrigger is the Schema for the timertriggers API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/wangpy1489/DNative/pkg/apis/batch/v1beta1.TimerTriggerSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/wangpy1489/DNative/pkg/apis/batch/v1beta1.TimerTriggerStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/wangpy1489/DNative/pkg/apis/batch/v1beta1.TimerTriggerSpec", "github.com/wangpy1489/DNative/pkg/apis/batch/v1beta1.TimerTriggerStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_batch_v1beta1_TimerTriggerSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "TimerTriggerSpec defines the desired state of TimerTrigger",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"cron": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"jobref": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/wangpy1489/DNative/pkg/apis/batch/v1beta1.JobReference"),
						},
					},
				},
				Required: []string{"cron", "jobref"},
			},
		},
		Dependencies: []string{
			"github.com/wangpy1489/DNative/pkg/apis/batch/v1beta1.JobReference"},
	}
}

func schema_pkg_apis_batch_v1beta1_TimerTriggerStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "TimerTriggerStatus defines the observed state of TimerTrigger",
				Type:        []string{"object"},
			},
		},
	}
}
