package repositories

import (
	"freshers-bootcamp-last/models"
	"sync"
	"time"

	"gorm.io/gorm"
)

type OrderRepository struct {
	db    *gorm.DB
	mutex sync.Mutex
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) Create(order *models.Order) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	return r.db.Create(order).Error
}

func (r *OrderRepository) FindByCustomerID(customerID uint) ([]models.Order, error) {
	var orders []models.Order
	err := r.db.Where("customer_id = ?", customerID).Find(&orders).Error
	return orders, err
}

func (r *OrderRepository) LastOrderTime(customerID uint) (time.Time, error) {
	var order models.Order
	err := r.db.Where("customer_id = ?", customerID).Order("created_at desc").First(&order).Error
	return order.CreatedAt, err
}

func (r *OrderRepository) FindByID(id uint) (*models.Order, error) {
	var order models.Order
	err := r.db.First(&order, id).Error
	return &order, err
}
