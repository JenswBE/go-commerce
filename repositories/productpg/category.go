package productpg

import (
	"github.com/JenswBE/go-commerce/api/openapi"
	"github.com/JenswBE/go-commerce/entities"
	"github.com/JenswBE/go-commerce/repositories/productpg/internal"
	"github.com/google/uuid"
	"gorm.io/gorm/clause"
)

func (r *ProductPostgres) GetCategory(id entities.ID) (*entities.Category, error) {
	category := &internal.Category{}
	err := r.db.Preload("Products").Preload("Image").Take(category, "id = ?", id).Error
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
	if e.ParentID != uuid.Nil {
		_, err := r.GetCategory(e.ParentID)
		if err != nil {
			return nil, err
		}
	}

	// Create category
	m := internal.CategoryEntityToPg(e)
	err := r.db.Create(m).Error
	if err != nil {
		return nil, translatePgError(err, m, m.ID)
	}
	return m.ToEntity(), nil
}

func (r *ProductPostgres) UpdateCategory(e *entities.Category) (*entities.Category, error) {
	// Check if provided parent ID exists
	if e.ParentID != uuid.Nil {
		_, err := r.GetCategory(e.ParentID)
		if err != nil {
			return nil, err
		}
	}

	// Update category
	m := internal.CategoryEntityToPg(e)
	err := r.db.Save(m).Error
	if err != nil {
		return nil, translatePgError(err, m, e.ID.String())
	}
	return m.ToEntity(), nil
}

func (r *ProductPostgres) DeleteCategory(id entities.ID) error {
	// Delete category
	var categories []internal.Category
	err := r.db.
		Clauses(clause.Returning{Columns: []clause.Column{{Name: "id"}}}).
		Delete(&categories, "id = ?", id).
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
