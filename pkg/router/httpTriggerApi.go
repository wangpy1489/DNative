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
	"github.com/wangpy1489/DNative/pkg/storage"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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

func (rou *Router) submitBatchJob(template *batchv1beta1.BatchTemplate) (*batchv1beta1.BatchJob, error) {
	temp := template.Spec.DeepCopy()
	if len(template.Spec.StorageName) != 0 {
		sc := storage.MakeSotreCore(rou.kubeclient)
		volume, err := sc.VolumeBuilder(*template, "test")
		if err != nil {
			return nil, err
		}
		switch temp.Type {
		case batchv1beta1.Batch:
			temp.Template.Batch.Spec.Volumes = append(temp.Template.Batch.Spec.Volumes, *volume)
		case batchv1beta1.Spark:
			temp.Template.Spark.Volumes = append(temp.Template.Spark.Volumes, *volume)
		}
	}
	newApp := &batchv1beta1.BatchJob{
		ObjectMeta: metav1.ObjectMeta{
			Name:      template.Name,
			Namespace: template.Namespace,
		},
		Spec: batchv1beta1.BatchJobSpec{
			Type:     template.Spec.Type,
			Template: temp.Template,
		},
	}
	newApp.Spec.Template = template.Spec.Template
	err := rou.kubeclient.Create(context.TODO(), newApp)
	if err != nil {
		return nil, err
	}
	// newApp, err := r.sparkClient.Sparkoperator().SparkApplications(job.Namespace).Create(newApp)

	return newApp, nil
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

func (rou *Router) HttpTriggerApi(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["httpTrigger"]
	namespace := r.URL.Query().Get("namespace")
	httpTrigger := &batchv1beta1.HttpTrigger{}
	rou.logger.Info(namespace, name)
	err := rou.kubeclient.Get(context.TODO(), types.NamespacedName{Namespace: namespace, Name: name}, httpTrigger)
	if err != nil {
		rou.logger.Error(err, err.Error())
		rou.respondWithError(w, err)
		return
	}
	batchTemplate := &batchv1beta1.BatchTemplate{}
	err = rou.kubeclient.Get(context.TODO(), types.NamespacedName{Namespace: httpTrigger.Namespace, Name: httpTrigger.Spec.JobReference.Name}, batchTemplate)
	if err != nil {
		rou.logger.Error(err, err.Error())
		rou.respondWithError(w, err)
		return
	}
	batchJob, err := rou.submitBatchJob(batchTemplate)
	if err != nil {
		rou.logger.Error(err, err.Error())
		rou.respondWithError(w, err)
		return
	}
	resp, err := json.Marshal(batchJob)
	if err != nil {
		rou.logger.Error(err, err.Error())
		rou.respondWithError(w, err)
		return
	}
	rou.respondWithSuccess(w, resp)
}

func (rou *Router) HttpTriggerApiUpdate(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		rou.logger.Error(err, err.Error())
		rou.respondWithError(w, err)
		return
	}
	httptrigger := batchv1beta1.HttpTrigger{}
	err = json.Unmarshal(body, &httptrigger)
	if err != nil {
		rou.logger.Error(err, err.Error())
		rou.respondWithError(w, err)
		return
	}
	triggerFound := &batchv1beta1.HttpTrigger{}
	err = rou.kubeclient.Get(context.TODO(), types.NamespacedName{Namespace: httptrigger.Namespace, Name: httptrigger.Name}, triggerFound)
	if err != nil {
		rou.logger.Error(err, err.Error())
		rou.respondWithError(w, err)
		return
	}
	triggerFound.Spec = httptrigger.Spec
	err = rou.kubeclient.Update(context.TODO(), triggerFound)
	if err != nil {
		rou.logger.Error(err, err.Error())
		rou.respondWithError(w, err)
		return
	}

	resp, err := json.Marshal(httptrigger)
	if err != nil {
		rou.logger.Error(err, err.Error())
		rou.respondWithError(w, err)
		return
	}
	rou.respondWithSuccess(w, resp)

}

func (rou *Router) HttpTriggerApiCreate(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		rou.logger.Error(err, err.Error())
		rou.respondWithError(w, err)
		return
	}
	httptrigger := batchv1beta1.HttpTrigger{}
	err = json.Unmarshal(body, &httptrigger)
	if err != nil {
		rou.logger.Error(err, err.Error())
		rou.respondWithError(w, err)
		return
	}
	triggerFound := &batchv1beta1.HttpTrigger{}
	err = rou.kubeclient.Get(context.TODO(), types.NamespacedName{Namespace: httptrigger.Namespace, Name: httptrigger.Name}, triggerFound)
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

	err = rou.kubeclient.Create(context.TODO(), &httptrigger)
	if err != nil {
		rou.logger.Error(err, err.Error())
		rou.respondWithError(w, err)
		return
	}

	resp, err := json.Marshal(httptrigger)
	if err != nil {
		rou.logger.Error(err, err.Error())
		rou.respondWithError(w, err)
		return
	}
	rou.respondWithSuccess(w, resp)
}

func (rou *Router) HttpTriggerApiDelete(w http.ResponseWriter, r *http.Request) {
	httptrigger, err := rou.findHttptrigger(r)
	if err != nil {
		rou.logger.Error(err, err.Error())
		rou.respondWithError(w, err)
		return
	}
	err = rou.kubeclient.Delete(context.TODO(), httptrigger)
	if err != nil {
		rou.logger.Error(err, err.Error())
		rou.respondWithError(w, err)
		return
	}

	resp, err := json.Marshal(httptrigger)
	rou.respondWithSuccess(w, resp)
}
