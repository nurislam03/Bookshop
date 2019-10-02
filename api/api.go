package api

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/nurislam03/Bookshop/data"
)

// API ...
type API struct {
	router        chi.Router
	BookDataStore data.BookDataStore
}

// GetRouter ...
func (api *API) GetRouter() http.Handler {
	return api.router
}

// RegisterRoutes ...
func (api *API) RegisterRoutes() {
	r := chi.NewRouter()
	r.Mount("/index", api.indexRouter())
	r.Mount("/books", api.booksRouter())
	api.router = r
}

func (api *API) indexRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/", api.indexGet)
	r.Post("/", api.indexPost)
	return r
}

func (api *API) booksRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/", api.getBooks)
	r.Get("/{id}", api.getBook)
	r.Post("/", api.createBook)
	r.Put("/{id}", api.updateBook)
	r.Delete("/{id}", api.deleteBook)
	return r
}
