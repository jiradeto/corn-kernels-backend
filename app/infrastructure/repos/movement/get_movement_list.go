package movementrepo

import (
	"fmt"
	"time"

	"github.com/jiradeto/corn-kernels-backend/app/constants"
	"github.com/jiradeto/corn-kernels-backend/app/entities"
	"github.com/jiradeto/corn-kernels-backend/app/infrastructure/models"
	"github.com/jiradeto/corn-kernels-backend/app/utils/gerrors"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// GetMovementListInput is an input for find one movement
type GetMovementListInput struct {
	Name            *string
	ProductID       *string
	Type            *string
	URL             *string
	FromCreatedDate *time.Time
	ToCreatedDate   *time.Time
	Limit           *int
}

func (repo *repo) GetMovementList(tx *gorm.DB, input GetMovementListInput) ([]*entities.StockMovement, error) {
	const errLocation = "movementRepo/GetMovementList: %s"
	var resultModel models.StockMovements
	query := repo.selectDB(tx)

	if input.FromCreatedDate != nil {
		query = query.Where("created_at >= ?", *input.FromCreatedDate)
	}

	if input.ProductID != nil {
		query = query.Where("product_id = ?", *input.ProductID)
	}

	if input.Type != nil {
		query = query.Where("movement_type = ?", *input.Type)
	}

	if input.ToCreatedDate != nil {
		query = query.Where("created_at <= ?", *input.ToCreatedDate)
	}

	if input.Name != nil {
		query = query.Where("name LIKE ?", "%"+*input.Name+"%")
	}

	if input.URL != nil {
		query = query.Where("url LIKE ?", "%"+*input.URL+"%")
	}

	if input.Limit != nil {
		query = query.Limit(*input.Limit)
	}

	result := query.Order("created_at desc").Find(&resultModel)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, gerrors.RecordNotFoundError{
				Code:    constants.StatusCodeEntryNotFound,
				Message: constants.ErrorMessageNotFound,
			}
		}
		return nil, errors.Wrap(result.Error, fmt.Sprintf(errLocation, "unable to find movements due to database error"))
	}

	resultEntities, err := resultModel.ToEntities()
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(errLocation, "unable to covert from model to entities"))
	}

	return resultEntities, nil
}
