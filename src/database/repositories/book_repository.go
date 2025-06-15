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
	var books []entities.Book
	result := b.dbConnection.Connection.Find(&books)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return []entities.Book{}, nil
	}
	return books, nil
}

func (b *BookRepository) AddBook(book entities.Book) (entities.Book, error) {
	result := b.dbConnection.Connection.Create(&book)
	if result.Error != nil {
		return entities.Book{}, result.Error
	}
	return book, nil
}

func (b *BookRepository) UpdateBookById(id int, book entities.Book) error {
	result := b.dbConnection.Connection.Model(&entities.Book{}).Where("id = ?", id).Updates(book)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return nil
	}
	return nil
}

func (b *BookRepository) DeleteBookById(id int) error {
	result := b.dbConnection.Connection.Delete(&entities.Book{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return nil
	}
	return nil
}
