package product

import (
	"github.com/JenswBE/go-commerce/entity"
)

// GetProduct fetches a single product by ID
func (s *Service) GetProduct(id entity.ID) (*entity.Product, error) {
	return s.db.GetProduct(id)
}

// ListProducts fetches all products
func (s *Service) ListProducts() ([]*entity.Product, error) {
	return s.db.ListProducts()
}

// CreateProduct creates a new product
func (s *Service) CreateProduct(product *entity.Product) (*entity.Product, error) {
	// Generate new ID
	product.ID = entity.NewID()

	// Validate entity
	err := product.Validate()
	if err != nil {
		return nil, err
	}

	// Persist entity
	return s.db.CreateProduct(product)
}

// UpdateProduct persists the provided product
func (s *Service) UpdateProduct(product *entity.Product) (*entity.Product, error) {
	// Validate entity
	err := product.Validate()
	if err != nil {
		return nil, err
	}

	// Persist entity
	return s.db.UpdateProduct(product)
}

// DeleteProduct deletes a single product by ID
func (s *Service) DeleteProduct(id entity.ID) error {
	return s.db.DeleteProduct(id)
}
