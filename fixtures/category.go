package fixtures

import (
	"github.com/JenswBE/go-commerce/entities"
	"github.com/google/uuid"
)

func Category() *entities.Category {
	return &entities.Category{
		ID:          uuid.MustParse(CategoryID),
		Name:        "test-name",
		Description: "test-description",
		Order:       1,
		Image:       Image(),
	}
}
