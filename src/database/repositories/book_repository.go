package repositories

import (
	"github.com/brianmorais/go-book-store-api/src/database"
	"github.com/brianmorais/go-book-store-api/src/domain/entities"
)

type BookRepository struct {
	dbConnection *database.DbConnection
}

func NewBookRepository(connection database.DbConnection) *BookRepository {
	return &BookRepository{
		dbConnection: &connection,
	}
}

func (b *BookRepository) GetAllBooks() ([]entities.Book, error) {
	return []entities.Book{}, nil
}

func (b *BookRepository) AddBook(book entities.Book) (entities.Book, error) {
	return entities.Book{}, nil
}

func (b *BookRepository) UpdateBookById(id int) error {
	return nil
}

func (b *BookRepository) DeleteBookById(id int) error {
	return nil
}
