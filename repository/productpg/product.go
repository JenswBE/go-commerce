package productpg

import (
	"github.com/JenswBE/go-commerce/entity"
	"github.com/JenswBE/go-commerce/repository/productpg/internal"
	"github.com/google/uuid"
)

func (r *ProductPostgres) GetProduct(id entity.ID) (*entity.Product, error) {
	product := &internal.Product{}
	err := r.db.First(product, "id = ?", id).Error
	if err != nil {
		return nil, translatePgError(err, "product")
	}
	return internal.ProductPgToEntity(product), nil
}

func (r *ProductPostgres) ListProducts() ([]*entity.Product, error) {
	products := []*internal.Product{}
	err := r.db.Find(&products).Error
	if err != nil {
		return nil, translatePgError(err, "product")
	}
	return internal.ProductsListPgToEntity(products), nil
}

func (r *ProductPostgres) CreateProduct(e *entity.Product) (*entity.Product, error) {
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
		return nil, translatePgError(err, "product")
	}
	return internal.ProductPgToEntity(m), nil
}

func (r *ProductPostgres) UpdateProduct(e *entity.Product) (*entity.Product, error) {
	// Check if provided manufacturer ID exists
	if e.ManufacturerID != uuid.Nil {
		_, err := r.GetManufacturer(e.ManufacturerID)
		if err != nil {
			return nil, err
		}
	}

	// Update product
	m := internal.ProductEntityToPg(e)
	err := r.db.Model(m).Updates(m).Error
	if err != nil {
		return nil, translatePgError(err, "product")
	}
	return internal.ProductPgToEntity(m), nil
}

func (r *ProductPostgres) DeleteProduct(id entity.ID) error {
	err := r.db.Delete(&internal.Product{}, "id = ?", id).Error
	if err != nil {
		return translatePgError(err, "product")
	}
	return nil
}
