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

func (rou *Router) findStorageSource(r *http.Request) (*batchv1beta1.StorageSource, error) {
	vars := mux.Vars(r)
	name := vars["storageSource"]
	namespace := r.URL.Query().Get("namespace")
	storageSource := &batchv1beta1.StorageSource{}
	err := rou.kubeclient.Get(context.TODO(), types.NamespacedName{Namespace: namespace, Name: name}, storageSource)
	if err != nil {
		rou.logger.Error(err, err.Error())
		return nil, err
	}
	return storageSource, nil
}

func (rou *Router) StorageSourceApiList(w http.ResponseWriter, r *http.Request) {
	timers := &batchv1beta1.StorageSourceList{}
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

func (rou *Router) StorageSourceApiGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["storageSource"]
	namespace := r.URL.Query().Get("namespace")
	storageSource := &batchv1beta1.StorageSource{}
	rou.logger.Info(namespace, name)
	err := rou.kubeclient.Get(context.TODO(), types.NamespacedName{Namespace: namespace, Name: name}, storageSource)
	if err != nil {
		rou.logger.Error(err, err.Error())
		rou.respondWithError(w, err)
		return
	}
	resp, err := json.Marshal(storageSource)
	if err != nil {
		rou.logger.Error(err, err.Error())
		rou.respondWithError(w, err)
		return
	}
	rou.respondWithSuccess(w, resp)
}

func (rou *Router) StorageSourceApiCreate(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		rou.logger.Error(err, err.Error())
		rou.respondWithError(w, err)
		return
	}
	storageSource := batchv1beta1.StorageSource{}
	err = json.Unmarshal(body, &storageSource)
	if err != nil {
		rou.logger.Error(err, err.Error())
		rou.respondWithError(w, err)
		return
	}
	triggerFound := &batchv1beta1.StorageSource{}
	err = rou.kubeclient.Get(context.TODO(), types.NamespacedName{Namespace: storageSource.Namespace, Name: storageSource.Name}, triggerFound)
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

	err = rou.kubeclient.Create(context.TODO(), &storageSource)
	if err != nil {
		rou.logger.Error(err, err.Error())
		rou.respondWithError(w, err)
		return
	}

	resp, err := json.Marshal(storageSource)
	if err != nil {
		rou.logger.Error(err, err.Error())
		rou.respondWithError(w, err)
		return
	}
	rou.respondWithSuccess(w, resp)
}

func (rou *Router) StorageSourceApiUpdate(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		rou.logger.Error(err, err.Error())
		rou.respondWithError(w, err)
		return
	}
	storageSource := batchv1beta1.StorageSource{}
	err = json.Unmarshal(body, &storageSource)
	if err != nil {
		rou.logger.Error(err, err.Error())
		rou.respondWithError(w, err)
		return
	}
	triggerFound := &batchv1beta1.StorageSource{}
	err = rou.kubeclient.Get(context.TODO(), types.NamespacedName{Namespace: storageSource.Namespace, Name: storageSource.Name}, triggerFound)
	if err != nil {
		rou.logger.Error(err, err.Error())
		rou.respondWithError(w, err)
		return
	}
	triggerFound.Spec = storageSource.Spec
	err = rou.kubeclient.Update(context.TODO(), triggerFound)
	if err != nil {
		rou.logger.Error(err, err.Error())
		rou.respondWithError(w, err)
		return
	}

	resp, err := json.Marshal(storageSource)
	if err != nil {
		rou.logger.Error(err, err.Error())
		rou.respondWithError(w, err)
		return
	}
	rou.respondWithSuccess(w, resp)

}

func (rou *Router) StorageSourceApiDelete(w http.ResponseWriter, r *http.Request) {
	storageSource, err := rou.findStorageSource(r)
	if err != nil {
		rou.logger.Error(err, err.Error())
		rou.respondWithError(w, err)
		return
	}
	err = rou.kubeclient.Delete(context.TODO(), storageSource)
	if err != nil {
		rou.logger.Error(err, err.Error())
		rou.respondWithError(w, err)
		return
	}

	resp, err := json.Marshal(storageSource)
	rou.respondWithSuccess(w, resp)
}
