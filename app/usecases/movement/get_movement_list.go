package movementusecase

import (
	"context"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/jiradeto/corn-kernels-backend/app/constants"
	"github.com/jiradeto/corn-kernels-backend/app/entities"
	movementrepo "github.com/jiradeto/corn-kernels-backend/app/infrastructure/repos/movement"
	"github.com/jiradeto/corn-kernels-backend/app/utils/gerrors"
	"github.com/pkg/errors"
)

// Validate is a function to validate function input
func (c *GetMovementListInput) Validate() error {
	const errLocation = "movementUsecase/GetMovementListInput/Validate: %s"
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

// GetMovementListInput is an input for GetMovementList
type GetMovementListInput struct {
	Name            *string `validate:"omitempty,max=128"`
	ProductID       *string `validate:"omitempty,max=128"`
	Type            *string `validate:"omitempty,max=128"`
	URL             *string `validate:"omitempty,max=128"`
	Limit           *int    `validate:"omitempty,min=0"`
	FromCreatedDate *time.Time
	ToCreatedDate   *time.Time
}

func (uc *useCase) GetMovementList(ctx context.Context, input GetMovementListInput) ([]*entities.StockMovement, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}
	products, err := uc.StockMovementRepo.GetMovementList(nil, movementrepo.GetMovementListInput{
		ProductID: input.ProductID,
		Type:      input.Type,
	})
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}

	return products, nil
}
