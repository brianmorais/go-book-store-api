package repositories

import (
	"github.com/brianmorais/go-book-store-api/src/database"
	"github.com/brianmorais/go-book-store-api/src/domain/entities"
)

type AuthorRepository struct {
	dbConnection *database.DbConnection
}

func NewAuthorRepository(connection database.DbConnection) *AuthorRepository {
	return &AuthorRepository{
		dbConnection: &connection,
	}
}

func (r *AuthorRepository) GetAllAuthors() ([]entities.Author, error) {
	return []entities.Author{}, nil
}

func (r *AuthorRepository) AddAuthor(author entities.Author) (entities.Author, error) {
	return entities.Author{}, nil
}
