package product

import (
	"github.com/JenswBE/go-commerce/api/openapi"
	"github.com/JenswBE/go-commerce/api/presenter"
	"github.com/JenswBE/go-commerce/entities"
	"github.com/google/uuid"
)

func CategoryFromEntity(p *presenter.Presenter, input *entities.Category) openapi.Category {
	// Set basic fields
	output := openapi.NewCategory(input.Name, int64(input.Order))
	output.SetId(p.EncodeID(input.ID))
	output.SetDescription(input.Description)
	output.SetProductIds(p.EncodeIDList(input.ProductIDs))

	// Set parent ID
	if input.ParentID != uuid.Nil {
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

func CategoryToEntity(p *presenter.Presenter, id uuid.UUID, input openapi.Category) (*entities.Category, error) {
	// Build entity
	output := &entities.Category{
		ID:          id,
		Name:        input.GetName(),
		Description: input.GetDescription(),
		ParentID:    uuid.Nil,
		Order:       int(input.Order),
		ProductIDs:  nil, // read-only
	}

	// Parse parent ID
	if input.GetParentId() != "" {
		pID, err := p.ParseID(input.GetParentId())
		if err != nil {
			return nil, entities.NewError(400, openapi.GOCOMERRORCODE_CATEGORY_PARENT_ID_INVALID, id.String(), err)
		}
		output.ParentID = pID
	}

	// Successful
	return output, nil
}
