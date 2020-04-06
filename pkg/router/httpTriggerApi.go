package router

import (
	// "fmt"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	batchv1beta1 "github.com/wangpy1489/DNative/pkg/apis/batch/v1beta1"
	"github.com/wangpy1489/DNative/pkg/router/resource"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
)

func (rou *Router) findHttptrigger(r *http.Request) (*batchv1beta1.HttpTrigger, error) {
	vars := mux.Vars(r)
	name := vars["httpTrigger"]
	namespace := r.URL.Query().Get("namespace")
	httpTrigger := &batchv1beta1.HttpTrigger{}
	err := rou.kubeclient.Get(context.TODO(), types.NamespacedName{Namespace: namespace, Name: name}, httpTrigger)
	if err != nil {
		rou.logger.Error(err, err.Error())
		return nil, err
	}
	return httpTrigger, nil
}

func (rou *Router) HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write([]byte(rou.info))
}

func (rou *Router) HttpTriggerApiList(w http.ResponseWriter, r *http.Request) {
	timers := &batchv1beta1.HttpTriggerList{}
	err := rou.kubeclient.List(context.TODO(), timers)
	if err != nil {
		rou.logger.Error(err, err.Error())
		return
	}
	resp, err := json.Marshal(timers)
	if err != nil {
		rou.logger.Error(err, err.Error())
		return
	}
	rou.respondWithSuccess(w, resp)
}

func (rou *Router) HttpTriggerApi(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["httpTrigger"]
	namespace := r.URL.Query().Get("namespace")
	httpTrigger := &batchv1beta1.HttpTrigger{}
	rou.logger.Info(namespace, name)
	err := rou.kubeclient.Get(context.TODO(), types.NamespacedName{Namespace: namespace, Name: name}, httpTrigger)
	if err != nil {
		rou.logger.Error(err, err.Error())
		return
	}
	batchJob := &batchv1beta1.BatchJob{}
	err = rou.kubeclient.Get(context.TODO(), types.NamespacedName{Namespace: httpTrigger.Namespace, Name: httpTrigger.Spec.JobReference.Name}, batchJob)
	if err != nil {
		rou.logger.Error(err, err.Error())
		return
	}
	sparkapp, err := resource.SubmitBatchJob(rou.kubeclient, batchJob)
	if err != nil {
		rou.logger.Error(err, err.Error())
		return
	}
	resp, err := json.Marshal(sparkapp)
	if err != nil {
		rou.logger.Error(err, err.Error())
		return
	}
	rou.respondWithSuccess(w, resp)
}

func (rou *Router) HttpTriggerApiUpdate(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		rou.logger.Error(err, err.Error())
		return
	}
	var httptrigger batchv1beta1.HttpTrigger
	err = json.Unmarshal(body, httptrigger)
	if err != nil {
		rou.logger.Error(err, err.Error())
		return
	}
	triggerFound := &batchv1beta1.HttpTrigger{}
	err = rou.kubeclient.Get(context.TODO(), types.NamespacedName{Namespace: httptrigger.Namespace, Name: httptrigger.Name}, triggerFound)
	if err != nil {
		rou.logger.Error(err, err.Error())
		return
	}
	err = rou.kubeclient.Update(context.TODO(), &httptrigger)
	if err != nil {
		rou.logger.Error(err, err.Error())
		return
	}

	resp, err := json.Marshal(httptrigger)
	if err != nil {
		rou.logger.Error(err, err.Error())
		return
	}
	rou.respondWithSuccess(w, resp)

}

func (rou *Router) HttpTriggerApiCreate(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		rou.logger.Error(err, err.Error())
		return
	}
	var httptrigger batchv1beta1.HttpTrigger
	err = json.Unmarshal(body, httptrigger)
	if err != nil {
		rou.logger.Error(err, err.Error())
		return
	}
	triggerFound := &batchv1beta1.HttpTrigger{}
	err = rou.kubeclient.Get(context.TODO(), types.NamespacedName{Namespace: httptrigger.Namespace, Name: httptrigger.Name}, triggerFound)
	if err != nil {
		if !errors.IsNotFound(err) {
			rou.logger.Error(err, err.Error())
			return
		}
	} else if triggerFound != nil {
		rou.logger.Error(fmt.Errorf("%s", "HTTP Trigger already existed."), "Existed")
		return
	}

	err = rou.kubeclient.Create(context.TODO(), &httptrigger)
	if err != nil {
		rou.logger.Error(err, err.Error())
		return
	}

	resp, err := json.Marshal(httptrigger)
	if err != nil {
		rou.logger.Error(err, err.Error())
		return
	}
	rou.respondWithSuccess(w, resp)
}

func (rou *Router) HttpTriggerApiDelete(w http.ResponseWriter, r *http.Request) {
	httptrigger, err := rou.findHttptrigger(r)
	if err != nil {
		rou.logger.Error(err, err.Error())
		return
	}
	err = rou.kubeclient.Delete(context.TODO(), httptrigger)
	if err != nil {
		rou.logger.Error(err, err.Error())
		return
	}

	resp, err := json.Marshal(httptrigger)
	rou.respondWithSuccess(w, resp)
}
