package timer

import (
	"fmt"
	// "bytes"
	"io/ioutil"
	"net/http"
	"net/url"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func cacheKey(metadata *metav1.ObjectMeta) string {
	return fmt.Sprintf("%v_%v", metadata.UID, metadata.ResourceVersion)
}

func makeHTTPRequest(head string) error{
	u, _ := url.Parse(head)
	q := u.Query()
	u.RawQuery = q.Encode()
	res, err := http.Get(u.String())
	if err != nil {
		return err
	}
	result, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return err
	}
	fmt.Println(result)
	return nil
}