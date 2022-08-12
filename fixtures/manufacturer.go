package fixtures

import (
	"github.com/JenswBE/go-commerce/api/openapi"
	"github.com/JenswBE/go-commerce/entities"
	"github.com/JenswBE/go-commerce/utils/generics"
)

// #############################
// #           ENTITY          #
// #############################

func Manufacturer() *entities.Manufacturer {
	return &entities.Manufacturer{
		ID:         generics.Must(entities.NewIDFromString(ManufacturerID)),
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
		ImageUrls:  &Image().URLs,
	}
}

func ManufacturerOpenAPISlice() []openapi.Manufacturer {
	return []openapi.Manufacturer{
		*ManufacturerOpenAPI(),
	}
}

func ManufacturerListOpenAPI() *openapi.ManufacturerList {
	return openapi.NewManufacturerList(ManufacturerOpenAPISlice())
}
