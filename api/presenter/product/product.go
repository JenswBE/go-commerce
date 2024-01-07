package product

import (
	"github.com/jinzhu/copier"
	"github.com/rs/zerolog/log"

	"github.com/JenswBE/go-commerce/api/openapi"
	"github.com/JenswBE/go-commerce/api/presenter"
	"github.com/JenswBE/go-commerce/entities"
)

const defaultStatus = openapi.PRODUCTSTATUS_ARCHIVED

func ProductFromEntity(p *presenter.Presenter, input *entities.Product) openapi.Product {
	// Set basic fields
	output := openapi.NewProduct(
		p.EncodeID(input.ID),
		input.CreatedAt,
		input.UpdatedAt,
		p.String(input.Name),
		int64(input.Price.Int()),
	)
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
	if !input.ManufacturerID.IsNil() {
		output.SetManufacturerId(p.EncodeID(input.ManufacturerID))
	}
	return *output
}

func ProductListFromEntity(p *presenter.Presenter, input []*entities.Product) openapi.ProductList {
	return *openapi.NewProductList(presenter.SliceFromEntity(p, input, ProductFromEntity))
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
	output.Categories = presenter.SliceFromEntity(p, input.Categories, CategoryFromEntity)
	return output, nil
}
