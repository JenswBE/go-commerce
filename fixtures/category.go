package fixtures

import (
	"github.com/JenswBE/go-commerce/api/openapi"
	"github.com/JenswBE/go-commerce/entities"
	"github.com/JenswBE/go-commerce/utils/generics"
)

// #############################
// #           ENTITY          #
// #############################

func Category() *entities.Category {
	return &entities.Category{
		ID:          generics.Must(entities.NewIDFromString(CategoryID)),
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
		Id:          CategoryID,
		Name:        "test-name",
		Description: openapi.PtrString("test-description"),
		Order:       1,
		ImageUrls:   &Image().URLs,
		ProductIds:  []string{},
	}
}

func CategoryOpenAPISlice() []openapi.Category {
	return []openapi.Category{
		*CategoryOpenAPI(),
	}
}

func CategoryListOpenAPI() *openapi.CategoryList {
	return openapi.NewCategoryList(CategoryOpenAPISlice())
}
