package productrepo

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

// GetProductListInput is an input for find one product
type GetProductListInput struct {
	Query           *string
	Name            *string
	URL             *string
	FromCreatedDate *time.Time
	ToCreatedDate   *time.Time
	Limit           *int
	PreloadStock    bool
}

func (repo *repo) GetProductList(tx *gorm.DB, input GetProductListInput) ([]*entities.Product, error) {
	const errLocation = "productRepo/GetProductList: %s"
	var resultModel models.Products
	query := repo.selectDB(tx)

	if input.FromCreatedDate != nil {
		query = query.Where("created_at >= ?", *input.FromCreatedDate)
	}

	if input.ToCreatedDate != nil {
		query = query.Where("created_at <= ?", *input.ToCreatedDate)
	}

	if input.Query != nil {
		query = query.Where("name LIKE ?", "%"+*input.Query+"%")
		query = query.Where("description LIKE ?", "%"+*input.Query+"%")
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

	if input.PreloadStock {
		query = query.Preload("ProductStock")
	}

	result := query.Order("created_at asc").Find(&resultModel)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, gerrors.RecordNotFoundError{
				Code:    constants.StatusCodeEntryNotFound,
				Message: constants.ErrorMessageNotFound,
			}
		}
		return nil, errors.Wrap(result.Error, fmt.Sprintf(errLocation, "unable to find products due to database error"))
	}

	resultEntities, err := resultModel.ToEntities()
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(errLocation, "unable to covert from model to entities"))
	}

	return resultEntities, nil
}
