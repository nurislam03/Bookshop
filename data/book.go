package data

import "github.com/nurislam03/Bookshop/model"

// BookDataStore ...
type BookDataStore interface {
	GetAllBooks() ([]*model.Book, error)
	GetBookByID(id string) (*model.Book, error)
	UpdateBookByID(id string, book *model.Book) (*model.Book, error)
	DeleteBookByID(id string) error
	AddNewBook(book *model.Book) (*model.Book, error)
}
