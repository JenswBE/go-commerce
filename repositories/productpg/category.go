package productpg

import (
	"errors"
	"net/http"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/JenswBE/go-commerce/api/openapi"
	"github.com/JenswBE/go-commerce/entities"
	"github.com/JenswBE/go-commerce/repositories/productpg/internal"
)

func (r *ProductPostgres) GetCategory(id entities.ID) (*entities.Category, error) {
	category := &internal.Category{}
	err := r.db.Preload("Products").Preload("Image").Take(category, "id = ?", id.String()).Error
	if err != nil {
		return nil, translatePgError(err, category, id.String())
	}
	return category.ToEntity(), nil
}

func (r *ProductPostgres) ListCategories() ([]*entities.Category, error) {
	categories := []*internal.Category{}
	err := r.db.Preload("Products").Preload("Image").Order(`"order"`).Find(&categories).Error
	if err != nil {
		return nil, translatePgError(err, categories, "")
	}
	return internal.CategoriesListPgToEntity(categories), nil
}

func (r *ProductPostgres) CreateCategory(e *entities.Category) (*entities.Category, error) {
	// Check if provided parent ID exists
	if !e.ParentID.IsNil() {
		_, err := r.GetCategory(e.ParentID)
		if err != nil {
			return nil, err
		}
	}

	// Fetch currently highest order
	maxOrder, err := r.getMaxOrderForParentID(e.ParentID)
	if err != nil {
		return nil, err
	}

	// Create category
	m := internal.CategoryEntityToPg(e)
	m.Order = maxOrder + 1
	err = r.db.Create(m).Error
	if err != nil {
		return nil, translatePgError(err, m, m.ID)
	}
	return m.ToEntity(), nil
}

func (r *ProductPostgres) UpdateCategory(e *entities.Category) (*entities.Category, error) {
	// Check if provided parent ID exists
	if !e.ParentID.IsNil() {
		_, err := r.GetCategory(e.ParentID)
		if err != nil {
			return nil, err
		}
	}

	// Fetch category
	currentCategory, err := r.GetCategory(e.ID)
	if err != nil {
		return nil, err
	}

	var m *internal.Category
	err = r.db.Transaction(func(tx *gorm.DB) error {
		// Check if order has to be updated
		if currentCategory.Order != e.Order {
			// Fetch category with same new order (if any)
			categorySameOrder, err := r.getCategoryByParentIDAndOrder(e.ParentID, e.Order)
			if err != nil {
				gocomErr, ok := err.(*entities.GoComError)
				if !ok || gocomErr.Status == http.StatusNotFound {
					return err
				}
			}

			// Image with same new order exists => Update with old order
			if categorySameOrder != nil {
				categorySameOrder.Order = currentCategory.Order
				err = tx.Save(categorySameOrder).Error
				if err != nil {
					return translatePgError(err, categorySameOrder, categorySameOrder.ID)
				}
			}
		}

		// Update category
		m = internal.CategoryEntityToPg(e)
		err := tx.Save(m).Error
		if err != nil {
			return translatePgError(err, m, e.ID.String())
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return m.ToEntity(), nil
}

func (r *ProductPostgres) DeleteCategory(id entities.ID) error {
	// Delete category
	var categories []internal.Category
	err := r.db.
		Clauses(clause.Returning{Columns: []clause.Column{{Name: "id"}}}).
		Delete(&categories, "id = ?", id.String()).
		Error
	if err != nil {
		return translatePgError(err, &internal.Category{}, id.String())
	}

	// Return error if not found
	if len(categories) == 0 {
		return entities.NewError(404, openapi.GOCOMERRORCODE_UNKNOWN_CATEGORY, id.String(), err)
	}
	return nil
}

func (r *ProductPostgres) getCategoryByParentIDAndOrder(parentID entities.ID, order int) (*internal.Category, error) {
	category := &internal.Category{}
	conditions := internal.NewConditionsGroup(`"order" = ?`, order)
	if parentID.IsNil() {
		conditions.AddANDCondition("parent_id IS NULL")
	} else {
		conditions.AddANDCondition("parent_id = ?", parentID.String())
	}
	err := takeWithConditions(r.db, category, conditions)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, translatePgError(err, category, "")
	}
	return category, nil
}

// getMaxOrderForParentID fetches the currently maximum order number of a parent.
func (r *ProductPostgres) getMaxOrderForParentID(parentID entities.ID) (int, error) {
	query := r.db.Model(&internal.Category{}).Select(`MAX("order") AS max_order`)
	if parentID.IsNil() {
		query = query.Where("parent_id IS NULL")
	} else {
		query = query.Where("parent_id = ?", parentID.String())
	}

	var result struct{ MaxOrder int }
	err := query.Scan(&result).Error
	if err != nil {
		return 0, translatePgError(err, nil, "")
	}
	return result.MaxOrder, nil
}
