package services

import (
	"errors"
	"freshers-bootcamp-last/models"
	"freshers-bootcamp-last/repositories"
)

type ProductService struct {
	repo *repositories.ProductRepository
}

func NewProductService(repo *repositories.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) CreateProduct(product *models.Product) error {
	return s.repo.Create(product)
}

func (s *ProductService) UpdateProduct(product *models.Product) error {
	return s.repo.Update(product)
}

func (s *ProductService) PatchUpdateProduct(id uint, updates map[string]interface{}) (*models.Product, error) {
	return s.repo.PatchUpdate(id, updates)
}

func (s *ProductService) GetAllProducts() ([]models.Product, error) {
	return s.repo.FindAll()
}

func (s *ProductService) GetProductById(id uint) (*models.Product, error) {
	return s.repo.FindByID(id)
}

func (s *ProductService) GetProductByID(id uint) (*models.Product, error) {
	return s.repo.FindByID(id)
}

func (s *ProductService) DeleteProduct(id uint) error {
	// Check if the product exists
	product, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("product not found")
	}

	// Delete the product
	return s.repo.Delete(product)
}
