package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var migrate20221215112000CreateCornKernelsTables = []string{
	`
	CREATE TABLE products (
		id INTEGER PRIMARY KEY,
		name TEXT NOT NULL,
		description TEXT,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
		updated_at DATETIME DEFAULT (STRFTIME('%Y-%m-%d %H:%M:%f', 'NOW')) NOT NULL
	);
	CREATE TABLE product_stocks (
		id INTEGER PRIMARY KEY,
		product_id INTEGER NOT NULL,
		quantity INTEGER NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
		updated_at DATETIME DEFAULT (STRFTIME('%Y-%m-%d %H:%M:%f', 'NOW')) NOT NULL,
		FOREIGN KEY (product_id) REFERENCES products(id)
	);
	
	CREATE TABLE stock_movements (
		id INTEGER PRIMARY KEY,
		product_id INTEGER NOT NULL,
		description TEXT,
		quantity INTEGER NOT NULL,
		movement_type TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
		updated_at DATETIME DEFAULT (STRFTIME('%Y-%m-%d %H:%M:%f', 'NOW')) NOT NULL,
		FOREIGN KEY (product_id) REFERENCES products(id)
	);
	`,
}

var migrate20221215112000DropCornKernelsTables = []string{
	`DROP TABLE "products"`,
	`DROP TABLE "product_stocks"`,
	`DROP TABLE "stock_movements"`,
}

func migrate20221215112000() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "20221215112000",
		Migrate: func(db *gorm.DB) error {
			for _, migrate := range migrate20221215112000CreateCornKernelsTables {
				if err := db.Exec(migrate).Error; err != nil {
					return err
				}
			}
			return nil
		},
		Rollback: func(db *gorm.DB) error {
			for _, migrate := range migrate20221215112000DropCornKernelsTables {
				if err := db.Exec(migrate).Error; err != nil {
					return err
				}
			}
			return nil
		},
	}
}
