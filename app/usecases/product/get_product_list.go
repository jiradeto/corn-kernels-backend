package productusecase

import (
	"context"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/jiradeto/corn-kernels-backend/app/constants"
	"github.com/jiradeto/corn-kernels-backend/app/entities"
	productrepo "github.com/jiradeto/corn-kernels-backend/app/infrastructure/repos/product"
	"github.com/jiradeto/corn-kernels-backend/app/utils/gerrors"
	"github.com/pkg/errors"
)

// Validate is a function to validate function input
func (c *GetProductListInput) Validate() error {
	const errLocation = "productUsecase/GetProductListInput/Validate: %s"
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

// GetProductListInput is an input for GetProductList
type GetProductListInput struct {
	Query           *string `validate:"omitempty,max=128"`
	Name            *string `validate:"omitempty,max=128"`
	URL             *string `validate:"omitempty,max=128"`
	Limit           *int    `validate:"omitempty,min=0"`
	FromCreatedDate *time.Time
	ToCreatedDate   *time.Time
}

func (uc *useCase) GetProductList(ctx context.Context, input GetProductListInput) ([]*entities.Product, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}
	products, err := uc.ProductRepo.GetProductList(nil, productrepo.GetProductListInput{
		PreloadStock: true,
		Query:        input.Query,
	})
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}

	return products, nil
}
