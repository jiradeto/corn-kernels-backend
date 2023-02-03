package stockrepo

import (
	"fmt"

	"github.com/jiradeto/corn-kernels-backend/app/entities"
	"github.com/jiradeto/corn-kernels-backend/app/infrastructure/models"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// CreateOneProductStockInput is a DTO for updating one Stock
type CreateOneProductStockInput struct {
	ProductStockEntity *entities.ProductStock
}

// CreateOneProductStock is a function for updating Stock
func (repo *repo) CreateOneProductStock(tx *gorm.DB, input CreateOneProductStockInput) (*entities.ProductStock, error) {
	const errLocation = "stockRepo/CreateOneProductStock %s"

	ProductStockModel, err := new(models.ProductStock).FromEntity(input.ProductStockEntity)
	if err != nil {
		return nil, errors.Wrapf(err, errLocation, "unable to parse an entity to model")
	}

	query := repo.selectDB(tx)
	result := query.Create(ProductStockModel)
	if result.Error != nil {
		return nil, errors.Wrap(result.Error, fmt.Sprintf(errLocation, "unable to create Product due to database error"))
	}

	resultEntity, err := ProductStockModel.ToEntity()
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(errLocation, "unable to covert from model to entity"))
	}

	return resultEntity, nil
}
