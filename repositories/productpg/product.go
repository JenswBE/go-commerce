package productpg

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/JenswBE/go-commerce/api/openapi"
	"github.com/JenswBE/go-commerce/entities"
	"github.com/JenswBE/go-commerce/repositories"
	"github.com/JenswBE/go-commerce/repositories/productpg/internal"
)

func (r *ProductPostgres) GetProduct(id entities.ID) (*entities.Product, error) {
	product := &internal.Product{}
	err := r.db.
		Preload("Categories", func(db *gorm.DB) *gorm.DB {
			return db.Order(`"order"`)
		}).
		Preload("Images", func(db *gorm.DB) *gorm.DB {
			return db.Order(`"order"`)
		}).
		Take(product, "id = ?", id.String()).Error
	if err != nil {
		return nil, translatePgError(err, product, id.String())
	}
	return product.ToEntity(), nil
}

func (r *ProductPostgres) ListProducts() ([]*entities.Product, error) {
	products := []*internal.Product{}
	err := r.db.
		Preload("Categories", func(db *gorm.DB) *gorm.DB {
			return db.Order(`"order"`)
		}).
		Preload("Images", func(db *gorm.DB) *gorm.DB {
			return db.Order(`"order"`)
		}).
		Order("LOWER(name)").
		Find(&products).Error
	if err != nil {
		return nil, translatePgError(err, products, "")
	}
	return repositories.ToEntitiesList(products, (*internal.Product).ToEntity), nil
}

func (r *ProductPostgres) CreateProduct(e *entities.Product) (*entities.Product, error) {
	// Check if provided manufacturer ID exists
	if !e.ManufacturerID.IsNil() {
		_, err := r.GetManufacturer(e.ManufacturerID)
		if err != nil {
			return nil, err
		}
	}

	// Create product
	m := internal.ProductEntityToPg(e)
	err := r.db.Create(m).Error
	if err != nil {
		return nil, translatePgError(err, m, m.ID)
	}
	return m.ToEntity(), nil
}

func (r *ProductPostgres) UpdateProduct(e *entities.Product) (*entities.Product, error) {
	// Check if provided manufacturer ID exists
	if !e.ManufacturerID.IsNil() {
		_, err := r.GetManufacturer(e.ManufacturerID)
		if err != nil {
			return nil, err
		}
	}

	// Update product
	m := internal.ProductEntityToPg(e)
	err := r.db.Save(m).Error
	if err != nil {
		return nil, translatePgError(err, m, m.ID)
	}
	err = r.db.Model(m).Association("Categories").Replace(m.Categories)
	if err != nil {
		return nil, translatePgError(err, m, m.ID)
	}
	return m.ToEntity(), nil
}

func (r *ProductPostgres) UpdateProductImages(id entities.ID, images []*entities.Image) ([]*entities.Image, error) {
	m := &internal.Product{Base: internal.Base{ID: id.String()}}
	mImages := internal.ImagesListEntityToPg(images)
	err := r.db.Model(m).Association("Images").Replace(mImages)
	if err != nil {
		return nil, translatePgError(err, m, m.ID)
	}
	return internal.ImagesListPgToEntity(mImages), nil
}

func (r *ProductPostgres) DeleteProduct(id entities.ID) error {
	// Delete product
	var products []internal.Product
	err := r.db.
		Select("Images").
		Clauses(clause.Returning{Columns: []clause.Column{{Name: "id"}}}).
		Delete(&products, "id = ?", id.String()).
		Error
	if err != nil {
		return translatePgError(err, &internal.Product{}, id.String())
	}

	// Return error if not found
	if len(products) == 0 {
		return entities.NewError(404, openapi.GOCOMERRORCODE_UNKNOWN_PRODUCT, id.String(), err)
	}
	return nil
}
