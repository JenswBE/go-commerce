package presenter

import (
	"errors"

	"github.com/JenswBE/go-commerce/api/openapi"
	"github.com/JenswBE/go-commerce/entities"
	"github.com/google/uuid"
)

func (p *Presenter) ProductFromEntity(e *entities.Product) openapi.Product {
	// Set basic fields
	m := openapi.NewProduct()
	m.SetId(p.EncodeID(e.ID))
	m.SetCreatedAt(e.CreatedAt)
	m.SetUpdatedAt(e.UpdatedAt)
	m.SetName(e.Name)
	m.SetDescriptionShort(e.DescriptionShort)
	m.SetDescriptionLong(e.DescriptionLong)
	m.SetPrice(int64(e.Price))
	m.SetCategoryIds(p.EncodeIDList(e.CategoryIDs))
	m.SetStatus(string(e.Status))
	m.SetStockCount(int64(e.StockCount))
	m.SetImageUrls(p.ImageURLListFromEntity(e.Images))

	// Set manufacturer ID
	if e.ManufacturerID != uuid.Nil {
		m.SetManufacturerId(p.EncodeID(e.ManufacturerID))
	}
	return *m
}

func (p *Presenter) ProductsListFromEntity(input []*entities.Product) []openapi.Product {
	output := make([]openapi.Product, 0, len(input))
	for _, product := range input {
		output = append(output, p.ProductFromEntity(product))
	}
	return output
}

func (p *Presenter) ProductToEntity(id uuid.UUID, product openapi.Product) (*entities.Product, error) {
	// Build entity
	e := &entities.Product{
		ID:               id,
		Name:             product.GetName(),
		DescriptionShort: product.GetDescriptionShort(),
		DescriptionLong:  product.GetDescriptionLong(),
		Price:            int(product.GetPrice()),
		CategoryIDs:      nil,
		ManufacturerID:   uuid.Nil,
		Status:           entities.ProductStatus(product.GetStatus()),
		StockCount:       int(product.GetStockCount()),
	}

	// Parse category ID's
	if len(product.GetCategoryIds()) > 0 {
		catIDs, err := p.ParseIDList(product.GetCategoryIds())
		if err != nil {
			return nil, entities.NewError(400, errors.New("category_ids is invalid"))
		}
		e.CategoryIDs = catIDs
	}

	// Parse manufacturer ID
	if product.GetManufacturerId() != "" {
		manID, err := p.ParseID(product.GetManufacturerId())
		if err != nil {
			return nil, entities.NewError(400, errors.New("manufacturer_id is invalid"))
		}
		e.ManufacturerID = manID
	}

	// Successful
	return e, nil
}
