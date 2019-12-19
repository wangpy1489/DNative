package router

import (
	"net/http"
	"fmt"
	"log"
	"github.com/gorilla/mux"
)

type(
	Router struct{
		logger *log.Logger
		info string
	}
)

func MakeRouter(in_logger *log.Logger) (*Router,error) {
	return &Router{
		logger: in_logger,
		info: "trigger router",
	}, nil
}

func (rou *Router) GetHandler() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/", rou.HomeHandler).Methods("GET")
	return r
}

func (rou *Router) Serve(port int){
	// r := mux.NewRouter()
	address := fmt.Sprintf(":%v", port)
	err := http.ListenAndServe(address, rou.GetHandler())
	rou.logger.Fatal("done listening", err)
}