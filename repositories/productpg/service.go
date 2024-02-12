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

func (r *ProductPostgres) GetService(id entities.ID) (*entities.Service, error) {
	svc := &internal.Service{}
	err := r.db.Take(svc, "id = ?", id.String()).Error
	if err != nil {
		return nil, translatePgError(err, svc, id.String())
	}
	return svc.ToEntity(), nil
}

func (r *ProductPostgres) ListServices(optionalServiceCategoryID entities.ID) ([]*entities.Service, error) {
	services := []*internal.Service{}
	query := r.db
	if !optionalServiceCategoryID.IsNil() {
		query = query.Where("service_category_id = ?", optionalServiceCategoryID)
	}
	err := query.Order(`"order"`).Find(&services).Error
	if err != nil {
		return nil, translatePgError(err, services, "")
	}
	return repositories.ToEntitiesList(services, (*internal.Service).ToEntity), nil
}

func (r *ProductPostgres) CreateService(e *entities.Service) (*entities.Service, error) {
	// Fetch currently highest order
	maxOrder, err := getMaxOrderForService(r.db, e.ServiceCategoryID)
	if err != nil {
		return nil, err
	}

	// Create service
	svc := internal.ServiceEntityToPg(e)
	svc.Order = maxOrder + 1
	err = r.db.Create(svc).Error
	if err != nil {
		return nil, translatePgError(err, svc, svc.ID)
	}
	return svc.ToEntity(), nil
}

func (r *ProductPostgres) UpdateService(e *entities.Service) (*entities.Service, error) {
	// Fetch service
	currentService, err := r.GetService(e.ID)
	if err != nil {
		return nil, err
	}

	var svc *internal.Service
	r.db = r.db.Debug()
	err = r.db.Transaction(func(tx *gorm.DB) error {
		if e.ServiceCategoryID == currentService.ServiceCategoryID {
			// Order within service category might have changed
			if currentService.Order != e.Order {
				// Fetch service with same new order (if any)
				var serviceSameOrder internal.Service
				err := r.db.Take(&serviceSameOrder, `service_category_id = ? AND "order" = ?`, e.ServiceCategoryID, e.Order).Error
				if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
					return translatePgError(err, serviceSameOrder, "")
				}

				// Defer all constraints
				err = tx.Exec("SET CONSTRAINTS ALL DEFERRED").Error
				if err != nil {
					return translatePgError(err, serviceSameOrder, serviceSameOrder.ID)
				}

				// Service with same new order exists => Update with old order
				if serviceSameOrder.Base.ID != "" {
					serviceSameOrder.Order = currentService.Order
					err = tx.Save(serviceSameOrder).Error
					if err != nil {
						return translatePgError(err, serviceSameOrder, serviceSameOrder.ID)
					}
				}
			}
		} else {
			// Service category changed => Append to last order
			maxOrder, err := getMaxOrderForService(r.db, e.ServiceCategoryID)
			if err != nil {
				return err
			}
			e.Order = maxOrder + 1
		}

		// Update service
		svc = internal.ServiceEntityToPg(e)
		err := tx.Save(svc).Error
		if err != nil {
			return translatePgError(err, svc, e.ID.String())
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return svc.ToEntity(), nil
}

func (r *ProductPostgres) DeleteService(id entities.ID) error {
	// Delete service
	var services []internal.Service
	err := r.db.
		Clauses(clause.Returning{Columns: []clause.Column{{Name: "id"}}}).
		Delete(&services, "id = ?", id.String()).
		Error
	if err != nil {
		return translatePgError(err, &internal.Service{}, id.String())
	}

	// Return error if not found
	if len(services) == 0 {
		return entities.NewError(404, openapi.GOCOMERRORCODE_UNKNOWN_SERVICE, id.String(), err)
	}
	return nil
}

// getMaxOrderForService fetches the currently maximum order number of a service for a service category.
func getMaxOrderForService(db *gorm.DB, serviceCategoryID entities.ID) (int, error) {
	var result struct{ MaxOrder int }
	query := db.
		Model(&internal.Service{}).
		Select(`MAX("order") AS max_order`).
		Where("service_category_id = ?", serviceCategoryID)
	err := query.Scan(&result).Error
	if err != nil {
		return 0, translatePgError(err, nil, "")
	}
	return result.MaxOrder, nil
}
