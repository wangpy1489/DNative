package router

import (
	"net/http"
	"fmt"
	"log"

	"sigs.k8s.io/controller-runtime/pkg/client"
	"github.com/gorilla/mux"
)

type(
	Router struct{
		logger *log.Logger
		kubeclient client.Client
		info string
	}
)

func MakeRouter(in_logger *log.Logger, kubeclient client.Client) (*Router,error) {
	return &Router{
		logger: in_logger,
		kubeclient: kubeclient,
		info: "trigger router",
	}, nil
}

func (rou *Router) GetHandler() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/", rou.HomeHandler).Methods("GET")
	r.HandleFunc("/v1",rou.CreateHttpTrigger).Methods("POST")
	r.HandleFunc("/v1/{httpTrigger}",rou.HttpTrigger).Methods("POST")
	r.HandleFunc("/v1/{httpTrigger}",rou.HttpTrigger).Methods("GET")
	r.HandleFunc("/v1/{httpTrigger}",rou.UpdateHttpTrigger).Methods("PUT")
	r.HandleFunc("/v1/{httpTrigger}",rou.DeleteHttpTrigger).Methods("DELETE")
	return r
}

func (rou *Router) Serve(port int){
	// r := mux.NewRouter()
	address := fmt.Sprintf(":%v", port)
	err := http.ListenAndServe(address, rou.GetHandler())
	rou.logger.Fatal("done listening", err)
}