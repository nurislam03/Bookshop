package data

import (
	"errors"
	"math/rand"
	"strconv"

	"github.com/nurislam03/Bookshop/model"
)

// Init books var as a slice Book struct
var books []*model.Book

// GetAllBooks ...
func GetAllBooks() ([]*model.Book, error) {
	if len(books) == 0 {
		return nil, errors.New("no book has been added")
	}
	return books, nil
}

// GetBookByID ...
func GetBookByID(id string) (*model.Book, error) {
	// Loop through books and find one with the id from the params
	for _, item := range books {
		if item.ID == id {
			return item, nil
		}
	}
	return nil, errors.New("not found")
}

// UpdateBookByID ...
func UpdateBookByID(id string, book *model.Book) (*model.Book, error) {
	// Loop through books and find one with the id from the params
	for index, item := range books {
		if item.ID == id {
			book.ID = id
			books[index] = book
			return book, nil
		}
	}
	return nil, errors.New("not found")
}

// DeleteBookByID ...
func DeleteBookByID(id string) ([]*model.Book, error) {
	for index, item := range books {
		if item.ID == id {
			books = append(books[:index], books[index+1:]...)
			return books, nil
		}
	}
	return nil, errors.New("not found")
}

// AddNewBook ...
func AddNewBook(book *model.Book) (*model.Book, error) {
	book.ID = strconv.Itoa(rand.Intn(100000000)) // Mock ID - not safe
	books = append(books, book)
	return book, nil
}
