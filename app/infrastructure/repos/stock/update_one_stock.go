package stockrepo

import (
	"fmt"

	"github.com/jiradeto/corn-kernels-backend/app/entities"
	"github.com/jiradeto/corn-kernels-backend/app/infrastructure/models"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// UpdateOneProductStockInput is a DTO for updating one movement
type UpdateOneProductStockInput struct {
	ProductStockEntity *entities.ProductStock
}

// UpdateOneProductStock is a function for updating stock
func (repo *repo) UpdateOneProductStock(tx *gorm.DB, input UpdateOneProductStockInput) error {
	const errLocation = "movementRepo/UpdateOneProductStock %s"

	productStockModel, err := new(models.ProductStock).FromEntity(input.ProductStockEntity)
	if err != nil {
		return errors.Wrapf(err, errLocation, "unable to parse an entity to model")
	}

	query := repo.selectDB(tx)
	result := query.Updates(&productStockModel)
	if result.Error != nil {
		return errors.Wrap(result.Error, fmt.Sprintf(errLocation, "unable to update due to database error"))
	}
	return nil
}
