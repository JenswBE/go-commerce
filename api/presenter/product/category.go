package product

import (
	"github.com/JenswBE/go-commerce/api/openapi"
	"github.com/JenswBE/go-commerce/api/presenter"
	"github.com/JenswBE/go-commerce/entities"
)

func CategoryFromEntity(p *presenter.Presenter, input *entities.Category) openapi.Category {
	// Set basic fields
	output := openapi.NewCategory(
		p.EncodeID(input.ID),
		p.String(input.Name),
		int64(input.Order),
	)
	output.SetDescription(p.String(input.Description))
	output.SetProductIds(p.EncodeIDList(input.ProductIDs))

	// Set parent ID
	if !input.ParentID.IsNil() {
		output.SetParentId(p.EncodeID(input.ParentID))
	}

	// Set image URL
	if input.Image != nil && len(input.Image.URLs) > 0 {
		output.SetImageUrls(input.Image.URLs)
	}
	return *output
}

func CategorySliceFromEntity(p *presenter.Presenter, input []*entities.Category) []openapi.Category {
	output := make([]openapi.Category, 0, len(input))
	for _, category := range input {
		output = append(output, CategoryFromEntity(p, category))
	}
	return output
}

func CategoryListFromEntity(p *presenter.Presenter, input []*entities.Category) openapi.CategoryList {
	return *openapi.NewCategoryList(CategorySliceFromEntity(p, input))
}
