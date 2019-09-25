package api

import "net/http"

func indexGet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Getting Index route!"))
}

func indexPost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Getting POST index route!"))
}
