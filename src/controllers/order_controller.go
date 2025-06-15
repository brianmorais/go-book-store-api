package controllers

import (
	"strconv"

	"github.com/brianmorais/go-book-store-api/src/domain/entities"
	"github.com/brianmorais/go-book-store-api/src/domain/interfaces"
	"github.com/labstack/echo/v4"
)

type OrderController struct {
	orderRepository interfaces.IOrderRepository
}

func NewOrderController(orderRepository interfaces.IOrderRepository) *OrderController {
	return &OrderController{
		orderRepository: orderRepository,
	}
}

func (o *OrderController) AddOrder(c echo.Context) error {
	var order entities.Order
	if err := c.Bind(&order); err != nil {
		return c.JSON(400, map[string]string{"error": "Invalid order data"})
	}

	createdOrder, err := o.orderRepository.AddOrder(order)
	if err != nil {
		return c.JSON(500, map[string]string{"error": "Failed to add order"})
	}
	return c.JSON(201, createdOrder)
}

func (o *OrderController) GetOrdersByUserId(c echo.Context) error {
	userId := c.Param("userId")
	if userId == "" {
		return c.JSON(400, map[string]string{"error": "User ID is required"})
	}

	orderInt, err := strconv.Atoi(userId)

	if err != nil {
		return c.JSON(400, map[string]string{"error": "Invalid user ID"})
	}

	orders, err := o.orderRepository.GetOrdersByUserId(orderInt)
	if err != nil {
		return c.JSON(500, map[string]string{"error": "Failed to retrieve orders"})
	}
	return c.JSON(200, orders)
}

func (o *OrderController) UpdateOrderStatus(c echo.Context) error {
	orderId := c.Param("orderId")
	status := c.Param("status")
	if orderId == "" || status == "" {
		return c.JSON(400, map[string]string{"error": "Order ID and status is required"})
	}

	orderInt, err := strconv.Atoi(orderId)

	if err != nil {
		return c.JSON(400, map[string]string{"error": "Invalid order ID"})
	}

	err = o.orderRepository.UpdateOrderStatus(orderInt, status)
	if err != nil {
		return c.JSON(500, map[string]string{"error": "Failed to update order status"})
	}
	return c.NoContent(204)
}
