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
	if os.Getenv("MIGRATE_TABLES") == "true" {
		databaseConnection.MigrateEntities()
	}
	ws.bookController = controllers.NewBookController(repositories.NewBookRepository(*databaseConnection))
	ws.authorController = controllers.NewAuthorController(repositories.NewAuthorRepository(*databaseConnection))
	ws.orderController = controllers.NewOrderController(repositories.NewOrderRepository(*databaseConnection))
}

func (ws *WebServer) setupRoutes(e *echo.Echo) {
	e.GET("/api/books", ws.bookController.GetAllBooks)
	e.POST("/api/books", ws.bookController.AddBook)
	e.PUT("/api/books/:id", ws.bookController.UpdateBookById)
	e.DELETE("/api/books/:id", ws.bookController.DeleteBookById)
	e.GET("/api/authors", ws.authorController.GetAllAuthors)
	e.POST("/api/authors", ws.authorController.AddAuthor)
	e.GET("/api/orders/:userId", ws.orderController.GetOrdersByUserId)
	e.POST("/api/orders", ws.orderController.AddOrder)
	e.PUT("/api/orders/:orderId/:status", ws.orderController.UpdateOrderStatus)
}
