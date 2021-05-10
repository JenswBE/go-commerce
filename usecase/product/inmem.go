package product

import (
	"strings"

	"github.com/JenswBE/go-commerce/entity"
)

type Inmem struct {
	manufacturers map[entity.ID]*entity.Manufacturer
}

func NewInmem() *Inmem {
	return &Inmem{
		manufacturers: map[entity.ID]*entity.Manufacturer{},
	}
}

func (r *Inmem) GetManufacturer(id entity.ID) (*entity.Manufacturer, error) {
	// Get entity
	e, ok := r.manufacturers[id]
	if !ok {
		return nil, entity.ErrNotFound
	}

	// Return clone
	clone := *e
	return &clone, nil
}

func (r *Inmem) ListManufacturers() ([]*entity.Manufacturer, error) {
	list := make([]*entity.Manufacturer, 0, len(r.manufacturers))
	for _, manufacturer := range r.manufacturers {
		clone := *manufacturer
		list = append(list, &clone)
	}
	return list, nil
}

func (r *Inmem) SearchManufacturers(query string) ([]*entity.Manufacturer, error) {
	query = strings.ToLower(query)
	list, _ := r.ListManufacturers()
	var result []*entity.Manufacturer
	for _, manufacturer := range list {
		if strings.Contains(strings.ToLower(manufacturer.Name), query) {
			// Found, no clone needed as ListManufacturers already returns a cloned list
			result = append(result, manufacturer)
		}
	}
	return result, nil
}

func (r *Inmem) CreateManufacturer(e *entity.Manufacturer) (*entity.Manufacturer, error) {
	clone := *e
	r.manufacturers[e.ID] = &clone
	return &clone, nil
}

func (r *Inmem) UpdateManufacturer(e *entity.Manufacturer) (*entity.Manufacturer, error) {
	// Fetch manufacturer
	_, err := r.GetManufacturer(e.ID)
	if err != nil {
		return nil, err
	}

	// Replace with clone
	clone := *e
	r.manufacturers[e.ID] = &clone
	return &clone, nil
}

func (r *Inmem) DeleteManufacturer(id entity.ID) error {
	if r.manufacturers[id] == nil {
		return entity.ErrNotFound
	}
	r.manufacturers[id] = nil
	return nil
}
