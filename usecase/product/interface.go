package product

import "github.com/JenswBE/go-commerce/entity"

type Reader interface {
	GetManufacturer(id entity.ID) (*entity.Manufacturer, error)
	ListManufacturers() ([]*entity.Manufacturer, error)
	SearchManufacturers(query string) ([]*entity.Manufacturer, error)
}

type Writer interface {
	CreateManufacturer(e *entity.Manufacturer) (*entity.Manufacturer, error)
	UpdateManufacturer(e *entity.Manufacturer) (*entity.Manufacturer, error)
	DeleteManufacturer(id entity.ID) error
}

type Repository interface {
	Reader
	Writer
}

type Usecase interface {
	GetManufacturer(id entity.ID) (*entity.Manufacturer, error)
	ListManufacturers() ([]*entity.Manufacturer, error)
	SearchManufacturers(query string) ([]*entity.Manufacturer, error)
	CreateManufacturer(name, websiteURL string) (*entity.Manufacturer, error)
	UpdateManufacturer(e *entity.Manufacturer) (*entity.Manufacturer, error)
	DeleteManufacturer(id entity.ID) error
}
