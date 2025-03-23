package repositories

import (
	"freshers-bootcamp-last/models"
	"sync"

	"gorm.io/gorm"
)

type ProductRepository struct {
	db    *gorm.DB
	mutex sync.Mutex
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) Create(product *models.Product) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	return r.db.Create(product).Error
}

func (r *ProductRepository) Update(product *models.Product) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	return r.db.Save(product).Error
}

func (r *ProductRepository) FindAll() ([]models.Product, error) {
	var products []models.Product
	err := r.db.Find(&products).Error
	return products, err
}

func (r *ProductRepository) FindByID(id uint) (*models.Product, error) {
	var product models.Product
	err := r.db.First(&product, id).Error
	return &product, err
}

func (r *ProductRepository) Delete(product *models.Product) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	return r.db.Delete(product).Error
}
