package router

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	batchv1beta1 "github.com/wangpy1489/DNative/pkg/apis/batch/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
)

func (rou *Router) findBatchTemplate(r *http.Request) (*batchv1beta1.BatchTemplate, error) {
	vars := mux.Vars(r)
	name := vars["batchTemplate"]
	namespace := r.URL.Query().Get("namespace")
	batchTemplate := &batchv1beta1.BatchTemplate{}
	err := rou.kubeclient.Get(context.TODO(), types.NamespacedName{Namespace: namespace, Name: name}, batchTemplate)
	if err != nil {
		rou.logger.Error(err, err.Error())
		return nil, err
	}
	return batchTemplate, nil
}

func (rou *Router) BatchTemplateApiList(w http.ResponseWriter, r *http.Request) {
	timers := &batchv1beta1.BatchTemplateList{}
	err := rou.kubeclient.List(context.TODO(), timers)
	if err != nil {
		rou.logger.Error(err, err.Error())
		rou.respondWithError(w, err)
		return
	}
	resp, err := json.Marshal(timers)
	if err != nil {
		rou.logger.Error(err, err.Error())
		rou.respondWithError(w, err)
		return
	}
	rou.respondWithSuccess(w, resp)
}

func (rou *Router) BatchTemplateApiGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["batchTemplate"]
	namespace := r.URL.Query().Get("namespace")
	batchTemplate := &batchv1beta1.BatchTemplate{}
	rou.logger.Info(namespace, name)
	err := rou.kubeclient.Get(context.TODO(), types.NamespacedName{Namespace: namespace, Name: name}, batchTemplate)
	if err != nil {
		rou.logger.Error(err, err.Error())
		rou.respondWithError(w, err)
		return
	}
	resp, err := json.Marshal(batchTemplate)
	if err != nil {
		rou.logger.Error(err, err.Error())
		rou.respondWithError(w, err)
		return
	}
	rou.respondWithSuccess(w, resp)
}

func (rou *Router) BatchTemplateApiCreate(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		rou.logger.Error(err, err.Error())
		rou.respondWithError(w, err)
		return
	}
	batchTemplate := batchv1beta1.BatchTemplate{}
	err = json.Unmarshal(body, &batchTemplate)
	if err != nil {
		rou.logger.Error(err, err.Error())
		rou.respondWithError(w, err)
		return
	}
	triggerFound := &batchv1beta1.BatchTemplate{}
	err = rou.kubeclient.Get(context.TODO(), types.NamespacedName{Namespace: batchTemplate.Namespace, Name: batchTemplate.Name}, triggerFound)
	if err != nil {
		if !errors.IsNotFound(err) {
			rou.logger.Error(err, err.Error())
			rou.respondWithError(w, err)
			return
		}
	} else if triggerFound != nil {
		rou.logger.Error(fmt.Errorf("%s", "HTTP Trigger already existed."), "Existed")
		rou.respondWithError(w, err)
		return
	}

	err = rou.kubeclient.Create(context.TODO(), &batchTemplate)
	if err != nil {
		rou.logger.Error(err, err.Error())
		rou.respondWithError(w, err)
		return
	}

	resp, err := json.Marshal(batchTemplate)
	if err != nil {
		rou.logger.Error(err, err.Error())
		rou.respondWithError(w, err)
		return
	}
	rou.respondWithSuccess(w, resp)
}

func (rou *Router) BatchTemplateApiUpdate(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		rou.logger.Error(err, err.Error())
		rou.respondWithError(w, err)
		return
	}
	batchTemplate := &batchv1beta1.BatchTemplate{}
	err = json.Unmarshal(body, batchTemplate)
	if err != nil {
		rou.logger.Error(err, err.Error())
		rou.respondWithError(w, err)
		return
	}
	triggerFound := &batchv1beta1.BatchTemplate{}
	err = rou.kubeclient.Get(context.TODO(), types.NamespacedName{Namespace: batchTemplate.Namespace, Name: batchTemplate.Name}, triggerFound)
	if err != nil {
		rou.logger.Error(err, err.Error())
		rou.respondWithError(w, err)
		return
	}
	triggerFound.Spec = batchTemplate.Spec
	// rou.logger.Info("now", "Update", batchTemplate, "origin", triggerFound, "json", string(body))
	err = rou.kubeclient.Update(context.TODO(), triggerFound)
	if err != nil {
		rou.logger.Error(err, err.Error())
		rou.respondWithError(w, err)
		return
	}

	resp, err := json.Marshal(batchTemplate)
	if err != nil {
		rou.logger.Error(err, err.Error())
		return
	}
	rou.respondWithSuccess(w, resp)

}

func (rou *Router) BatchTemplateApiDelete(w http.ResponseWriter, r *http.Request) {
	batchTemplate, err := rou.findBatchTemplate(r)
	if err != nil {
		rou.logger.Error(err, err.Error())
		return
	}
	err = rou.kubeclient.Delete(context.TODO(), batchTemplate)
	if err != nil {
		rou.logger.Error(err, err.Error())
		rou.respondWithError(w, err)
		return
	}

	resp, err := json.Marshal(batchTemplate)
	rou.respondWithSuccess(w, resp)
}
