package product

import (
	"errors"

	"github.com/JenswBE/go-commerce/entity"
	"github.com/JenswBE/go-commerce/utils/imageproxy"
)

// GetProduct fetches a single product by ID
func (s *Service) GetProduct(id entity.ID, imageConfig *imageproxy.ImageConfig) (*entity.Product, error) {
	// Fetch product
	product, err := s.db.GetProduct(id)
	if err != nil {
		return nil, err
	}

	// Generate URL's
	if imageConfig != nil {
		err := s.setImageURLsFromConfig(product.Images, *imageConfig)
		if err != nil {
			return nil, err
		}
	}

	// Get successful
	return product, nil
}

// ListProducts fetches all products
func (s *Service) ListProducts(imageConfig *imageproxy.ImageConfig) ([]*entity.Product, error) {
	// Fetch product
	products, err := s.db.ListProducts()
	if err != nil {
		return nil, err
	}

	// Generate URL's
	if imageConfig != nil {
		for _, product := range products {
			err := s.setImageURLsFromConfig(product.Images, *imageConfig)
			if err != nil {
				return nil, err
			}
		}
	}

	// Get successful
	return products, nil
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
func (s *Service) AddProductImages(productID entity.ID, images map[string][]byte, imageConfig *imageproxy.ImageConfig) (*entity.Product, error) {
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
	if imageConfig != nil {
		err := s.setImageURLsFromConfig(product.Images, *imageConfig)
		if err != nil {
			return nil, err
		}
	}

	// Add images successful
	return product, nil
}

// Update product image
func (s *Service) UpdateProductImage(productID, imageID entity.ID, order int) ([]*entity.Image, error) {
	return s.db.UpdateImage(imageID, productID, order)
}

// DeleteProductImage deletes a single image for a product
func (s *Service) DeleteProductImage(productID, imageID entity.ID) error {
	// Fetch product
	product, err := s.db.GetProduct(productID)
	if err != nil {
		return err
	}

	// Check if image owned by product
	var image *entity.Image
	for _, productImage := range product.Images {
		if productImage.ID == imageID {
			image = productImage
			break
		}
	}
	if image == nil {
		return entity.NewError(404, errors.New("product has no image with this ID"))
	}

	// Delete image
	return s.deleteImage(image)
}
