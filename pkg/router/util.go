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

func (rou *Router) respondWithError(w http.ResponseWriter, in_err error) error {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	_, err := w.Write([]byte(in_err.Error()))
	if err != nil {
		return err
	}
	return nil
}
