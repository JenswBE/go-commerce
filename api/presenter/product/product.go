package product

import (
	"github.com/JenswBE/go-commerce/api/openapi"
	"github.com/JenswBE/go-commerce/api/presenter"
	"github.com/JenswBE/go-commerce/entities"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"github.com/rs/zerolog/log"
)

const defaultStatus = openapi.PRODUCTSTATUS_ARCHIVED

func ProductFromEntity(p *presenter.Presenter, input *entities.Product) openapi.Product {
	// Set basic fields
	output := openapi.NewProduct(p.String(input.Name), int64(input.Price))
	output.SetId(p.EncodeID(input.ID))
	output.SetCreatedAt(input.CreatedAt)
	output.SetUpdatedAt(input.UpdatedAt)
	output.SetDescriptionShort(p.String(input.DescriptionShort))
	output.SetDescriptionLong(p.String(input.DescriptionLong))
	output.SetCategoryIds(p.EncodeIDList(input.CategoryIDs))
	output.SetStockCount(int64(input.StockCount))
	output.SetImageUrls(ImageURLsSliceFromEntity(p, input.Images))

	// Set status
	status, err := openapi.NewProductStatusFromValue(input.Status.String())
	if err != nil {
		log.Warn().Err(err).Stringer("status", input.Status).Msgf("Unknown product status received from entity, defaulting to %s", defaultStatus)
		status = defaultStatus.Ptr()
	}
	output.SetStatus(*status)

	// Set manufacturer ID
	if input.ManufacturerID != uuid.Nil {
		output.SetManufacturerId(p.EncodeID(input.ManufacturerID))
	}
	return *output
}

func ProductSliceFromEntity(p *presenter.Presenter, input []*entities.Product) []openapi.Product {
	output := make([]openapi.Product, 0, len(input))
	for _, product := range input {
		output = append(output, ProductFromEntity(p, product))
	}
	return output
}

func ProductListFromEntity(p *presenter.Presenter, input []*entities.Product) openapi.ProductList {
	return *openapi.NewProductList(ProductSliceFromEntity(p, input))
}

func ResolvedProductFromEntity(p *presenter.Presenter, input *entities.ResolvedProduct) (openapi.ResolvedProduct, error) {
	// Convert to basic product
	product := ProductFromEntity(p, &input.Product)
	output := openapi.ResolvedProduct{}
	err := copier.Copy(&output, &product)
	if err != nil {
		return openapi.ResolvedProduct{}, entities.NewError(500, openapi.GOCOMERRORCODE_UNKNOWN_ERROR, input.ID.String(), err)
	}

	// Set manufacturer
	if input.Manufacturer != nil {
		manufacturer := ManufacturerFromEntity(p, input.Manufacturer)
		output.Manufacturer = &manufacturer
	}

	// Set categories
	categories := make([]openapi.Category, 0, len(input.Categories))
	for _, category := range input.Categories {
		categories = append(categories, CategoryFromEntity(p, category))
	}
	output.Categories = categories
	return output, nil
}
