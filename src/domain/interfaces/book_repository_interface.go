package interfaces

import entities "github.com/brianmorais/go-book-store-api/src/domain/entities"

type IBookRepository interface {
	GetAllBooks() ([]entities.Book, error)
	AddBook(book entities.Book) (entities.Book, error)
	UpdateBookById(id int, book entities.Book) error
	DeleteBookById(id int) error
}
