package productpg

import (
	"github.com/JenswBE/go-commerce/entities"
	"github.com/JenswBE/go-commerce/repositories/productpg/internal"
)

func (r *ProductPostgres) GetManufacturer(id entities.ID) (*entities.Manufacturer, error) {
	manufacturer := &internal.Manufacturer{}
	err := r.db.Preload("Image").Take(manufacturer, "id = ?", id).Error
	if err != nil {
		return nil, translatePgError(err, "manufacturer")
	}
	return internal.ManufacturerPgToEntity(manufacturer), nil
}

func (r *ProductPostgres) ListManufacturers() ([]*entities.Manufacturer, error) {
	manufacturers := []*internal.Manufacturer{}
	err := r.db.Preload("Image").Find(&manufacturers).Error
	if err != nil {
		return nil, translatePgError(err, "manufacturer")
	}
	return internal.ManufacturersListPgToEntity(manufacturers), nil
}

func (r *ProductPostgres) CreateManufacturer(e *entities.Manufacturer) (*entities.Manufacturer, error) {
	m := internal.ManufacturerEntityToPg(e)
	err := r.db.Create(m).Error
	if err != nil {
		return nil, translatePgError(err, "manufacturer")
	}
	return internal.ManufacturerPgToEntity(m), nil
}

func (r *ProductPostgres) UpdateManufacturer(e *entities.Manufacturer) (*entities.Manufacturer, error) {
	m := internal.ManufacturerEntityToPg(e)
	err := r.db.Save(m).Error
	if err != nil {
		return nil, translatePgError(err, "manufacturer")
	}
	return internal.ManufacturerPgToEntity(m), nil
}

func (r *ProductPostgres) DeleteManufacturer(id entities.ID) error {
	err := r.db.Delete(&internal.Manufacturer{}, "id = ?", id).Error
	if err != nil {
		return translatePgError(err, "manufacturer")
	}
	return nil
}
