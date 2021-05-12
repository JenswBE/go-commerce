package presenter

import (
	"errors"

	"github.com/JenswBE/go-commerce/entity"
	"github.com/google/uuid"
)

type Category struct {
	ID string `json:"id"`
	CategoryData
}

type CategoryData struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	ParentID    string `json:"parent_id"`
}

func (p *Presenter) CategoryFromEntity(e *entity.Category) Category {
	c := Category{
		ID: p.shortIDService.Encode(e.ID),
		CategoryData: CategoryData{
			Name:        e.Name,
			Description: e.Description,
			ParentID:    "",
		},
	}
	if e.ParentID != uuid.Nil {
		c.CategoryData.ParentID = p.shortIDService.Encode(e.ParentID)
	}
	return c
}

func (p *Presenter) CategoriesListFromEntity(input []*entity.Category) []Category {
	output := make([]Category, 0, len(input))
	for _, category := range input {
		output = append(output, p.CategoryFromEntity(category))
	}
	return output
}

func (p *Presenter) CategoryToEntity(id uuid.UUID, data CategoryData) (*entity.Category, error) {
	// Build entity
	e := &entity.Category{
		ID:          id,
		Name:        data.Name,
		Description: data.Description,
		ParentID:    uuid.Nil,
	}

	// Parse parent ID
	if data.ParentID != "" {
		pID, err := p.shortIDService.Decode(data.ParentID)
		if err != nil {
			return nil, entity.NewError(400, errors.New("parent_id is invalid"))
		}
		e.ParentID = pID
	}

	// Successful
	return e, nil
}
