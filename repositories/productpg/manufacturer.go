package productpg

import (
	"github.com/JenswBE/go-commerce/entities"
	"github.com/JenswBE/go-commerce/repositories/productpg/internal"
)

func (r *ProductPostgres) GetManufacturer(id entities.ID) (*entities.Manufacturer, error) {
	manufacturer := &internal.Manufacturer{}
	err := r.db.Preload("Image").Take(manufacturer, "id = ?", id).Error
	if err != nil {
		return nil, translatePgError(err, manufacturer, id.String())
	}
	return manufacturer.ToEntity(), nil
}

func (r *ProductPostgres) ListManufacturers() ([]*entities.Manufacturer, error) {
	manufacturers := []*internal.Manufacturer{}
	err := r.db.Preload("Image").Find(&manufacturers).Error
	if err != nil {
		return nil, translatePgError(err, manufacturers, "")
	}
	return internal.ManufacturersListPgToEntity(manufacturers), nil
}

func (r *ProductPostgres) CreateManufacturer(e *entities.Manufacturer) (*entities.Manufacturer, error) {
	m := internal.ManufacturerEntityToPg(e)
	err := r.db.Create(m).Error
	if err != nil {
		return nil, translatePgError(err, m, m.ID)
	}
	return m.ToEntity(), nil
}

func (r *ProductPostgres) UpdateManufacturer(e *entities.Manufacturer) (*entities.Manufacturer, error) {
	m := internal.ManufacturerEntityToPg(e)
	err := r.db.Save(m).Error
	if err != nil {
		return nil, translatePgError(err, m, m.ID)
	}
	return m.ToEntity(), nil
}

func (r *ProductPostgres) DeleteManufacturer(id entities.ID) error {
	err := r.db.Delete(&internal.Manufacturer{}, "id = ?", id).Error
	if err != nil {
		return translatePgError(err, &internal.Manufacturer{}, id.String())
	}
	return nil
}
