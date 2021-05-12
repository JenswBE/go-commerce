package productpg

import (
	"log"

	"github.com/JenswBE/go-commerce/entity"
	"github.com/JenswBE/go-commerce/infrastructure/repository/productpg/internal"
	"github.com/google/uuid"
)

func (r *ProductPostgres) GetCategory(id entity.ID) (*entity.Category, error) {
	category := &internal.Category{}
	err := r.db.First(category, "id = ?", id).Error
	if err != nil {
		return nil, translatePgError(err, "category")
	}
	return internal.CategoryPgToEntity(category), nil
}

func (r *ProductPostgres) ListCategories() ([]*entity.Category, error) {
	categories := []*internal.Category{}
	err := r.db.Find(&categories).Error
	if err != nil {
		return nil, translatePgError(err, "category")
	}
	return internal.CategoriesListPgToEntity(categories), nil
}

func (r *ProductPostgres) CreateCategory(e *entity.Category) (*entity.Category, error) {
	// Check if provided parent ID exists
	if e.ParentID != uuid.Nil {
		_, err := r.GetCategory(e.ParentID)
		if err != nil {
			return nil, translatePgError(err, "category")
		}
	}

	// Create category
	m := internal.CategoryEntityToPg(e)
	err := r.db.Create(m).Error
	log.Printf("CreateCategory error type: %T", err)
	if err != nil {
		return nil, translatePgError(err, "category")
	}
	return internal.CategoryPgToEntity(m), nil
}

func (r *ProductPostgres) UpdateCategory(e *entity.Category) (*entity.Category, error) {
	// Check if provided parent ID exists
	if e.ParentID != uuid.Nil {
		_, err := r.GetCategory(e.ParentID)
		if err != nil {
			return nil, translatePgError(err, "category")
		}
	}

	// Update category
	m := internal.CategoryEntityToPg(e)
	err := r.db.Model(m).Updates(m).Error
	if err != nil {
		return nil, translatePgError(err, "category")
	}
	return internal.CategoryPgToEntity(m), nil
}

func (r *ProductPostgres) DeleteCategory(id entity.ID) error {
	err := r.db.Delete(&internal.Category{}, "id = ?", id).Error
	if err != nil {
		return translatePgError(err, "category")
	}
	return nil
}
