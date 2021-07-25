package product

import (
	"github.com/JenswBE/go-commerce/entities"
	"github.com/JenswBE/go-commerce/utils/imageproxy"
)

type DatabaseRepository interface {
	GetCategory(id entities.ID) (*entities.Category, error)
	ListCategories() ([]*entities.Category, error)
	CreateCategory(e *entities.Category) (*entities.Category, error)
	UpdateCategory(e *entities.Category) (*entities.Category, error)
	DeleteCategory(id entities.ID) error

	GetManufacturer(id entities.ID) (*entities.Manufacturer, error)
	ListManufacturers() ([]*entities.Manufacturer, error)
	CreateManufacturer(e *entities.Manufacturer) (*entities.Manufacturer, error)
	UpdateManufacturer(e *entities.Manufacturer) (*entities.Manufacturer, error)
	DeleteManufacturer(id entities.ID) error

	GetProduct(id entities.ID) (*entities.Product, error)
	ListProducts() ([]*entities.Product, error)
	CreateProduct(e *entities.Product) (*entities.Product, error)
	UpdateProduct(e *entities.Product) (*entities.Product, error)
	DeleteProduct(id entities.ID) error

	GetImage(id entities.ID) (*entities.Image, error)
	UpdateImage(id entities.ID, ownerID entities.ID, newOrder int) ([]*entities.Image, error)
	DeleteImage(id entities.ID) error
}

type StorageRepository interface {
	SaveFile(filename string, content []byte) error
	DeleteFile(filename string) error
}

type Usecase interface {
	GetCategory(id entities.ID, imageConfig *imageproxy.ImageConfig) (*entities.Category, error)
	ListCategories(imageConfig *imageproxy.ImageConfig) ([]*entities.Category, error)
	CreateCategory(*entities.Category) (*entities.Category, error)
	UpdateCategory(e *entities.Category) (*entities.Category, error)
	DeleteCategory(id entities.ID) error
	UpsertCategoryImage(categoryID entities.ID, imageName string, imageContent []byte, imageConfig *imageproxy.ImageConfig) (*entities.Category, error)
	DeleteCategoryImage(categoryID entities.ID) error

	GetManufacturer(id entities.ID, imageConfig *imageproxy.ImageConfig) (*entities.Manufacturer, error)
	ListManufacturers(imageConfig *imageproxy.ImageConfig) ([]*entities.Manufacturer, error)
	CreateManufacturer(*entities.Manufacturer) (*entities.Manufacturer, error)
	UpdateManufacturer(e *entities.Manufacturer) (*entities.Manufacturer, error)
	DeleteManufacturer(id entities.ID) error
	UpsertManufacturerImage(manufacturerID entities.ID, imageName string, imageContent []byte, imageConfig *imageproxy.ImageConfig) (*entities.Manufacturer, error)
	DeleteManufacturerImage(manufacturerID entities.ID) error

	GetProduct(id entities.ID, imageConfig *imageproxy.ImageConfig) (*entities.Product, error)
	ListProducts(imageConfig *imageproxy.ImageConfig) ([]*entities.Product, error)
	CreateProduct(*entities.Product) (*entities.Product, error)
	UpdateProduct(e *entities.Product) (*entities.Product, error)
	DeleteProduct(id entities.ID) error
	AddProductImages(id entities.ID, images map[string][]byte, imageConfig *imageproxy.ImageConfig) (*entities.Product, error)
	UpdateProductImage(productID, imageID entities.ID, order int) ([]*entities.Image, error)
	DeleteProductImage(productID, imageID entities.ID) error
}
