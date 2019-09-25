package main

import (
	"log"
	"net/http"

	"github.com/nurislam03/Bookshop/api"
)

func main() {

	// Hardcoded data - @todo: add database
	// books = append(books, Book{ID: "1", Isbn: "438227", Title: "Book One", Author: &Author{Firstname: "John", Lastname: "Doe"}})
	// books = append(books, Book{ID: "2", Isbn: "454555", Title: "Book Two", Author: &Author{Firstname: "Steve", Lastname: "Smith"}})

	// starting server
	log.Println("Server Running at port: 8080")
	http.ListenAndServe(":8080", api.Router())
}
