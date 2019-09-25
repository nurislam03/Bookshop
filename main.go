package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

// Book struct (Model)
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Author struct
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// Init books var as a slice Book struct
var books []Book

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

	// Hardcoded data - @todo: add database
	books = append(books, Book{ID: "1", Isbn: "438227", Title: "Book One", Author: &Author{Firstname: "John", Lastname: "Doe"}})
	books = append(books, Book{ID: "2", Isbn: "454555", Title: "Book Two", Author: &Author{Firstname: "Steve", Lastname: "Smith"}})

	// initializing router
	r := chi.NewRouter()
	r.Mount("/index", indexRouter())

	// starting server
	log.Println("Server Running at port: 8080")
	http.ListenAndServe(":8080", r)
}
