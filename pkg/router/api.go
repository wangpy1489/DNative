package router

import (
	// "fmt"
	"net/http"
	"context"
	"encoding/json"

	batchv1beta1 "github.com/wangpy1489/DNative/pkg/apis/batch/v1beta1"
	"k8s.io/apimachinery/pkg/types"
	"github.com/gorilla/mux"
	"github.com/wangpy1489/DNative/pkg/router/resource"
)



func (rou *Router) respondWithSuccess(w http.ResponseWriter, resp []byte) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	_, err := w.Write(resp)
	if err != nil {
		return  err
	}
	return nil
}

func (rou *Router) HomeHandler(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write([]byte(rou.info))
}

func (rou *Router) HttpTrigger(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	name := vars["httpTrigger"]
	namespace := r.URL.Query().Get("namespace")
	httpTrigger := &batchv1beta1.HttpTrigger{}
	rou.logger.Println(namespace,name)
	err := rou.kubeclient.Get(context.TODO(), types.NamespacedName{Namespace: namespace, Name: name,}, httpTrigger)
	if err != nil {
		rou.logger.Fatal(err)
		return
	}
	batchJob := &batchv1beta1.BatchJob{}
	err = rou.kubeclient.Get(context.TODO(),types.NamespacedName{Namespace: httpTrigger.Namespace, Name: httpTrigger.Spec.JobReference.Nmae},batchJob)
	if err != nil {
		rou.logger.Fatal(err)
		return
	}
	sparkapp, err := resource.SubmitBatchJob(rou.kubeclient,batchJob)
	if err != nil {
		rou.logger.Fatal(err)
		return
	}
	resp, err := json.Marshal(sparkapp)
	if err != nil {
		rou.logger.Fatal(err)
		return
	}
	rou.respondWithSuccess(w, resp)
}