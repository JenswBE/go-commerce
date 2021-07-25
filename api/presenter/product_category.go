package presenter

import (
	"errors"

	"github.com/JenswBE/go-commerce/api/openapi"
	"github.com/JenswBE/go-commerce/entities"
	"github.com/google/uuid"
)

func (p *Presenter) CategoryFromEntity(e *entities.Category) openapi.Category {
	// Set basic fields
	c := openapi.NewCategory(e.Name, int64(e.Order))
	c.SetId(p.EncodeID(e.ID))
	c.SetDescription(e.Description)
	c.SetProductIds(p.EncodeIDList(e.ProductIDs))

	// Set parent ID
	if e.ParentID != uuid.Nil {
		c.SetParentId(p.EncodeID(e.ParentID))
	}

	// Set image URL
	if e.Image != nil && e.Image.URL != "" {
		c.SetImageUrl(e.Image.URL)
	}
	return *c
}

func (p *Presenter) CategoriesListFromEntity(input []*entities.Category) []openapi.Category {
	output := make([]openapi.Category, 0, len(input))
	for _, category := range input {
		output = append(output, p.CategoryFromEntity(category))
	}
	return output
}

func (p *Presenter) CategoryToEntity(id uuid.UUID, category openapi.Category) (*entities.Category, error) {
	// Build entity
	e := &entities.Category{
		ID:          id,
		Name:        category.GetName(),
		Description: category.GetDescription(),
		ParentID:    uuid.Nil,
		Order:       int(category.Order),
		ProductIDs:  nil, // read-only
	}

	// Parse parent ID
	if category.GetParentId() != "" {
		pID, err := p.ParseID(category.GetParentId())
		if err != nil {
			return nil, entities.NewError(400, errors.New("parent_id is invalid"))
		}
		e.ParentID = pID
	}

	// Successful
	return e, nil
}
