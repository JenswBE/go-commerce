package entities

import (
	"strings"
	"time"

	"github.com/JenswBE/go-commerce/api/openapi"
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
	Price AmountInCents

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

func (p *Product) Clean() {
	p.Name = strings.TrimSpace(p.Name)
	p.DescriptionShort = strings.TrimSpace(p.DescriptionShort)
	p.DescriptionLong = strings.TrimSpace(p.DescriptionLong)
}

// Validate cleans and validates the product data
func (p *Product) Validate() error {
	// Clean entity
	p.Clean()

	// Validate simple fields
	if p.Name == "" {
		return NewError(400, openapi.GOCOMERRORCODE_PRODUCT_NAME_EMPTY, p.ID.String(), nil)
	}
	if p.Price.Int() < 0 {
		return NewError(400, openapi.GOCOMERRORCODE_PRODUCT_PRICE_NEGATIVE, p.ID.String(), nil)
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
