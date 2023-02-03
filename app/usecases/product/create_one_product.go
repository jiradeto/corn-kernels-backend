package productusecase

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/jiradeto/corn-kernels-backend/app/constants"
	"github.com/jiradeto/corn-kernels-backend/app/entities"
	productrepo "github.com/jiradeto/corn-kernels-backend/app/infrastructure/repos/product"
	stockrepo "github.com/jiradeto/corn-kernels-backend/app/infrastructure/repos/stock"
	"github.com/jiradeto/corn-kernels-backend/app/utils/gerrors"
	"github.com/pkg/errors"
)

// Validate is a function to validate function input
func (c *CreateOneProductInput) Validate() error {
	const errLocation = "productUsecase/CreateOneProductInput/Validate: %s"
	validate := validator.New()
	err := validate.Struct(c)
	if err != nil {
		ve, ok := err.(validator.ValidationErrors)
		if !ok {
			return gerrors.InternalError{
				Code:    constants.StatusCodeInvalidParameters,
				Message: constants.ErrorMessageUnableProcessParameter,
			}.Wrap(errors.Wrapf(err, errLocation, "failed to convert validation error"))
		}
		return gerrors.ParameterError{
			Code:            constants.StatusCodeInvalidParameters,
			ValidatorErrors: &ve,
		}.Wrap(errors.Wrapf(err, errLocation, "unable to process the request due to some parameter(s) are invalid"))
	}
	return nil
}

// CreateOneProductInput is an input for CreateOneProduct
type CreateOneProductInput struct {
	Name        *string `json:"name" validate:"required,min=1"`
	Description *string `json:"description" validate:"required,min=1"`
}

func (uc *useCase) CreateOneProduct(ctx context.Context, input CreateOneProductInput) (*entities.Product, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}
	product, err := uc.ProductRepo.CreateOneProduct(nil, productrepo.CreateOneProductInput{
		ProductEntity: &entities.Product{
			Name:        input.Name,
			Description: input.Description,
		},
	})
	if err != nil {
		return nil, err
	}

	//  initialize stock to new product
	_, err = uc.StockRepo.CreateOneProductStock(nil, stockrepo.CreateOneProductStockInput{
		ProductStockEntity: &entities.ProductStock{
			ProductID: product.ID,
			Quantity:  0,
		},
	})

	if err != nil {
		return nil, err
	}

	return product, nil
}
