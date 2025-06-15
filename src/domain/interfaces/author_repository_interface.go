package interfaces

import entities "github.com/brianmorais/go-book-store-api/src/domain/entities"

type IAuthorRepository interface {
	GetAllAuthors() ([]entities.Author, error)
	AddAuthor(author entities.Author) (entities.Author, error)
}
