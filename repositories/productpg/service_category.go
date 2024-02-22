package productpg

import (
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/JenswBE/go-commerce/api/openapi"
	"github.com/JenswBE/go-commerce/entities"
	"github.com/JenswBE/go-commerce/repositories"
	"github.com/JenswBE/go-commerce/repositories/productpg/internal"
)

func (r *ProductPostgres) GetServiceCategory(id entities.ID) (*entities.ServiceCategory, error) {
	category := &internal.ServiceCategory{}
	err := r.db.
		Preload("Services", func(db *gorm.DB) *gorm.DB {
			return db.Order(`"order"`)
		}).
		Take(category, "id = ?", id.String()).Error
	if err != nil {
		return nil, translatePgError(err, category, id.String())
	}
	return category.ToEntity(), nil
}

func (r *ProductPostgres) ListServiceCategories() ([]*entities.ServiceCategory, error) {
	categories := []*internal.ServiceCategory{}
	err := r.db.
		Preload("Services", func(db *gorm.DB) *gorm.DB {
			return db.Order(`"order"`)
		}).
		Order(`"order"`).Find(&categories).Error
	if err != nil {
		return nil, translatePgError(err, categories, "")
	}
	return repositories.ToEntitiesList(categories, (*internal.ServiceCategory).ToEntity), nil
}

func (r *ProductPostgres) CreateServiceCategory(e *entities.ServiceCategory) (*entities.ServiceCategory, error) {
	// Fetch currently highest order
	var result struct{ MaxOrder int }
	query := r.db.Model(&internal.ServiceCategory{}).Select(`MAX("order") AS max_order`)
	err := query.Scan(&result).Error
	if err != nil {
		return nil, translatePgError(err, nil, "")
	}

	// Create category
	m := internal.ServiceCategoryEntityToPg(e)
	m.Order = result.MaxOrder + 1
	err = r.db.Create(m).Error
	if err != nil {
		return nil, translatePgError(err, m, m.ID)
	}
	return m.ToEntity(), nil
}

func (r *ProductPostgres) UpdateServiceCategory(e *entities.ServiceCategory) (*entities.ServiceCategory, error) {
	// Fetch service category
	currentServiceCategory, err := r.GetServiceCategory(e.ID)
	if err != nil {
		return nil, err
	}

	var cat *internal.ServiceCategory
	err = r.db.Transaction(func(tx *gorm.DB) error {
		// Check if order has to be updated
		if currentServiceCategory.Order != e.Order {
			// Fetch category with same new order (if any)
			var categorySameOrder internal.ServiceCategory
			err := r.db.Take(&categorySameOrder, `"order" = ?`, e.Order).Error
			if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
				return translatePgError(err, categorySameOrder, "")
			}

			// Category with same new order exists => Update with old order
			if categorySameOrder.Base.ID != "" {
				categorySameOrder.Order = currentServiceCategory.Order
				err = tx.Save(categorySameOrder).Error
				if err != nil {
					return translatePgError(err, categorySameOrder, categorySameOrder.ID)
				}
			}
		}

		// Update category
		cat = internal.ServiceCategoryEntityToPg(e)
		err = tx.Save(cat).Error
		if err != nil {
			return translatePgError(err, cat, e.ID.String())
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return cat.ToEntity(), nil
}

func (r *ProductPostgres) DeleteServiceCategory(id entities.ID) error {
	// Delete category
	var categories []internal.ServiceCategory
	err := r.db.
		Clauses(clause.Returning{Columns: []clause.Column{{Name: "id"}}}).
		Delete(&categories, "id = ?", id.String()).
		Error
	if err != nil {
		return translatePgError(err, &internal.ServiceCategory{}, id.String())
	}

	// Return error if not found
	if len(categories) == 0 {
		return entities.NewError(404, openapi.GOCOMERRORCODE_UNKNOWN_SERVICE_CATEGORY, id.String(), err)
	}
	return nil
}
