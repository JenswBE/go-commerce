package presenter

import (
	"github.com/JenswBE/go-commerce/entity"
	"github.com/JenswBE/go-commerce/pkg/shortid"
	"github.com/google/uuid"
)

type Presenter struct {
	shortIDService shortid.Service
}

func New(shortIDService shortid.Service) *Presenter {
	return &Presenter{shortIDService}
}

func (p *Presenter) ParseID(id string) (uuid.UUID, error) {
	// Parse UUID
	pID, err := uuid.Parse(id)
	if err != nil {
		// Try to parse as short ID
		var shortErr error
		pID, shortErr = p.shortIDService.Decode(id)
		if shortErr != nil {
			// When also short ID parsing fails, we return the initial error
			return uuid.Nil, entity.NewError(400, err)
		}
	}
	return pID, nil
}
