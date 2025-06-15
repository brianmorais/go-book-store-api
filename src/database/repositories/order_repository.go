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
	result := r.dbConnection.Connection.Create(&order)
	if result.Error != nil {
		return entities.Order{}, result.Error
	}
	return order, nil
}

func (r *orderRepository) GetOrdersByUserId(userId int) ([]entities.Order, error) {
	var orders []entities.Order
	result := r.dbConnection.Connection.Where("user_id = ?", userId).Find(&orders)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return []entities.Order{}, nil
	}
	return orders, nil
}

func (r *orderRepository) UpdateOrderStatus(orderId int, status string) error {
	result := r.dbConnection.Connection.Model(&entities.Order{}).Where("id = ?", orderId).Update("status", status)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return nil
	}
	return nil
}
