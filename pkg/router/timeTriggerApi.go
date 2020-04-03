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

func (rou *Router) TimerTriggerApiList(w http.ResponseWriter, r *http.Request) {
	timers := &batchv1beta1.TimerTriggerList{}
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

func (rou *Router) TimerTriggerApiGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["timerTrigger"]
	namespace := r.URL.Query().Get("namespace")
	timerTrigger := &batchv1beta1.TimerTrigger{}
	rou.logger.Info(namespace, name)
	err := rou.kubeclient.Get(context.TODO(), types.NamespacedName{Namespace: namespace, Name: name}, timerTrigger)
	if err != nil {
		rou.logger.Error(err, err.Error())
		return
	}
	resp, err := json.Marshal(timerTrigger)
	if err != nil {
		rou.logger.Error(err, err.Error())
		return
	}
	rou.respondWithSuccess(w, resp)
}

func (rou *Router) TimerTriggerApiCreate(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		rou.logger.Error(err, err.Error())
		return
	}
	var timerTrigger batchv1beta1.TimerTrigger
	err = json.Unmarshal(body, timerTrigger)
	if err != nil {
		rou.logger.Error(err, err.Error())
		return
	}
	triggerFound := &batchv1beta1.TimerTrigger{}
	err = rou.kubeclient.Get(context.TODO(), types.NamespacedName{Namespace: timerTrigger.Namespace, Name: timerTrigger.Name}, triggerFound)
	if err != nil {
		if !errors.IsNotFound(err) {
			rou.logger.Error(err, err.Error())
			return
		}
	} else if triggerFound != nil {
		rou.logger.Error(fmt.Errorf("%s", "HTTP Trigger already existed."), "Existed")
		return
	}

	err = rou.kubeclient.Create(context.TODO(), &timerTrigger)
	if err != nil {
		rou.logger.Error(err, err.Error())
		return
	}

	resp, err := json.Marshal(timerTrigger)
	if err != nil {
		rou.logger.Error(err, err.Error())
		return
	}
	rou.respondWithSuccess(w, resp)
}

func (rou *Router) TimerTriggerApiUpdate(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		rou.logger.Error(err, err.Error())
		return
	}
	var timerTrigger batchv1beta1.TimerTrigger
	err = json.Unmarshal(body, &timerTrigger)
	if err != nil {
		rou.logger.Error(err, err.Error())
		return
	}
	triggerFound := &batchv1beta1.TimerTrigger{}
	err = rou.kubeclient.Get(context.TODO(), types.NamespacedName{Namespace: timerTrigger.Namespace, Name: timerTrigger.Name}, triggerFound)
	if err != nil {
		rou.logger.Error(err, err.Error())
		return
	}
	err = rou.kubeclient.Update(context.TODO(), &timerTrigger)
	if err != nil {
		rou.logger.Error(err, err.Error())
		return
	}

	resp, err := json.Marshal(timerTrigger)
	if err != nil {
		rou.logger.Error(err, err.Error())
		return
	}
	rou.respondWithSuccess(w, resp)

}

func (rou *Router) TimerTriggerApiDelete(w http.ResponseWriter, r *http.Request) {
	timerTrigger, err := rou.findHttptrigger(r)
	if err != nil {
		rou.logger.Error(err, err.Error())
		return
	}
	err = rou.kubeclient.Delete(context.TODO(), timerTrigger)
	if err != nil {
		rou.logger.Error(err, err.Error())
		return
	}

	resp, err := json.Marshal(timerTrigger)
	rou.respondWithSuccess(w, resp)
}
