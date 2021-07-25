package fixtures

import (
	"github.com/JenswBE/go-commerce/entities"
	"github.com/google/uuid"
)

func Manufacturer() *entities.Manufacturer {
	return &entities.Manufacturer{
		ID:         uuid.MustParse(ManufacturerID),
		Name:       "test-name",
		WebsiteURL: "https://manufacturer.test",
		Image:      Image(),
	}
}
