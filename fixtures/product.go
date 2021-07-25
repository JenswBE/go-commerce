package fixtures

import (
	"github.com/JenswBE/go-commerce/entities"
	"github.com/google/uuid"
)

func Product() *entities.Product {
	return &entities.Product{
		ID:               uuid.MustParse(ManufacturerID),
		Name:             "test-name",
		DescriptionShort: "test-description-short",
		DescriptionLong:  "test-description-long",
		Price:            2050,
		ManufacturerID:   Manufacturer().ID,
		CategoryIDs:      []uuid.UUID{Category().ID},
		Status:           entities.ProductStatusAvailable,
		StockCount:       5,
		Images:           []*entities.Image{Image()},
	}
}
