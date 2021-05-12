package product

import "github.com/JenswBE/go-commerce/entity"

type Reader interface {
	GetCategory(id entity.ID) (*entity.Category, error)
	ListCategories() ([]*entity.Category, error)

	GetManufacturer(id entity.ID) (*entity.Manufacturer, error)
	ListManufacturers() ([]*entity.Manufacturer, error)
	SearchManufacturers(query string) ([]*entity.Manufacturer, error)
}

type Writer interface {
	CreateCategory(e *entity.Category) (*entity.Category, error)
	UpdateCategory(e *entity.Category) (*entity.Category, error)
	DeleteCategory(id entity.ID) error

	CreateManufacturer(e *entity.Manufacturer) (*entity.Manufacturer, error)
	UpdateManufacturer(e *entity.Manufacturer) (*entity.Manufacturer, error)
	DeleteManufacturer(id entity.ID) error
}

type Repository interface {
	Reader
	Writer
}

type Usecase interface {
	GetCategory(id entity.ID) (*entity.Category, error)
	ListCategories() ([]*entity.Category, error)
	CreateCategory(name, description string, parentID entity.ID) (*entity.Category, error)
	UpdateCategory(e *entity.Category) (*entity.Category, error)
	DeleteCategory(id entity.ID) error

	GetManufacturer(id entity.ID) (*entity.Manufacturer, error)
	ListManufacturers() ([]*entity.Manufacturer, error)
	SearchManufacturers(query string) ([]*entity.Manufacturer, error)
	CreateManufacturer(name, websiteURL string) (*entity.Manufacturer, error)
	UpdateManufacturer(e *entity.Manufacturer) (*entity.Manufacturer, error)
	DeleteManufacturer(id entity.ID) error
}
