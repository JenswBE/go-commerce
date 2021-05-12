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
	pID, err := p.shortIDService.Decode(id)
	if err != nil {
		return uuid.Nil, entity.NewError(400, err)
	}
	return pID, nil
}
