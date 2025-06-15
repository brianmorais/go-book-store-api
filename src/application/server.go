package application

import (
	"fmt"
	"os"

	"github.com/brianmorais/go-book-store-api/src/controllers"
	"github.com/brianmorais/go-book-store-api/src/database"
	"github.com/brianmorais/go-book-store-api/src/database/repositories"
	"github.com/labstack/echo/v4"
)

type WebServer struct {
	bookController   *controllers.BookController
	authorController *controllers.AuthorController
	orderController  *controllers.OrderController
}

func NewWebServer() *WebServer {
	return &WebServer{}
}

func (ws *WebServer) Serve() {
	ws.setupControllers()
	e := echo.New()
	ws.setupRoutes(e)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", os.Getenv("API_PORT"))))
}

func (ws *WebServer) setupControllers() {
	databaseConnection := database.NewDbConnection()
	ws.bookController = controllers.NewBookController(repositories.NewBookRepository(*databaseConnection))
	ws.authorController = controllers.NewAuthorController(repositories.NewAuthorRepository(*databaseConnection))
	ws.orderController = controllers.NewOrderController(repositories.NewOrderRepository(*databaseConnection))
}

func (ws *WebServer) setupRoutes(e *echo.Echo) {
	e.GET("/books", ws.bookController.GetAllBooks)
	e.POST("/books", ws.bookController.AddBook)
	e.PUT("/books/:id", ws.bookController.UpdateBookById)
	e.DELETE("/books/:id", ws.bookController.DeleteBookById)
	e.GET("/authors", ws.authorController.GetAllAuthors)
	e.POST("/authors", ws.authorController.AddAuthor)
	e.GET("/orders/:userId", ws.orderController.GetOrdersByUserId)
	e.POST("/orders", ws.orderController.AddOrder)
	e.PUT("/orders/:orderId/:status", ws.orderController.UpdateOrderStatus)
}
