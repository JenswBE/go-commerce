package productpg

import (
	"github.com/JenswBE/go-commerce/entities"
	"github.com/JenswBE/go-commerce/repositories/productpg/internal"
	"github.com/google/uuid"
)

func (r *ProductPostgres) GetProduct(id entities.ID) (*entities.Product, error) {
	product := &internal.Product{}
	err := r.db.Preload("Categories").Preload("Images").Take(product, "id = ?", id).Error
	if err != nil {
		return nil, translatePgError(err, product, id.String())
	}
	return internal.ProductPgToEntity(product), nil
}

func (r *ProductPostgres) ListProducts() ([]*entities.Product, error) {
	products := []*internal.Product{}
	err := r.db.Preload("Categories").Preload("Images").Find(&products).Error
	if err != nil {
		return nil, translatePgError(err, products, "")
	}
	return internal.ProductsListPgToEntity(products), nil
}

func (r *ProductPostgres) CreateProduct(e *entities.Product) (*entities.Product, error) {
	// Check if provided manufacturer ID exists
	if e.ManufacturerID != uuid.Nil {
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
	return internal.ProductPgToEntity(m), nil
}

func (r *ProductPostgres) UpdateProduct(e *entities.Product) (*entities.Product, error) {
	// Check if provided manufacturer ID exists
	if e.ManufacturerID != uuid.Nil {
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
	err = r.db.Model(m).Association("Images").Replace(m.Images)
	if err != nil {
		return nil, translatePgError(err, m, m.ID)
	}
	return internal.ProductPgToEntity(m), nil
}

func (r *ProductPostgres) DeleteProduct(id entities.ID) error {
	err := r.db.Select("Images").Delete(&internal.Product{}, "id = ?", id).Error
	if err != nil {
		return translatePgError(err, &internal.Product{}, id.String())
	}
	return nil
}
