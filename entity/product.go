package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

// Product data
type Product struct {
	ID               ID
	CreatedAt        time.Time
	UpdatedAt        time.Time
	Name             string
	DescriptionShort string
	DescriptionLong  string

	// Price of a single product in cents (1/100)
	Price int

	CategoryIDs    []ID
	ManufacturerID ID
	Status         ProductStatus
	StockCount     int
}

type ProductStatus string

const ProductStatusAvailable = "AVAILABLE"

// NewProduct creates a new product
func NewProduct(name string, price int) (*Product, error) {
	b := &Product{
		ID:               NewID(),
		CreatedAt:        time.Now().UTC(),
		UpdatedAt:        time.Now().UTC(),
		Name:             name,
		DescriptionShort: "",
		DescriptionLong:  "",
		Price:            price,
		CategoryIDs:      nil,
		ManufacturerID:   uuid.Nil,
		Status:           ProductStatusAvailable,
		StockCount:       0,
	}
	err := b.Validate()
	if err != nil {
		return nil, NewError(400, ErrInvalidEntity)
	}
	return b, nil
}

// Validate validates the product data
func (c *Product) Validate() error {
	// Validate simple fields
	if c.Name == "" {
		return NewError(400, errors.New("product name is mandatory"))
	}
	if c.Price < 0 {
		return NewError(400, errors.New("product price cannot be negative"))
	}

	// Entity is valid
	return nil
}
