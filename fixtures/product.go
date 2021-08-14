package fixtures

import (
	"time"

	"github.com/JenswBE/go-commerce/api/openapi"
	"github.com/JenswBE/go-commerce/entities"
	"github.com/google/uuid"
)

// #############################
// #           ENTITY          #
// #############################

func Product() *entities.Product {
	return &entities.Product{
		ID:               uuid.MustParse(ProductID),
		CreatedAt:        time.Date(2021, 8, 6, 10, 0, 0, 0, time.UTC),
		UpdatedAt:        time.Date(2021, 8, 7, 14, 0, 0, 0, time.UTC),
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

func ProductSlice() []*entities.Product {
	return []*entities.Product{
		Product(),
	}
}

// #############################
// #          OPENAPI          #
// #############################

func ProductOpenAPI() *openapi.Product {
	return &openapi.Product{
		Id:               openapi.PtrString(ProductID),
		CreatedAt:        openapi.PtrTime(time.Date(2021, 8, 6, 10, 0, 0, 0, time.UTC)),
		UpdatedAt:        openapi.PtrTime(time.Date(2021, 8, 7, 14, 0, 0, 0, time.UTC)),
		Name:             "test-name",
		DescriptionShort: openapi.PtrString("test-description-short"),
		DescriptionLong:  openapi.PtrString("test-description-long"),
		Price:            2050,
		ManufacturerId:   openapi.PtrString(ManufacturerID),
		CategoryIds:      &[]string{CategoryID},
		Status:           openapi.PRODUCTSTATUS_AVAILABLE.Ptr(),
		StockCount:       openapi.PtrInt64(5),
		ImageUrls:        &[]map[string]string{Image().URLs},
	}
}

func ProductOpenAPISlice() []openapi.Product {
	return []openapi.Product{
		*ProductOpenAPI(),
	}
}

func ProductListOpenAPI() *openapi.ProductList {
	return openapi.NewProductList(ProductOpenAPISlice())
}
