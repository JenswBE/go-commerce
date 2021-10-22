package presenter

import (
	"github.com/JenswBE/go-commerce/api/openapi"
	"github.com/JenswBE/go-commerce/entities"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"github.com/rs/zerolog/log"
)

func (p *Presenter) ProductFromEntity(e *entities.Product) openapi.Product {
	// Set basic fields
	output := openapi.NewProduct(e.Name, int64(e.Price))
	output.SetId(p.EncodeID(e.ID))
	output.SetCreatedAt(e.CreatedAt)
	output.SetUpdatedAt(e.UpdatedAt)
	output.SetDescriptionShort(e.DescriptionShort)
	output.SetDescriptionLong(e.DescriptionLong)
	output.SetCategoryIds(p.EncodeIDList(e.CategoryIDs))
	output.SetStockCount(int64(e.StockCount))
	output.SetImageUrls(p.ImageURLsSliceFromEntity(e.Images))

	// Set status
	status, err := openapi.NewProductStatusFromValue(e.Status.String())
	if err != nil {
		defaultStatus := openapi.PRODUCTSTATUS_ARCHIVED
		log.Warn().Err(err).Stringer("status", e.Status).Msgf("Unknown product status received from entity, defaulting to %s", defaultStatus)
		status = defaultStatus.Ptr()
	}
	output.SetStatus(*status)

	// Set manufacturer ID
	if e.ManufacturerID != uuid.Nil {
		output.SetManufacturerId(p.EncodeID(e.ManufacturerID))
	}
	return *output
}

func (p *Presenter) ProductSliceFromEntity(input []*entities.Product) []openapi.Product {
	output := make([]openapi.Product, 0, len(input))
	for _, product := range input {
		output = append(output, p.ProductFromEntity(product))
	}
	return output
}

func (p *Presenter) ProductListFromEntity(input []*entities.Product) openapi.ProductList {
	return *openapi.NewProductList(p.ProductSliceFromEntity(input))
}

func (p *Presenter) ResolvedProductFromEntity(e *entities.ResolvedProduct) (openapi.ResolvedProduct, error) {
	// Convert to basic product
	product := p.ProductFromEntity(&e.Product)
	output := openapi.ResolvedProduct{}
	err := copier.Copy(&output, &product)
	if err != nil {
		return openapi.ResolvedProduct{}, entities.NewError(500, openapi.ERRORCODE_UNKNOWN_ERROR, e.ID.String(), err)
	}

	// Set manufacturer
	if e.Manufacturer != nil {
		manufacturer := p.ManufacturerFromEntity(e.Manufacturer)
		output.Manufacturer = &manufacturer
	}

	// Set categories
	categories := make([]openapi.Category, 0, len(e.Categories))
	for _, category := range e.Categories {
		categories = append(categories, p.CategoryFromEntity(category))
	}
	output.Categories = &categories
	return output, nil
}

func (p *Presenter) ProductToEntity(id uuid.UUID, product openapi.Product) (*entities.Product, error) {
	// Build entity
	e := &entities.Product{
		ID:               id,
		Name:             product.GetName(),
		DescriptionShort: product.GetDescriptionShort(),
		DescriptionLong:  product.GetDescriptionLong(),
		Price:            int(product.GetPrice()),
		CategoryIDs:      nil,
		ManufacturerID:   uuid.Nil,
		Status:           entities.ProductStatus(product.GetStatus()),
		StockCount:       int(product.GetStockCount()),
	}

	// Parse category ID's
	if len(product.GetCategoryIds()) > 0 {
		catIDs, err := p.ParseIDList(product.GetCategoryIds())
		if err != nil {
			return nil, entities.NewError(400, openapi.ERRORCODE_PRODUCT_CATEGORY_IDS_INVALID, id.String(), err)
		}
		e.CategoryIDs = catIDs
	}

	// Parse manufacturer ID
	if product.GetManufacturerId() != "" {
		manID, err := p.ParseID(product.GetManufacturerId())
		if err != nil {
			return nil, entities.NewError(400, openapi.ERRORCODE_PRODUCT_MANUFACTURER_ID_INVALID, id.String(), err)
		}
		e.ManufacturerID = manID
	}

	// Successful
	return e, nil
}
