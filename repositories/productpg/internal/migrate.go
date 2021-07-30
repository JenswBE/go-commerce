package internal

import (
	"time"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		// create persons table
		{
			ID: "202107302030",
			Migrate: func(tx *gorm.DB) error {
				type Base struct {
					ID        string `gorm:"type:uuid"`
					CreatedAt time.Time
					UpdatedAt time.Time
				}
				type Category struct {
					Base
					Name        string
					Description string
					ParentID    *string    `gorm:"type:uuid"`
					Children    []Category `gorm:"foreignkey:ParentID"`
					Products    []*Product `gorm:"many2many:product_categories;"`
					Order       int
					Image       *Image `gorm:"polymorphic:Owner;"`
				}
				type Manufacturer struct {
					Base
					Name       string
					WebsiteURL string
					Image      *Image `gorm:"polymorphic:Owner;"`
				}
				type Product struct {
					Base
					Name             string
					DescriptionShort string
					DescriptionLong  string
					Price            int
					Categories       []Category `gorm:"many2many:product_categories;"`
					ManufacturerID   *string    `gorm:"type:uuid"`
					Status           string
					StockCount       int
					Images           []Image `gorm:"polymorphic:Owner;"`
				}
				type Image struct {
					ID        string `gorm:"type:uuid"`
					OwnerID   string `gorm:"type:uuid"`
					OwnerType string
					Extension string // File extension
					Order     int
				}
				return tx.AutoMigrate(
					&Category{},
					&Manufacturer{},
					&Product{},
					&Image{},
				)
			},
		},
	})

	// Run migrations
	return m.Migrate()
}
