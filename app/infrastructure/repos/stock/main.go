package stockrepo

import (
	"github.com/jiradeto/corn-kernels-backend/app/entities"
	"gorm.io/gorm"
)

// Repo ...
type Repo interface {
	CreateOneProductStock(tx *gorm.DB, input CreateOneProductStockInput) (*entities.ProductStock, error)
	FindOneProductStock(tx *gorm.DB, input FindOneProductStockInput) (*entities.ProductStock, error)
	UpdateOneProductStock(tx *gorm.DB, input UpdateOneProductStockInput) error
}

type repo struct {
	DB *gorm.DB
}

// New is a constructor method of Repo
func New(db *gorm.DB) Repo {
	return &repo{
		DB: db,
	}
}

func (repo *repo) selectDB(injectedDB *gorm.DB) *gorm.DB {
	if injectedDB == nil {
		return repo.DB
	}
	return injectedDB
}
