package productpg

import (
	"github.com/JenswBE/go-commerce/entities"
	"github.com/JenswBE/go-commerce/repositories/productpg/internal"
	"github.com/google/uuid"
)

func (r *ProductPostgres) GetCategory(id entities.ID) (*entities.Category, error) {
	category := &internal.Category{}
	err := r.db.Preload("Products").Take(category, "id = ?", id).Error
	if err != nil {
		return nil, translatePgError(err, "category")
	}
	return internal.CategoryPgToEntity(category), nil
}

func (r *ProductPostgres) ListCategories() ([]*entities.Category, error) {
	categories := []*internal.Category{}
	err := r.db.Preload("Products").Find(&categories).Error
	if err != nil {
		return nil, translatePgError(err, "category")
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
		return nil, translatePgError(err, "category")
	}
	return internal.CategoryPgToEntity(m), nil
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
		return nil, translatePgError(err, "category")
	}
	return internal.CategoryPgToEntity(m), nil
}

func (r *ProductPostgres) DeleteCategory(id entities.ID) error {
	err := r.db.Delete(&internal.Category{}, "id = ?", id).Error
	if err != nil {
		return translatePgError(err, "category")
	}
	return nil
}
