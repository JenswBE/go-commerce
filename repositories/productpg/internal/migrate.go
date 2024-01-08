package internal

import (
	"strings"
	"time"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			// Initial migration
			ID: "202107302030",
			Migrate: func(db *gorm.DB) error {
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
				return db.AutoMigrate(
					&Category{},
					&Manufacturer{},
					&Product{},
					&Image{},
				)
			},
		},
		{
			// Setup correct FK constraints
			ID: "202110281930",
			Migrate: func(db *gorm.DB) error {
				return runStatements(db, []string{
					"ALTER TABLE products DROP CONSTRAINT IF EXISTS fk_products_manufacturer",
					"ALTER TABLE products ADD CONSTRAINT fk_products_manufacturer FOREIGN KEY (manufacturer_id) REFERENCES manufacturers (id) ON UPDATE RESTRICT ON DELETE RESTRICT",
					"ALTER TABLE product_categories DROP CONSTRAINT IF EXISTS fk_product_categories_product",
					"ALTER TABLE product_categories ADD CONSTRAINT fk_product_categories_product FOREIGN KEY (product_id) REFERENCES products (id) ON UPDATE RESTRICT ON DELETE CASCADE",
					"ALTER TABLE product_categories DROP CONSTRAINT IF EXISTS fk_product_categories_category",
					"ALTER TABLE product_categories ADD CONSTRAINT fk_product_categories_category FOREIGN KEY (category_id) REFERENCES categories (id) ON UPDATE RESTRICT ON DELETE RESTRICT",
				})
			},
		},
		{
			// Add services and service categories
			ID: "202401051920",
			Migrate: func(db *gorm.DB) error {
				return runStatements(db, []string{
					joinStatement(
						`CREATE TABLE service_categories (`,
						`id uuid NOT NULL,`,
						`created_at timestamp with time zone,`,
						`updated_at timestamp with time zone,`,
						`name text NOT NULL,`,
						`"order" bigint NOT NULL,`,
						`CONSTRAINT "service_categories_pkey" PRIMARY KEY ("id")`,
						`);`,
					),
					joinStatement(
						`CREATE TABLE services (`,
						`id uuid NOT NULL,`,
						`created_at timestamp with time zone,`,
						`updated_at timestamp with time zone,`,
						`name text NOT NULL,`,
						`description text NOT NULL,`,
						`price bigint NOT NULL,`,
						`service_category_id uuid NOT NULL,`,
						`"order" bigint NOT NULL,`,
						`CONSTRAINT "services_pkey" PRIMARY KEY ("id")`,
						`);`,
					),
					"ALTER TABLE services ADD CONSTRAINT fk_service_service_category_id FOREIGN KEY (service_category_id) REFERENCES service_categories (id) ON UPDATE RESTRICT ON DELETE CASCADE NOT DEFERRABLE",
					`ALTER TABLE services ADD CONSTRAINT uniq_service_category_order UNIQUE (service_category_id,"order") DEFERRABLE INITIALLY IMMEDIATE`,
				})
			},
		},
	})

	// Run migrations
	return m.Migrate()
}

func joinStatement(parts ...string) string {
	return strings.Join(parts, " ")
}

func runStatements(db *gorm.DB, statements []string) error {
	return db.Transaction(func(tx *gorm.DB) error {
		for _, statement := range statements {
			if err := db.Exec(statement).Error; err != nil {
				return err
			}
		}
		return nil
	})
}
