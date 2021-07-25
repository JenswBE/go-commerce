package product

import "github.com/JenswBE/go-commerce/entities"

type Inmem struct {
	categories    map[entities.ID]*entities.Category
	manufacturers map[entities.ID]*entities.Manufacturer
	products      map[entities.ID]*entities.Product
}

func NewInmem() *Inmem {
	return &Inmem{
		categories:    map[entities.ID]*entities.Category{},
		manufacturers: map[entities.ID]*entities.Manufacturer{},
		products:      map[entities.ID]*entities.Product{},
	}
}

func (r *Inmem) GetCategory(id entities.ID) (*entities.Category, error) {
	// Get entity
	e, ok := r.categories[id]
	if !ok {
		return nil, entities.ErrNotFound
	}

	// Return clone
	clone := *e
	return &clone, nil
}

func (r *Inmem) ListCategories() ([]*entities.Category, error) {
	list := make([]*entities.Category, 0, len(r.categories))
	for _, category := range r.categories {
		clone := *category
		list = append(list, &clone)
	}
	return list, nil
}

func (r *Inmem) CreateCategory(e *entities.Category) (*entities.Category, error) {
	clone := *e
	r.categories[e.ID] = &clone
	return &clone, nil
}

func (r *Inmem) UpdateCategory(e *entities.Category) (*entities.Category, error) {
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

func (r *Inmem) DeleteCategory(id entities.ID) error {
	if r.categories[id] == nil {
		return entities.ErrNotFound
	}
	r.categories[id] = nil
	return nil
}

func (r *Inmem) GetManufacturer(id entities.ID) (*entities.Manufacturer, error) {
	// Get entity
	e, ok := r.manufacturers[id]
	if !ok {
		return nil, entities.ErrNotFound
	}

	// Return clone
	clone := *e
	return &clone, nil
}

func (r *Inmem) ListManufacturers() ([]*entities.Manufacturer, error) {
	list := make([]*entities.Manufacturer, 0, len(r.manufacturers))
	for _, manufacturer := range r.manufacturers {
		clone := *manufacturer
		list = append(list, &clone)
	}
	return list, nil
}

func (r *Inmem) CreateManufacturer(e *entities.Manufacturer) (*entities.Manufacturer, error) {
	clone := *e
	r.manufacturers[e.ID] = &clone
	return &clone, nil
}

func (r *Inmem) UpdateManufacturer(e *entities.Manufacturer) (*entities.Manufacturer, error) {
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

func (r *Inmem) DeleteManufacturer(id entities.ID) error {
	if r.manufacturers[id] == nil {
		return entities.ErrNotFound
	}
	r.manufacturers[id] = nil
	return nil
}

func (r *Inmem) GetProduct(id entities.ID) (*entities.Product, error) {
	// Get entity
	e, ok := r.products[id]
	if !ok {
		return nil, entities.ErrNotFound
	}

	// Return clone
	clone := *e
	return &clone, nil
}

func (r *Inmem) ListProducts() ([]*entities.Product, error) {
	list := make([]*entities.Product, 0, len(r.products))
	for _, product := range r.products {
		clone := *product
		list = append(list, &clone)
	}
	return list, nil
}

func (r *Inmem) CreateProduct(e *entities.Product) (*entities.Product, error) {
	clone := *e
	r.products[e.ID] = &clone
	return &clone, nil
}

func (r *Inmem) UpdateProduct(e *entities.Product) (*entities.Product, error) {
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

func (r *Inmem) DeleteProduct(id entities.ID) error {
	if r.products[id] == nil {
		return entities.ErrNotFound
	}
	r.products[id] = nil
	return nil
}
