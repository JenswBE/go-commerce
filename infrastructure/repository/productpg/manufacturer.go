package productpg

import (
	"github.com/JenswBE/go-commerce/entity"
	"github.com/JenswBE/go-commerce/infrastructure/repository/productpg/internal"
)

func (r *ProductPostgres) GetManufacturer(id entity.ID) (*entity.Manufacturer, error) {
	manufacturer := &internal.Manufacturer{}
	err := r.db.First(manufacturer, "id = ?", id).Error
	if err != nil {
		return nil, translatePgError(err, "manufacturer")
	}
	return internal.ManufacturerPgToEntity(manufacturer), nil
}

func (r *ProductPostgres) ListManufacturers() ([]*entity.Manufacturer, error) {
	manufacturers := []*internal.Manufacturer{}
	err := r.db.Find(&manufacturers).Error
	if err != nil {
		return nil, translatePgError(err, "manufacturer")
	}
	return internal.ManufacturersListPgToEntity(manufacturers), nil
}

func (r *ProductPostgres) SearchManufacturers(query string) ([]*entity.Manufacturer, error) {
	manufacturers := []*internal.Manufacturer{}
	err := r.db.Where("name ILIKE ?", "%"+query+"%").Find(&manufacturers).Error
	if err != nil {
		return nil, translatePgError(err, "manufacturer")
	}
	return internal.ManufacturersListPgToEntity(manufacturers), nil
}

func (r *ProductPostgres) CreateManufacturer(e *entity.Manufacturer) (*entity.Manufacturer, error) {
	m := internal.ManufacturerEntityToPg(e)
	err := r.db.Create(m).Error
	if err != nil {
		return nil, translatePgError(err, "manufacturer")
	}
	return internal.ManufacturerPgToEntity(m), nil
}

func (r *ProductPostgres) UpdateManufacturer(e *entity.Manufacturer) (*entity.Manufacturer, error) {
	m := internal.ManufacturerEntityToPg(e)
	err := r.db.Model(m).Updates(m).Error
	if err != nil {
		return nil, translatePgError(err, "manufacturer")
	}
	return internal.ManufacturerPgToEntity(m), nil
}

func (r *ProductPostgres) DeleteManufacturer(id entity.ID) error {
	err := r.db.Delete(&internal.Manufacturer{}, "id = ?", id).Error
	if err != nil {
		return translatePgError(err, "manufacturer")
	}
	return nil
}
