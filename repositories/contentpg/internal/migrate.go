package internal

import (
	"time"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			// Initial migration
			ID: "202111201700",
			Migrate: func(db *gorm.DB) error {
				type Base struct {
					ID        string `gorm:"type:uuid"`
					CreatedAt time.Time
					UpdatedAt time.Time
				}
				type Content struct {
					Name        string `gorm:"primaryKey"`
					ContentType string
					Content     string
				}
				type Event struct {
					Base
					Name        string
					Description string
					EventType   string
					Start       time.Time
					End         time.Time
					WholeDay    bool
				}
				return db.AutoMigrate(
					&Content{},
					&Event{},
				)
			},
		},
	})

	// Run migrations
	return m.Migrate()
}
