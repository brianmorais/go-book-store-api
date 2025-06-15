package controllers

import (
	"strconv"

	"github.com/brianmorais/go-book-store-api/src/domain/entities"
	"github.com/brianmorais/go-book-store-api/src/domain/interfaces"
	"github.com/labstack/echo/v4"
)

type BookController struct {
	bookRepository interfaces.IBookRepository
}

func NewBookController(bookRepository interfaces.IBookRepository) *BookController {
	return &BookController{
		bookRepository: bookRepository,
	}
}

func (b *BookController) GetAllBooks(c echo.Context) error {
	books, err := b.bookRepository.GetAllBooks()
	if err != nil {
		return c.JSON(500, map[string]string{"error": "Failed to retrieve books"})
	}
	return c.JSON(200, books)
}

func (b *BookController) AddBook(c echo.Context) error {
	var book entities.Book
	if err := c.Bind(&book); err != nil {
		return c.JSON(400, map[string]string{"error": "Invalid book data"})
	}

	createdBook, err := b.bookRepository.AddBook(book)
	if err != nil {
		return c.JSON(500, map[string]string{"error": "Failed to add book"})
	}
	return c.JSON(201, createdBook)
}

func (b *BookController) UpdateBookById(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(400, map[string]string{"error": "Book ID is required"})
	}

	var book entities.Book
	if err := c.Bind(&book); err != nil {
		return c.JSON(400, map[string]string{"error": "Invalid book data"})
	}

	intId, err := strconv.Atoi(id)

	if err != nil {
		return c.JSON(400, map[string]string{"error": "Invalid book ID"})
	}

	err = b.bookRepository.UpdateBookById(intId, book)
	if err != nil {
		return c.JSON(500, map[string]string{"error": "Failed to update book"})
	}
	return c.JSON(200, map[string]string{"message": "Book updated successfully"})
}

func (b *BookController) DeleteBookById(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(400, map[string]string{"error": "Book ID is required"})
	}

	intId, err := strconv.Atoi(id)

	if err != nil {
		return c.JSON(400, map[string]string{"error": "Invalid book ID"})
	}

	err = b.bookRepository.DeleteBookById(intId)
	if err != nil {
		return c.JSON(500, map[string]string{"error": "Failed to delete book"})
	}
	return c.JSON(200, map[string]string{"message": "Book deleted successfully"})
}
