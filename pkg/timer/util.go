package timer

import (
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func cacheKey(metadata *metav1.ObjectMeta) string {
	return fmt.Sprintf("%v_%v", metadata.UID, metadata.ResourceVersion)
}