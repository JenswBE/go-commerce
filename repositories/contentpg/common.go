package contentpg

import (
	"errors"
	"reflect"

	"github.com/JenswBE/go-commerce/api/openapi"
	"github.com/JenswBE/go-commerce/entities"
	"github.com/JenswBE/go-commerce/repositories/contentpg/internal"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type ContentPostgres struct {
	db *gorm.DB
}

func NewContentPostgres(db *gorm.DB) (*ContentPostgres, error) {
	// Migrate database
	err := internal.Migrate(db)
	if err != nil {
		return nil, err
	}

	// Build repository
	return &ContentPostgres{db: db}, nil
}

// translatePgError converts well-known errors (e.g. ErrRecordNotFound)
// to a more specific GoComError. Otherwise, provided error is returned
// as-is.
func translatePgError(err error, object interface{}, instance string) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		switch object.(type) {
		case *internal.Content, []*internal.Content:
			return entities.NewError(404, openapi.GOCOMERRORCODE_UNKNOWN_CONTENT, instance, err)
		case *internal.Event, []*internal.Event:
			return entities.NewError(404, openapi.GOCOMERRORCODE_UNKNOWN_EVENT, instance, err)
		default:
			log.Warn().Err(err).Stringer("object", reflect.TypeOf(object)).Msg("Unknown object in translatePgError.ErrRecordNotFound")
			return entities.NewError(404, openapi.GOCOMERRORCODE_UNKNOWN_ERROR, instance, err)
		}
	}
	return err
}
