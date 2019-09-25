package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func indexGet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Getting Index route!"))
}

func indexPost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Getting POST index route!"))
}

func indexRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/", indexGet)
	r.Post("/", indexPost)
	return r
}

func main() {
	r := chi.NewRouter()
	r.Mount("/index", indexRouter())

	log.Println("Server Running at port: 8080")
	http.ListenAndServe(":8080", r)
}
