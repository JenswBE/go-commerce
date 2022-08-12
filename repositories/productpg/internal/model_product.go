package internal

import (
	"github.com/JenswBE/go-commerce/entities"
	"github.com/JenswBE/go-commerce/utils/generics"
)

type Product struct {
	Base
	Name             string
	DescriptionShort string
	DescriptionLong  string
	Price            int
	Categories       []Category `gorm:"many2many:product_categories;"`
	ManufacturerID   *string    `gorm:"type:uuid"`
	Status           string
	StockCount       int
	Images           []Image `gorm:"polymorphic:Owner;"`
}

func ProductPgToEntity(c *Product) *entities.Product {
	product := &entities.Product{
		ID:               generics.Must(entities.NewIDFromString(c.ID)),
		CreatedAt:        c.CreatedAt,
		UpdatedAt:        c.UpdatedAt,
		Name:             c.Name,
		DescriptionShort: c.DescriptionShort,
		DescriptionLong:  c.DescriptionLong,
		Price:            entities.NewAmountInCents(c.Price),
		CategoryIDs:      nil,
		ManufacturerID:   entities.NewNilID(),
		Status:           entities.ProductStatus(c.Status),
		StockCount:       c.StockCount,
		Images:           ImagesListPgToEntity(c.Images),
	}
	if c.ManufacturerID != nil {
		product.ManufacturerID = generics.Must(entities.NewIDFromString(*c.ManufacturerID))
	}
	if len(c.Categories) > 0 {
		product.CategoryIDs = make([]entities.ID, len(c.Categories))
		for i, cat := range c.Categories {
			product.CategoryIDs[i] = generics.Must(entities.NewIDFromString(cat.ID))
		}
	}
	if len(c.Categories) > 0 {
		product.CategoryIDs = make([]entities.ID, len(c.Categories))
		for i, cat := range c.Categories {
			product.CategoryIDs[i] = generics.Must(entities.NewIDFromString(cat.ID))
		}
	}
	return product
}

func ProductsListPgToEntity(p []*Product) []*entities.Product {
	output := make([]*entities.Product, 0, len(p))
	for _, product := range p {
		output = append(output, ProductPgToEntity(product))
	}
	return output
}

func ProductEntityToPg(e *entities.Product) *Product {
	product := &Product{
		Base:             Base{ID: e.ID.String()},
		Name:             e.Name,
		DescriptionShort: e.DescriptionShort,
		DescriptionLong:  e.DescriptionLong,
		Price:            e.Price.Int(),
		Categories:       nil,
		ManufacturerID:   nil,
		Images:           ImagesListEntityToPg(e.Images),
		StockCount:       e.StockCount,
	}
	if !e.ManufacturerID.IsNil() {
		id := e.ManufacturerID.String()
		product.ManufacturerID = &id
	}
	if len(e.CategoryIDs) > 0 {
		product.Categories = make([]Category, len(e.CategoryIDs))
		for i, catID := range e.CategoryIDs {
			product.Categories[i] = Category{Base: Base{ID: catID.String()}}
		}
	}
	return product
}
