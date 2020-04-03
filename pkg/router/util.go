package router

import "net/http"

func (rou *Router) respondWithSuccess(w http.ResponseWriter, resp []byte) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	_, err := w.Write(resp)
	if err != nil {
		return err
	}
	return nil
}
