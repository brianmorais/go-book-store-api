package interfaces

import entities "github.com/brianmorais/go-book-store-api/src/domain/entities"

type IOrderRepository interface {
	AddOrder(order entities.Order) (entities.Order, error)
	GetOrdersByUserId(userId int) ([]entities.Order, error)
	UpdateOrderStatus(orderId int, status string) error
}
