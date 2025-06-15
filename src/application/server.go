package application

import (
	"github.com/brianmorais/go-book-store-api/src/controllers"
	"github.com/brianmorais/go-book-store-api/src/repositories"
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

func (ws *WebServer) SetupControllers() {
	ws.bookController = controllers.NewBookController(repositories.NewBookRepository())
	ws.authorController = controllers.NewAuthorController(repositories.NewAuthorRepository())
	ws.orderController = controllers.NewOrderController(repositories.NewOrderRepository())
}

func (ws *WebServer) Serve() {
	e := echo.New()
	e.GET("/books", ws.bookController.GetAllBooks)
	e.POST("/books", ws.bookController.AddBook)
	e.PUT("/books/:id", ws.bookController.UpdateBookById)
	e.DELETE("/books/:id", ws.bookController.DeleteBookById)
	e.GET("/authors", ws.authorController.GetAllAuthors)
	e.POST("/authors", ws.authorController.AddAuthor)
	e.GET("/orders/:userId", ws.orderController.GetOrdersByUserId)
	e.POST("/orders", ws.orderController.AddOrder)
	e.PUT("/orders/:orderId/:status", ws.orderController.UpdateOrderStatus)
	e.Logger.Fatal(e.Start(":8080"))
}
