package fixtures

import (
	"github.com/JenswBE/go-commerce/api/openapi"
	"github.com/JenswBE/go-commerce/entities"
	"github.com/google/uuid"
)

// #############################
// #           ENTITY          #
// #############################

func Category() *entities.Category {
	return &entities.Category{
		ID:          uuid.MustParse(CategoryID),
		Name:        "test-name",
		Description: "test-description",
		Order:       1,
		Image:       Image(),
	}
}

func CategorySlice() []*entities.Category {
	return []*entities.Category{
		Category(),
	}
}

// #############################
// #          OPENAPI          #
// #############################

func CategoryOpenAPI() *openapi.Category {
	return &openapi.Category{
		Id:          openapi.PtrString(CategoryID),
		Name:        "test-name",
		Description: openapi.PtrString("test-description"),
		Order:       1,
		ImageUrl:    &Image().URL,
		ProductIds:  &[]string{},
	}
}

func CategoryOpenAPISlice() []*openapi.Category {
	return []*openapi.Category{
		CategoryOpenAPI(),
	}
}
