package api

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/nurislam03/Bookshop/model"
)

// Get all books
func (api *API) getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	books, err := api.BookDataStore.GetAllBooks()
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(404), 404)
		return
	}
	json.NewEncoder(w).Encode(books)
}

// Get single book
func (api *API) getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	bID := chi.URLParam(r, "id")
	log.Println("BOOK-ID:", bID)

	book, err := api.BookDataStore.GetBookByID(bID)
	if err != nil {
		http.Error(w, http.StatusText(404), 404)
		return
	}
	json.NewEncoder(w).Encode(book)
}

type createBookBody struct {
	Isbn   string `json:"isbn"`
	Title  string `json:"title"`
	Author struct {
		Firstname string `json:"firstname"`
		Lastname  string `json:"lastname"`
	} `json:"author"`
}

// Add new book
func (api *API) createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	body := &createBookBody{}
	_ = json.NewDecoder(r.Body).Decode(body)

	book := &model.Book{
		ID:    strconv.Itoa(rand.Intn(100000000)),
		Isbn:  body.Isbn,
		Title: body.Title,
		Author: &model.Author{
			Firstname: body.Author.Firstname,
			Lastname:  body.Author.Lastname,
		},
	}

	book, err := api.BookDataStore.AddNewBook(book)
	if err != nil {
		http.Error(w, http.StatusText(404), 404)
		return
	}
	json.NewEncoder(w).Encode(book)
}

type updateBookBody struct {
	Isbn   string `json:"isbn"`
	Title  string `json:"title"`
	Author struct {
		Firstname string `json:"firstname"`
		Lastname  string `json:"lastname"`
	} `json:"author"`
}

// Update book
func (api *API) updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	body := &updateBookBody{}
	_ = json.NewDecoder(r.Body).Decode(body)

	bID := chi.URLParam(r, "id")
	// log.Println(bID)

	book := &model.Book{
		ID:    bID,
		Isbn:  body.Isbn,
		Title: body.Title,
		Author: &model.Author{
			Firstname: body.Author.Firstname,
			Lastname:  body.Author.Lastname,
		},
	}

	book, err := api.BookDataStore.UpdateBookByID(bID, book)

	if err != nil {
		http.Error(w, http.StatusText(404), 404)
		return
	}
	json.NewEncoder(w).Encode(book)
}

// Delete book
func (api *API) deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	bID := chi.URLParam(r, "id")

	err := api.BookDataStore.DeleteBookByID(bID)
	if err != nil {
		http.Error(w, http.StatusText(404), 404)
		return
	}
	json.NewEncoder(w).Encode("{status: success}")
}
