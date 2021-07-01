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

func (p *Presenter) ParseIDList(ids []string) ([]uuid.UUID, error) {
	output := make([]uuid.UUID, 0, len(ids))
	for _, id := range ids {
		pID, err := p.ParseID(id)
		if err != nil {
			return nil, err
		}
		output = append(output, pID)
	}
	return output, nil
}

func (p *Presenter) EncodeID(id uuid.UUID) string {
	return p.shortIDService.Encode(id)
}

func (p *Presenter) EncodeIDList(ids []uuid.UUID) []string {
	output := make([]string, 0, len(ids))
	for _, id := range ids {
		output = append(output, p.shortIDService.Encode(id))
	}
	return output
}
