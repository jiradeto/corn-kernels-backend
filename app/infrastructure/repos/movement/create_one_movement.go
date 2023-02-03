package movementrepo

import (
	"fmt"

	"github.com/jiradeto/corn-kernels-backend/app/entities"
	"github.com/jiradeto/corn-kernels-backend/app/infrastructure/models"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// CreateOneStockMovementInput is a DTO for updating one movement
type CreateOneStockMovementInput struct {
	StockMovementEntity *entities.StockMovement
}

// CreateOneStockMovement is a function for updating movement
func (repo *repo) CreateOneStockMovement(tx *gorm.DB, input CreateOneStockMovementInput) (*entities.StockMovement, error) {
	const errLocation = "movementRepo/CreateOneStockMovement %s"

	stockMovementModel, err := new(models.StockMovement).FromEntity(input.StockMovementEntity)
	if err != nil {
		return nil, errors.Wrapf(err, errLocation, "unable to parse an entity to model")
	}

	query := repo.selectDB(tx)
	result := query.Create(stockMovementModel)
	if result.Error != nil {
		return nil, errors.Wrap(result.Error, fmt.Sprintf(errLocation, "unable to create Product due to database error"))
	}

	resultEntity, err := stockMovementModel.ToEntity()
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(errLocation, "unable to covert from model to entity"))
	}

	return resultEntity, nil
}
