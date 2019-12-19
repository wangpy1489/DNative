package router

import (
	// "fmt"
	"net/http"
)

func (rou *Router) HomeHandler(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write([]byte(rou.info))
}

