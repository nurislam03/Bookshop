package api

import "net/http"

func (api *API) indexGet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Getting Index route!"))
}

func (api *API) indexPost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Getting POST index route!"))
}
