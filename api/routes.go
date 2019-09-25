package api

import (
	"net/http"

	"github.com/go-chi/chi"
)

func indexRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/", indexGet)
	r.Post("/", indexPost)
	return r
}

func booksRouter() http.Handler {
	r := chi.NewRouter()

	r.Get("/", getBooks)
	r.Get("/{id}", getBook)
	r.Post("/", createBook)
	r.Put("/{id}", updateBook)
	r.Delete("/{id}", deleteBook)
	return r
}

// Router ...
func Router() chi.Router {
	// initializing router
	r := chi.NewRouter()
	r.Mount("/index", indexRouter())
	r.Mount("/books", booksRouter())
	return r
}
