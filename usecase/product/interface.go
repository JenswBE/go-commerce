package product

import "github.com/JenswBE/go-commerce/entity"

type DatabaseRepository interface {
	GetCategory(id entity.ID) (*entity.Category, error)
	ListCategories() ([]*entity.Category, error)
	CreateCategory(e *entity.Category) (*entity.Category, error)
	UpdateCategory(e *entity.Category) (*entity.Category, error)
	DeleteCategory(id entity.ID) error

	GetManufacturer(id entity.ID) (*entity.Manufacturer, error)
	ListManufacturers() ([]*entity.Manufacturer, error)
	CreateManufacturer(e *entity.Manufacturer) (*entity.Manufacturer, error)
	UpdateManufacturer(e *entity.Manufacturer) (*entity.Manufacturer, error)
	DeleteManufacturer(id entity.ID) error

	GetProduct(id entity.ID) (*entity.Product, error)
	ListProducts() ([]*entity.Product, error)
	CreateProduct(e *entity.Product) (*entity.Product, error)
	UpdateProduct(e *entity.Product) (*entity.Product, error)
	DeleteProduct(id entity.ID) error
}

type StorageRepository interface {
	SaveFile(filename string, content []byte) error
	DeleteFile(filename string) error
}

type Usecase interface {
	GetCategory(id entity.ID) (*entity.Category, error)
	ListCategories() ([]*entity.Category, error)
	CreateCategory(*entity.Category) (*entity.Category, error)
	UpdateCategory(e *entity.Category) (*entity.Category, error)
	DeleteCategory(id entity.ID) error

	GetManufacturer(id entity.ID) (*entity.Manufacturer, error)
	ListManufacturers() ([]*entity.Manufacturer, error)
	CreateManufacturer(*entity.Manufacturer) (*entity.Manufacturer, error)
	UpdateManufacturer(e *entity.Manufacturer) (*entity.Manufacturer, error)
	DeleteManufacturer(id entity.ID) error

	GetProduct(id entity.ID) (*entity.Product, error)
	ListProducts() ([]*entity.Product, error)
	CreateProduct(*entity.Product) (*entity.Product, error)
	UpdateProduct(e *entity.Product) (*entity.Product, error)
	DeleteProduct(id entity.ID) error
	AddProductImages(id entity.ID, images map[string][]byte) (*entity.Product, error)
}
