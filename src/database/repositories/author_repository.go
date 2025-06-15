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
	var authors []entities.Author
	result := r.dbConnection.Connection.Find(&authors)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return []entities.Author{}, nil
	}
	return authors, nil
}

func (r *AuthorRepository) AddAuthor(author entities.Author) (entities.Author, error) {
	result := r.dbConnection.Connection.Create(&author)
	if result.Error != nil {
		return entities.Author{}, result.Error
	}
	return author, nil
}
