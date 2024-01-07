package productpg

import (
	"errors"
	"reflect"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"

	"github.com/JenswBE/go-commerce/api/openapi"
	"github.com/JenswBE/go-commerce/entities"
	"github.com/JenswBE/go-commerce/repositories/productpg/internal"
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
func translatePgError(err error, object any, instance string) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		switch object.(type) {
		case *internal.Category, []*internal.Category:
			return entities.NewError(404, openapi.GOCOMERRORCODE_UNKNOWN_CATEGORY, instance, err)
		case *internal.Image, []*internal.Image:
			return entities.NewError(404, openapi.GOCOMERRORCODE_UNKNOWN_IMAGE, instance, err)
		case *internal.Manufacturer, []*internal.Manufacturer:
			return entities.NewError(404, openapi.GOCOMERRORCODE_UNKNOWN_MANUFACTURER, instance, err)
		case *internal.Product, []*internal.Product:
			return entities.NewError(404, openapi.GOCOMERRORCODE_UNKNOWN_PRODUCT, instance, err)
		case *internal.Service, []*internal.Service:
			return entities.NewError(404, openapi.GOCOMERRORCODE_UNKNOWN_SERVICE, instance, err)
		case *internal.ServiceCategory, []*internal.ServiceCategory:
			return entities.NewError(404, openapi.GOCOMERRORCODE_UNKNOWN_SERVICE_CATEGORY, instance, err)
		default:
			log.Warn().Err(err).Stringer("object", reflect.TypeOf(object)).Msg("Unknown object in translatePgError.ErrRecordNotFound")
			return entities.NewError(404, openapi.GOCOMERRORCODE_UNKNOWN_ERROR, instance, err)
		}
	}
	return err
}

// takeWithConditions fetches a single DB entry with the provided conditions and args.
// Provided conditions are AND-ed.
func takeWithConditions(db *gorm.DB, dest any, conditions *internal.ConditionsGroup) error {
	return db.Take(dest, conditions.GetGORMConds()...).Error
}
