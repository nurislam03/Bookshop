package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/nurislam03/Bookshop/data"
	"github.com/nurislam03/Bookshop/model"
)

// Get all books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	books := data.GetAllBooks()
	json.NewEncoder(w).Encode(books)
}

// Get single book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	bID := chi.URLParam(r, "id")
	log.Println("BOOK-ID:", bID)

	book, err := data.GetBookByID(bID)
	if err != nil {
		http.Error(w, http.StatusText(404), 404)
		return
	}

	json.NewEncoder(w).Encode(book)
}

// Add new book
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	book := &model.Book{}
	_ = json.NewDecoder(r.Body).Decode(book)

	book, err := data.AddNewBook(book)
	if err != nil {
		http.Error(w, http.StatusText(404), 404)
		return
	}
	json.NewEncoder(w).Encode(book)
}

// Update book
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	bID := chi.URLParam(r, "id")

	book := &model.Book{}
	_ = json.NewDecoder(r.Body).Decode(book)

	book, err := data.UpdateBookByID(bID, book)

	if err != nil {
		http.Error(w, http.StatusText(404), 404)
		return
	}
	json.NewEncoder(w).Encode(book)
}

// Delete book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	bID := chi.URLParam(r, "id")

	books, err := data.DeleteBookByID(bID)
	if err != nil {
		http.Error(w, http.StatusText(404), 404)
		return
	}
	json.NewEncoder(w).Encode(books)
}
