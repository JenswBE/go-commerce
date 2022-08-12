package productpg

import (
	"github.com/JenswBE/go-commerce/api/openapi"
	"github.com/JenswBE/go-commerce/entities"
	"github.com/JenswBE/go-commerce/repositories/productpg/internal"
	"gorm.io/gorm/clause"
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
	err := r.db.Preload("Image").Order("name").Find(&manufacturers).Error
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
	// Delete manufacturer
	var manufacturers []internal.Manufacturer
	err := r.db.
		Clauses(clause.Returning{Columns: []clause.Column{{Name: "id"}}}).
		Delete(&manufacturers, "id = ?", id).
		Error
	if err != nil {
		return translatePgError(err, &internal.Manufacturer{}, id.String())
	}

	// Return error if not found
	if len(manufacturers) == 0 {
		return entities.NewError(404, openapi.GOCOMERRORCODE_UNKNOWN_MANUFACTURER, id.String(), err)
	}
	return nil
}
