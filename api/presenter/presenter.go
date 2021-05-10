package presenter

import (
	"github.com/JenswBE/go-commerce/pkg/shortid"
	"github.com/google/uuid"
)

type Presenter struct {
	shortIDService shortid.Service
}

func New() *Presenter {
	return &Presenter{
		shortIDService: shortid.NewBase58Service(),
	}
}

func (p *Presenter) ParseID(id string) (uuid.UUID, error) {
	// Parse UUID
	pID, err := uuid.Parse(id)
	if err != nil {
		// Try to parse as short ID
		return p.shortIDService.Decode(id)
	}
	return pID, err
}
