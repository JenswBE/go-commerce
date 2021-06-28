package product

import (
	"github.com/JenswBE/go-commerce/entity"
)

type Inmem struct {
	categories    map[entity.ID]*entity.Category
	manufacturers map[entity.ID]*entity.Manufacturer
	products      map[entity.ID]*entity.Product
}

func NewInmem() *Inmem {
	return &Inmem{
		categories:    map[entity.ID]*entity.Category{},
		manufacturers: map[entity.ID]*entity.Manufacturer{},
		products:      map[entity.ID]*entity.Product{},
	}
}

func (r *Inmem) GetCategory(id entity.ID) (*entity.Category, error) {
	// Get entity
	e, ok := r.categories[id]
	if !ok {
		return nil, entity.ErrNotFound
	}

	// Return clone
	clone := *e
	return &clone, nil
}

func (r *Inmem) ListCategories() ([]*entity.Category, error) {
	list := make([]*entity.Category, 0, len(r.categories))
	for _, category := range r.categories {
		clone := *category
		list = append(list, &clone)
	}
	return list, nil
}

func (r *Inmem) CreateCategory(e *entity.Category) (*entity.Category, error) {
	clone := *e
	r.categories[e.ID] = &clone
	return &clone, nil
}

func (r *Inmem) UpdateCategory(e *entity.Category) (*entity.Category, error) {
	// Fetch category
	_, err := r.GetCategory(e.ID)
	if err != nil {
		return nil, err
	}

	// Replace with clone
	clone := *e
	r.categories[e.ID] = &clone
	return &clone, nil
}

func (r *Inmem) DeleteCategory(id entity.ID) error {
	if r.categories[id] == nil {
		return entity.ErrNotFound
	}
	r.categories[id] = nil
	return nil
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

func (r *Inmem) GetProduct(id entity.ID) (*entity.Product, error) {
	// Get entity
	e, ok := r.products[id]
	if !ok {
		return nil, entity.ErrNotFound
	}

	// Return clone
	clone := *e
	return &clone, nil
}

func (r *Inmem) ListProducts() ([]*entity.Product, error) {
	list := make([]*entity.Product, 0, len(r.products))
	for _, product := range r.products {
		clone := *product
		list = append(list, &clone)
	}
	return list, nil
}

func (r *Inmem) CreateProduct(e *entity.Product) (*entity.Product, error) {
	clone := *e
	r.products[e.ID] = &clone
	return &clone, nil
}

func (r *Inmem) UpdateProduct(e *entity.Product) (*entity.Product, error) {
	// Fetch product
	_, err := r.GetProduct(e.ID)
	if err != nil {
		return nil, err
	}

	// Replace with clone
	clone := *e
	r.products[e.ID] = &clone
	return &clone, nil
}

func (r *Inmem) DeleteProduct(id entity.ID) error {
	if r.products[id] == nil {
		return entity.ErrNotFound
	}
	r.products[id] = nil
	return nil
}
