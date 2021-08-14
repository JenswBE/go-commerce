package product

import (
	"errors"

	"github.com/JenswBE/go-commerce/entities"
	"github.com/JenswBE/go-commerce/utils/imageproxy"
)

// GetProduct fetches a single product by ID
func (s *Service) GetProduct(id entities.ID, imageConfigs map[string]imageproxy.ImageConfig) (*entities.Product, error) {
	// Fetch product
	product, err := s.db.GetProduct(id)
	if err != nil {
		return nil, err
	}

	// Generate URL's
	if len(imageConfigs) > 0 {
		err := s.setImageURLsFromConfig(product.Images, imageConfigs)
		if err != nil {
			return nil, err
		}
	}

	// Get successful
	return product, nil
}

// ListProducts fetches all products
func (s *Service) ListProducts(imageConfigs map[string]imageproxy.ImageConfig) ([]*entities.Product, error) {
	// Fetch products
	products, err := s.db.ListProducts()
	if err != nil {
		return nil, err
	}

	// Generate URL's
	if len(imageConfigs) > 0 {
		for _, product := range products {
			err := s.setImageURLsFromConfig(product.Images, imageConfigs)
			if err != nil {
				return nil, err
			}
		}
	}

	// Get successful
	return products, nil
}

// CreateProduct creates a new product
func (s *Service) CreateProduct(product *entities.Product) (*entities.Product, error) {
	// Generate new ID
	product.ID = entities.NewID()

	// Validate entity
	err := product.Validate()
	if err != nil {
		return nil, err
	}

	// Persist entity
	return s.db.CreateProduct(product)
}

// UpdateProduct persists the provided product
func (s *Service) UpdateProduct(product *entities.Product) (*entities.Product, error) {
	// Validate entity
	err := product.Validate()
	if err != nil {
		return nil, err
	}

	// Persist entity
	return s.db.UpdateProduct(product)
}

// DeleteProduct deletes a single product by ID
func (s *Service) DeleteProduct(id entities.ID) error {
	// Fetch product
	product, err := s.db.GetProduct(id)
	if err != nil {
		return err
	}

	// Delete all images
	for _, image := range product.Images {
		err := s.deleteImage(image)
		if err != nil {
			return err
		}
	}

	// Delete product
	return s.db.DeleteProduct(id)
}

// AddProductImages adds multiple images to a product
func (s *Service) AddProductImages(productID entities.ID, images map[string][]byte, imageConfigs map[string]imageproxy.ImageConfig) (*entities.Product, error) {
	// Fetch product
	product, err := s.db.GetProduct(productID)
	if err != nil {
		return nil, err
	}

	// Add images
	for filename, imageBytes := range images {
		imageEntity, err := s.saveImage(filename, imageBytes)
		if err != nil {
			return nil, err
		}
		imageEntity.Order = len(product.Images)
		product.Images = append(product.Images, imageEntity)
	}

	// Save product
	product, err = s.db.UpdateProduct(product)
	if err != nil {
		return nil, err
	}

	// Generate URL's
	if len(imageConfigs) > 0 {
		err := s.setImageURLsFromConfig(product.Images, imageConfigs)
		if err != nil {
			return nil, err
		}
	}

	// Add images successful
	return product, nil
}

// Update product image
func (s *Service) UpdateProductImage(productID, imageID entities.ID, order int) ([]*entities.Image, error) {
	return s.db.UpdateImage(imageID, productID, order)
}

// DeleteProductImage deletes a single image for a product
func (s *Service) DeleteProductImage(productID, imageID entities.ID) error {
	// Fetch product
	product, err := s.db.GetProduct(productID)
	if err != nil {
		return err
	}

	// Check if image owned by product
	var image *entities.Image
	for _, productImage := range product.Images {
		if productImage.ID == imageID {
			image = productImage
			break
		}
	}
	if image == nil {
		return entities.NewError(404, errors.New("product has no image with this ID"))
	}

	// Delete image
	return s.deleteImage(image)
}
