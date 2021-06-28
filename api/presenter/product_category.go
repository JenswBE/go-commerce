package presenter

import (
	"errors"

	"github.com/JenswBE/go-commerce/api/openapi"
	"github.com/JenswBE/go-commerce/entity"
	"github.com/google/uuid"
)

func (p *Presenter) CategoryFromEntity(e *entity.Category) openapi.Category {
	c := openapi.NewCategory()
	c.SetId(p.shortIDService.Encode(e.ID))
	c.SetName(e.Name)
	c.SetDescription(e.Description)
	c.SetParentId("")
	if e.ParentID != uuid.Nil {
		c.SetParentId(p.shortIDService.Encode(e.ParentID))
	}
	return *c
}

func (p *Presenter) CategoriesListFromEntity(input []*entity.Category) []openapi.Category {
	output := make([]openapi.Category, 0, len(input))
	for _, category := range input {
		output = append(output, p.CategoryFromEntity(category))
	}
	return output
}

func (p *Presenter) CategoryToEntity(id uuid.UUID, category openapi.Category) (*entity.Category, error) {
	// Build entity
	e := &entity.Category{
		ID:          id,
		Name:        category.GetName(),
		Description: category.GetDescription(),
		ParentID:    uuid.Nil,
	}

	// Parse parent ID
	if category.GetParentId() != "" {
		pID, err := p.shortIDService.Decode(category.GetParentId())
		if err != nil {
			return nil, entity.NewError(400, errors.New("parent_id is invalid"))
		}
		e.ParentID = pID
	}

	// Successful
	return e, nil
}
