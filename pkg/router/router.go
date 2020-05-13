package router

import (
	"fmt"
	"net/http"

	"github.com/go-logr/logr"
	"github.com/gorilla/mux"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type (
	Router struct {
		logger     logr.Logger
		kubeclient client.Client
		info       string
	}
)

func MakeRouter(in_logger logr.Logger, kubeclient client.Client) (*Router, error) {
	return &Router{
		logger:     in_logger,
		kubeclient: kubeclient,
		info:       "trigger router",
	}, nil
}

func (rou *Router) GetHandler() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/", rou.HomeHandler).Methods("GET")
	r.HandleFunc("/v1/triggers/http", rou.HttpTriggerApiList).Methods("GET")
	r.HandleFunc("/v1/triggers/http", rou.HttpTriggerApiCreate).Methods("POST")
	r.HandleFunc("/v1/triggers/http", rou.HttpTriggerApiUpdate).Methods("PUT")
	r.HandleFunc("/v1/triggers/http/{httpTrigger}", rou.HttpTriggerApi).Methods("POST")
	r.HandleFunc("/v1/triggers/http/{httpTrigger}", rou.HttpTriggerApi).Methods("GET")
	r.HandleFunc("/v1/triggers/http/{httpTrigger}", rou.HttpTriggerApiDelete).Methods("DELETE")

	r.HandleFunc("/v1/triggers/timer", rou.TimerTriggerApiList).Methods("Get")
	r.HandleFunc("/v1/triggers/timer", rou.TimerTriggerApiCreate).Methods("POST")
	r.HandleFunc("/v1/triggers/timer", rou.TimerTriggerApiUpdate).Methods("PUT")
	r.HandleFunc("/v1/triggers/timer/{timerTrigger}", rou.TimerTriggerApiGet).Methods("GET")
	r.HandleFunc("/v1/triggers/timer/{timerTrigger}", rou.TimerTriggerApiDelete).Methods("DELETE")

	r.HandleFunc("/v1/storage", rou.StorageSourceApiList).Methods("Get")
	r.HandleFunc("/v1/storage", rou.StorageSourceApiCreate).Methods("POST")
	r.HandleFunc("/v1/storage", rou.StorageSourceApiUpdate).Methods("PUT")
	r.HandleFunc("/v1/storage/{storageSource}", rou.StorageSourceApiGet).Methods("GET")
	r.HandleFunc("/v1/storage/{storageSource}", rou.StorageSourceApiDelete).Methods("DELETE")

	r.HandleFunc("/v1/templates", rou.BatchTemplateApiList).Methods("Get")
	r.HandleFunc("/v1/templates", rou.BatchTemplateApiCreate).Methods("POST")
	r.HandleFunc("/v1/templates", rou.BatchTemplateApiUpdate).Methods("PUT")
	r.HandleFunc("/v1/templates/{batchTemplate}", rou.BatchTemplateApiGet).Methods("GET")
	r.HandleFunc("/v1/templates/{batchTemplate}", rou.BatchTemplateApiDelete).Methods("DELETE")

	r.HandleFunc("/v1/jobs", rou.BatchJobApiList).Methods("Get")
	// r.HandleFunc("/v1/jobs", rou.BatchJobApiCreate).Methods("POST")
	// r.HandleFunc("/v1/jobs", rou.BatchJobApiUpdate).Methods("PUT")
	r.HandleFunc("/v1/jobs/{batchJob}", rou.BatchJobApiGet).Methods("GET")
	r.HandleFunc("/v1/jobs/{batchJob}", rou.BatchJobApiDelete).Methods("DELETE")

	return r
}

func (rou *Router) serve(port int) {
	// r := mux.NewRouter()
	address := fmt.Sprintf(":%v", port)
	err := http.ListenAndServe(address, rou.GetHandler())
	rou.logger.Error(err, "done listening")
}

func Start(port int, in_logger logr.Logger, kubeclient client.Client) {
	logger := in_logger.WithName("Router")
	rou, err := MakeRouter(logger, kubeclient)
	if err != nil {
		in_logger.Error(err, "filed to make router:")
	}
	rou.serve(port)
}
