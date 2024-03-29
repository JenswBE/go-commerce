package fixtures

import (
	"time"

	"github.com/JenswBE/go-commerce/api/openapi"
	"github.com/JenswBE/go-commerce/entities"
	"github.com/JenswBE/go-commerce/utils/generics"
)

// #############################
// #           ENTITY          #
// #############################

func Product() *entities.Product {
	return &entities.Product{
		ID:               generics.Must(entities.NewIDFromString(ProductID)),
		CreatedAt:        time.Date(2021, 8, 6, 10, 0, 0, 0, time.UTC),
		UpdatedAt:        time.Date(2021, 8, 7, 14, 0, 0, 0, time.UTC),
		Name:             "test-name",
		DescriptionShort: "test-description-short",
		DescriptionLong:  "test-description-long",
		Price:            entities.NewAmountInCents(2050),
		ManufacturerID:   Manufacturer().ID,
		CategoryIDs:      []entities.ID{Category().ID},
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

func ResolvedProduct() *entities.ResolvedProduct {
	return &entities.ResolvedProduct{
		Product:      *Product(),
		Manufacturer: Manufacturer(),
		Categories:   []*entities.Category{Category()},
	}
}

// #############################
// #          OPENAPI          #
// #############################

func ProductOpenAPI() *openapi.Product {
	return &openapi.Product{
		Id:               ProductID,
		CreatedAt:        time.Date(2021, 8, 6, 10, 0, 0, 0, time.UTC),
		UpdatedAt:        time.Date(2021, 8, 7, 14, 0, 0, 0, time.UTC),
		Name:             "test-name",
		DescriptionShort: openapi.PtrString("test-description-short"),
		DescriptionLong:  openapi.PtrString("test-description-long"),
		Price:            2050,
		ManufacturerId:   openapi.PtrString(ManufacturerID),
		CategoryIds:      []string{CategoryID},
		Status:           openapi.PRODUCTSTATUS_AVAILABLE.Ptr(),
		StockCount:       openapi.PtrInt64(5),
		ImageUrls:        []map[string]string{Image().URLs},
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

func ResolvedProductOpenAPI() *openapi.ResolvedProduct {
	return &openapi.ResolvedProduct{
		Id:               ProductID,
		CreatedAt:        time.Date(2021, 8, 6, 10, 0, 0, 0, time.UTC),
		UpdatedAt:        time.Date(2021, 8, 7, 14, 0, 0, 0, time.UTC),
		Name:             "test-name",
		DescriptionShort: openapi.PtrString("test-description-short"),
		DescriptionLong:  openapi.PtrString("test-description-long"),
		Price:            2050,
		ManufacturerId:   openapi.PtrString(ManufacturerID),
		CategoryIds:      []string{CategoryID},
		Status:           openapi.PRODUCTSTATUS_AVAILABLE.Ptr(),
		StockCount:       openapi.PtrInt64(5),
		ImageUrls:        []map[string]string{Image().URLs},
		Manufacturer:     ManufacturerOpenAPI(),
		Categories:       []openapi.Category{*CategoryOpenAPI()},
	}
}
