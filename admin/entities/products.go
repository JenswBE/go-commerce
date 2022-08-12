package entities

import (
	"fmt"

	"github.com/JenswBE/go-commerce/entities"
)

type ProductsListTemplate struct {
	BaseData
	Products         []*entities.Product
	ManufacturersMap map[entities.ID]entities.Manufacturer
	PublicURLMap     map[entities.ID]string
}

func (t ProductsListTemplate) GetTemplateName() string {
	return "productsList"
}

type ProductsFormTemplate struct {
	BaseData
	IsNew         bool
	Product       Product
	Categories    []*entities.Category
	Manufacturers []*entities.Manufacturer
}

func (t ProductsFormTemplate) GetTemplateName() string {
	return "productsForm"
}

type Product struct {
	Name             string   `form:"name"`
	Price            string   `form:"price"`
	ManufacturerID   string   `form:"manufacturer_id"`
	CategoryIDs      []string `form:"category_ids"`
	StockCount       int      `form:"stock_count"`
	DescriptionShort string   `form:"description_short"`
	DescriptionLong  string   `form:"description_long"`
}

func ProductFromEntity(e *entities.ResolvedProduct) Product {
	categoryIDs := make([]string, 0, len(e.CategoryIDs))
	for _, catID := range e.CategoryIDs {
		categoryIDs = append(categoryIDs, catID.String())
	}

	return Product{
		Name:             e.Name,
		Price:            e.Price.String(),
		ManufacturerID:   e.ManufacturerID.String(),
		CategoryIDs:      categoryIDs,
		StockCount:       e.StockCount,
		DescriptionShort: e.DescriptionShort,
		DescriptionLong:  e.DescriptionLong,
	}
}

func (e Product) ToEntity() (entities.Product, error) {
	// Parse price
	price, err := entities.NewAmountInCentsFromString(e.Price)
	if err != nil {
		return entities.Product{}, fmt.Errorf("failed to convert price string to AmountInCents: %w", err)
	}

	// Parse manufacturer ID
	manufacturerID := entities.NewNilID()
	if e.ManufacturerID != "" {
		manufacturerID, err = entities.NewIDFromString(e.ManufacturerID)
		if err != nil {
			return entities.Product{}, fmt.Errorf("failed to parse manufacturer ID to a UUID: %w", err)
		}
	}

	// Parse category ID's
	categoryIDs := make([]entities.ID, 0, len(e.CategoryIDs))
	for _, id := range e.CategoryIDs {
		catID, err := entities.NewIDFromString(id)
		if err != nil {
			return entities.Product{}, fmt.Errorf("failed to parse category ID to a UUID: %w", err)
		}
		categoryIDs = append(categoryIDs, catID)
	}

	return entities.Product{
		Name:             e.Name,
		Price:            price,
		ManufacturerID:   manufacturerID,
		CategoryIDs:      categoryIDs,
		StockCount:       e.StockCount,
		DescriptionShort: e.DescriptionShort,
		DescriptionLong:  e.DescriptionLong,
	}, nil
}

func (e Product) HasCategoryID(categoryID string) bool {
	for _, catID := range e.CategoryIDs {
		if catID == categoryID {
			return true
		}
	}
	return false
}
