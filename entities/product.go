package entities

import (
	"errors"
	"time"
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
	Images         []*Image
}

type ProductStatus string

const ProductStatusAvailable = "AVAILABLE"
const ProductStatusArchived = "ARCHIVED"

func (status ProductStatus) String() string {
	return string(status)
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

// ResolvedProduct is a product for which related entities
// are included. This way all information is immediately at hand.
type ResolvedProduct struct {
	Product
	Manufacturer *Manufacturer
	Categories   []*Category
}
