package repositories

import (
	"github.com/brianmorais/go-book-store-api/src/database"
	"github.com/brianmorais/go-book-store-api/src/domain/entities"
)

type orderRepository struct {
	dbConnection *database.DbConnection
}

func NewOrderRepository(conection database.DbConnection) *orderRepository {
	return &orderRepository{
		dbConnection: &conection,
	}
}

func (r *orderRepository) AddOrder(order entities.Order) (entities.Order, error) {
	return entities.Order{}, nil
}

func (r *orderRepository) GetOrdersByUserId(userId int) ([]entities.Order, error) {
	return []entities.Order{}, nil
}

func (r *orderRepository) UpdateOrderStatus(orderId int, status string) error {
	return nil
}
