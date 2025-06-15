package repositories

import "github.com/brianmorais/go-book-store-api/src/domain/entities"

type AuthorRepository struct {
}

func NewAuthorRepository() *AuthorRepository {
	return &AuthorRepository{}
}

func (r *AuthorRepository) GetAllAuthors() ([]entities.Author, error) {
	return []entities.Author{}, nil
}

func (r *AuthorRepository) AddAuthor(author entities.Author) (entities.Author, error) {
	return entities.Author{}, nil
}
