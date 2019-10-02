package memory

import (
	"errors"

	"github.com/nurislam03/Bookshop/model"
)

// BookDataStore ...
type BookDataStore struct {
	books []*model.Book
}

// GetAllBooks ...
func (data *BookDataStore) GetAllBooks() ([]*model.Book, error) {
	if len(data.books) == 0 {
		return nil, errors.New("no book has been added")
	}
	return data.books, nil
}

// GetBookByID ...
func (data *BookDataStore) GetBookByID(id string) (*model.Book, error) {
	// Loop through books and find one with the id from the params
	for _, item := range data.books {
		if item.ID == id {
			return item, nil
		}
	}
	return nil, errors.New("not found")
}

// UpdateBookByID ...
func (data *BookDataStore) UpdateBookByID(id string, book *model.Book) (*model.Book, error) {
	// Loop through books and find one with the id from the params
	for index, item := range data.books {
		if item.ID == id {
			book.ID = id
			data.books[index] = book
			return book, nil
		}
	}
	return nil, errors.New("not found")
}

// DeleteBookByID ...
func (data *BookDataStore) DeleteBookByID(id string) ([]*model.Book, error) {
	for index, item := range data.books {
		if item.ID == id {
			data.books = append(data.books[:index], data.books[index+1:]...)
			return data.books, nil
		}
	}
	return nil, errors.New("not found")
}

// AddNewBook ...
func (data *BookDataStore) AddNewBook(book *model.Book) (*model.Book, error) {
	data.books = append(data.books, book)
	return book, nil
}
