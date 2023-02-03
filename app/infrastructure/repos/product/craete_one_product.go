package productrepo

import (
	"fmt"

	"github.com/jiradeto/corn-kernels-backend/app/entities"
	"github.com/jiradeto/corn-kernels-backend/app/infrastructure/models"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// CreateOneProductInput is a DTO for creating one Product
type CreateOneProductInput struct {
	ProductEntity *entities.Product
}

// CreateOneProduct is a function for creating one Product from data model
func (repo *repo) CreateOneProduct(tx *gorm.DB, input CreateOneProductInput) (*entities.Product, error) {
	const errLocation = "[productRepo/CreateOneProduct] %s"

	productModel, err := new(models.Product).FromEntity(input.ProductEntity)
	if err != nil {
		return nil, errors.Wrapf(err, errLocation, "unable to parse an entity to model")
	}

	query := repo.selectDB(tx)
	result := query.Create(productModel)
	if result.Error != nil {
		return nil, errors.Wrap(result.Error, fmt.Sprintf(errLocation, "unable to create Product due to database error"))
	}

	resultEntity, err := productModel.ToEntity()
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(errLocation, "unable to covert from model to entity"))
	}

	return resultEntity, nil
}
