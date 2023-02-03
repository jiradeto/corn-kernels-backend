package movementrepo

import (
	"github.com/jiradeto/corn-kernels-backend/app/entities"
	"gorm.io/gorm"
)

// Repo ...
type Repo interface {
	CreateOneStockMovement(tx *gorm.DB, input CreateOneStockMovementInput) (*entities.StockMovement, error)
	GetMovementList(tx *gorm.DB, input GetMovementListInput) ([]*entities.StockMovement, error)
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
