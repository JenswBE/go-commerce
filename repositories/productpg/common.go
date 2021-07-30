package productpg

import (
	"errors"
	"fmt"

	"github.com/JenswBE/go-commerce/entities"
	"github.com/JenswBE/go-commerce/repositories/productpg/internal"
	"gorm.io/gorm"
)

type ProductPostgres struct {
	db *gorm.DB
}

func NewProductPostgres(db *gorm.DB) (*ProductPostgres, error) {
	// Migrate database
	err := internal.Migrate(db)
	if err != nil {
		return nil, err
	}

	// Build repository
	return &ProductPostgres{db: db}, nil
}

// translatePgError converts well-known errors (e.g. ErrRecordNotFound)
// to a more specific GoComError. Otherwise, provided error is returned
// as-is.
func translatePgError(err error, object string) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return entities.NewError(404, fmt.Errorf(`%s not found`, object))
	}
	return err
}