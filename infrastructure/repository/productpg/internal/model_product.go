package internal

import (
	"github.com/JenswBE/go-commerce/entity"
	"github.com/google/uuid"
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
}

func ProductPgToEntity(c *Product) *entity.Product {
	product := &entity.Product{
		ID:               uuid.MustParse(c.ID),
		CreatedAt:        c.CreatedAt,
		UpdatedAt:        c.UpdatedAt,
		Name:             c.Name,
		DescriptionShort: c.DescriptionShort,
		DescriptionLong:  c.DescriptionLong,
		Price:            c.Price,
		CategoryIDs:      nil,
		ManufacturerID:   uuid.Nil,
		Status:           entity.ProductStatus(c.Status),
		StockCount:       c.StockCount,
	}
	if c.ManufacturerID != nil {
		product.ManufacturerID = uuid.MustParse(*c.ManufacturerID)
	}
	if len(c.Categories) > 0 {
		product.CategoryIDs = make([]uuid.UUID, len(c.Categories))
		for i, cat := range c.Categories {
			product.CategoryIDs[i] = uuid.MustParse(cat.ID)
		}
	}
	return product
}

func ProductsListPgToEntity(p []*Product) []*entity.Product {
	output := make([]*entity.Product, 0, len(p))
	for _, product := range p {
		output = append(output, ProductPgToEntity(product))
	}
	return output
}

func ProductEntityToPg(e *entity.Product) *Product {
	product := &Product{
		Base:             Base{ID: e.ID.String()},
		Name:             e.Name,
		DescriptionShort: e.DescriptionShort,
		DescriptionLong:  e.DescriptionLong,
		Price:            e.Price,
		Categories:       nil,
		ManufacturerID:   nil,
	}
	if e.ManufacturerID != uuid.Nil {
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
