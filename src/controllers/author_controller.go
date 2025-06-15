package controllers

import (
	"github.com/brianmorais/go-book-store-api/src/domain/entities"
	"github.com/brianmorais/go-book-store-api/src/domain/interfaces"
	"github.com/labstack/echo/v4"
)

type AuthorController struct {
	authorRepository interfaces.IAuthorRepository
}

func NewAuthorController(authorRepository interfaces.IAuthorRepository) *AuthorController {
	return &AuthorController{
		authorRepository: authorRepository,
	}
}

func (a *AuthorController) GetAllAuthors(c echo.Context) error {
	authors, err := a.authorRepository.GetAllAuthors()
	if err != nil {
		return c.JSON(500, map[string]string{"error": "Failed to retrieve authors"})
	}
	return c.JSON(200, authors)
}

func (a *AuthorController) AddAuthor(c echo.Context) error {
	var author entities.Author
	if err := c.Bind(&author); err != nil {
		return c.JSON(400, map[string]string{"error": "Invalid author data"})
	}

	createdAuthor, err := a.authorRepository.AddAuthor(author)
	if err != nil {
		return c.JSON(500, map[string]string{"error": "Failed to add author"})
	}
	return c.JSON(201, createdAuthor)
}
