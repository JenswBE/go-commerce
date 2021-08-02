package fixtures

import (
	"github.com/JenswBE/go-commerce/api/openapi"
	"github.com/JenswBE/go-commerce/entities"
	"github.com/google/uuid"
)

// #############################
// #           ENTITY          #
// #############################

func Manufacturer() *entities.Manufacturer {
	return &entities.Manufacturer{
		ID:         uuid.MustParse(ManufacturerID),
		Name:       "test-name",
		WebsiteURL: "https://manufacturer.test",
		Image:      Image(),
	}
}

func ManufacturerSlice() []*entities.Manufacturer {
	return []*entities.Manufacturer{
		Manufacturer(),
	}
}

// #############################
// #          OPENAPI          #
// #############################

func ManufacturerOpenAPI() *openapi.Manufacturer {
	return &openapi.Manufacturer{
		Id:         openapi.PtrString(ManufacturerID),
		Name:       "test-name",
		WebsiteUrl: openapi.PtrString("https://manufacturer.test"),
		ImageUrl:   openapi.PtrString("http://image.test"),
	}
}

func ManufacturerOpenAPISlice() []*openapi.Manufacturer {
	return []*openapi.Manufacturer{
		ManufacturerOpenAPI(),
	}
}
