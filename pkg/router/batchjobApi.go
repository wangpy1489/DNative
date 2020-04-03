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

func (rou *Router) BatchJobApiList(w http.ResponseWriter, r *http.Request) {
	batchjobs := &batchv1beta1.BatchJobList{}
	err := rou.kubeclient.List(context.TODO(), batchjobs)
	if err != nil {
		rou.logger.Error(err, err.Error())
		return
	}
	resp, err := json.Marshal(batchjobs)
	if err != nil {
		rou.logger.Error(err, err.Error())
		return
	}
	rou.respondWithSuccess(w, resp)
}

func (rou *Router) BatchJobApiGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["batchjob"]
	namespace := r.URL.Query().Get("namespace")
	batchjob := &batchv1beta1.BatchJob{}
	rou.logger.Info(namespace, name)
	err := rou.kubeclient.Get(context.TODO(), types.NamespacedName{Namespace: namespace, Name: name}, batchjob)
	if err != nil {
		rou.logger.Error(err, err.Error())
		return
	}
	resp, err := json.Marshal(batchjob)
	if err != nil {
		rou.logger.Error(err, err.Error())
		return
	}
	rou.respondWithSuccess(w, resp)
}

func (rou *Router) BatchJobApiCreate(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		rou.logger.Error(err, err.Error())
		return
	}
	var batchJob batchv1beta1.BatchJob
	err = json.Unmarshal(body, &batchJob)
	if err != nil {
		rou.logger.Error(err, err.Error())
		return
	}
	triggerFound := &batchv1beta1.BatchJob{}
	err = rou.kubeclient.Get(context.TODO(), types.NamespacedName{Namespace: batchJob.Namespace, Name: batchJob.Name}, triggerFound)
	if err != nil {
		if !errors.IsNotFound(err) {
			rou.logger.Error(err, err.Error())
			return
		}
	} else if triggerFound != nil {
		rou.logger.Error(fmt.Errorf("%s", "HTTP Trigger already existed."), "Existed")
		return
	}

	err = rou.kubeclient.Create(context.TODO(), &batchJob)
	if err != nil {
		rou.logger.Error(err, err.Error())
		return
	}

	resp, err := json.Marshal(batchJob)
	if err != nil {
		rou.logger.Error(err, err.Error())
		return
	}
	rou.respondWithSuccess(w, resp)
}

func (rou *Router) BatchJobApiUpdate(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		rou.logger.Error(err, err.Error())
		return
	}
	var batchJob batchv1beta1.BatchJob
	err = json.Unmarshal(body, &batchJob)
	if err != nil {
		rou.logger.Error(err, err.Error())
		return
	}
	triggerFound := &batchv1beta1.BatchJob{}
	err = rou.kubeclient.Get(context.TODO(), types.NamespacedName{Namespace: batchJob.Namespace, Name: batchJob.Name}, triggerFound)
	if err != nil {
		rou.logger.Error(err, err.Error())
		return
	}
	err = rou.kubeclient.Update(context.TODO(), &batchJob)
	if err != nil {
		rou.logger.Error(err, err.Error())
		return
	}

	resp, err := json.Marshal(batchJob)
	if err != nil {
		rou.logger.Error(err, err.Error())
		return
	}
	rou.respondWithSuccess(w, resp)

}

func (rou *Router) BatchJobApiDelete(w http.ResponseWriter, r *http.Request) {
	batchJob, err := rou.findHttptrigger(r)
	if err != nil {
		rou.logger.Error(err, err.Error())
		return
	}
	err = rou.kubeclient.Delete(context.TODO(), batchJob)
	if err != nil {
		rou.logger.Error(err, err.Error())
		return
	}

	resp, err := json.Marshal(batchJob)
	rou.respondWithSuccess(w, resp)
}
