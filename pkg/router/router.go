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
	r.HandleFunc("/v1", rou.CreateHttpTrigger).Methods("POST")
	r.HandleFunc("/v1/{httpTrigger}", rou.HttpTrigger).Methods("POST")
	r.HandleFunc("/v1/{httpTrigger}", rou.HttpTrigger).Methods("GET")
	r.HandleFunc("/v1/{httpTrigger}", rou.UpdateHttpTrigger).Methods("PUT")
	r.HandleFunc("/v1/{httpTrigger}", rou.DeleteHttpTrigger).Methods("DELETE")
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
