package product

import (
	"github.com/JenswBE/go-commerce/entity"
)

// GetProduct fetches a single product by ID
func (s *Service) GetProduct(id entity.ID) (*entity.Product, error) {
	return s.repo.GetProduct(id)
}

// ListProducts fetches all products
func (s *Service) ListProducts() ([]*entity.Product, error) {
	return s.repo.ListProducts()
}

// CreateProduct creates a new product
func (s *Service) CreateProduct(name string, price int) (*entity.Product, error) {
	// Create entity
	m, err := entity.NewProduct(name, price)
	if err != nil {
		return nil, err
	}

	// Persist entity
	return s.repo.CreateProduct(m)
}

// UpdateProduct persists the provided product
func (s *Service) UpdateProduct(e *entity.Product) (*entity.Product, error) {
	// Validate entity
	err := e.Validate()
	if err != nil {
		return nil, err
	}

	// Persist entity
	return s.repo.UpdateProduct(e)
}

// DeleteProduct deletes a single product by ID
func (s *Service) DeleteProduct(id entity.ID) error {
	return s.repo.DeleteProduct(id)
}
