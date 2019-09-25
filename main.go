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

	// data
	data.Books = append(data.Books, model.Book{ID: "1", Isbn: "438227", Title: "Book One", Author: &model.Author{Firstname: "John", Lastname: "Doe"}})
	data.Books = append(data.Books, model.Book{ID: "2", Isbn: "454555", Title: "Book Two", Author: &model.Author{Firstname: "Steve", Lastname: "Smith"}})

	// initializing router
	r := chi.NewRouter()
	r.Mount("/index", indexRouter())

	// starting server
	log.Println("Server Running at port: 8080")
	http.ListenAndServe(":8080", r)
}
