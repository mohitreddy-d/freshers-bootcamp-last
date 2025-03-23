package services

import (
	"errors"
	"freshers-bootcamp-last/models"
	"freshers-bootcamp-last/repositories"
	"time"
)

type OrderService struct {
	orderRepo *repositories.OrderRepository
	prodRepo  *repositories.ProductRepository
}

func NewOrderService(orderRepo *repositories.OrderRepository, prodRepo *repositories.ProductRepository) *OrderService {
	return &OrderService{orderRepo: orderRepo, prodRepo: prodRepo}
}

func (s *OrderService) PlaceOrder(order *models.Order) error {
	// Check cool-down period
	lastOrderTime, err := s.orderRepo.LastOrderTime(order.CustomerID)
	if err == nil && time.Since(lastOrderTime) < 5*time.Minute {
		return errors.New("cool-down period not over")
	}

	// Check product availability
	product, err := s.prodRepo.FindByID(order.ProductID)
	if err != nil {
		return err
	}
	if product.Quantity < order.Quantity {
		return errors.New("insufficient stock")
	}

	// Update product quantity
	product.Quantity -= order.Quantity
	if err := s.prodRepo.Update(product); err != nil {
		return err
	}

	// Create order
	order.Status = "order placed"
	return s.orderRepo.Create(order)
}

func (s *OrderService) GetOrderHistory(customerID uint) ([]models.Order, error) {
	return s.orderRepo.FindByCustomerID(customerID)
}

func (s *OrderService) GetOrderDetails(orderID uint) (*models.Order, error) {
	// Fetch the order from the repository
	order, err := s.orderRepo.FindByID(orderID)
	if err != nil {
		return nil, errors.New("order not found")
	}

	return order, nil
}
