package repositories

import "github.com/brianmorais/go-book-store-api/src/domain/entities"

type BookRepository struct{}

func NewBookRepository() *BookRepository {
	return &BookRepository{}
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
