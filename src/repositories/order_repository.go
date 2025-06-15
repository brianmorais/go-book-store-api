package repositories

import "github.com/brianmorais/go-book-store-api/src/domain/entities"

type orderRepository struct{}

func NewOrderRepository() *orderRepository {
	return &orderRepository{}
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
