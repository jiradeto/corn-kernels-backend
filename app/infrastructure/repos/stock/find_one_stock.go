package stockrepo

import (
	"fmt"

	"github.com/jiradeto/corn-kernels-backend/app/constants"
	"github.com/jiradeto/corn-kernels-backend/app/entities"
	"github.com/jiradeto/corn-kernels-backend/app/infrastructure/models"
	"github.com/jiradeto/corn-kernels-backend/app/utils/gerrors"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// FindOneProductStockInput is an input for find one item
type FindOneProductStockInput struct {
	ID *string
}

func (repo *repo) FindOneProductStock(tx *gorm.DB, input FindOneProductStockInput) (*entities.ProductStock, error) {
	const errLocation = "stockRepo/FindOneProductStock: %s"
	var resultModel models.ProductStock

	query := repo.selectDB(tx)

	if input.ID != nil {
		query = query.Where(`id = ?`, *input.ID)
	}

	result := query.First(&resultModel)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, gerrors.RecordNotFoundError{
				Code:    constants.StatusCodeEntryNotFound,
				Message: constants.ErrorMessageNotFound,
			}
		}
		return nil, errors.Wrap(result.Error, fmt.Sprintf(errLocation, "unable to find item due to database error"))
	}

	resultEntity, err := resultModel.ToEntity()
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(errLocation, "unable to covert from model to entity"))
	}

	return resultEntity, nil
}
