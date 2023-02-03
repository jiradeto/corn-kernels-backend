package movementusecase

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/jiradeto/corn-kernels-backend/app/constants"
	"github.com/jiradeto/corn-kernels-backend/app/entities"
	movementrepo "github.com/jiradeto/corn-kernels-backend/app/infrastructure/repos/movement"
	productrepo "github.com/jiradeto/corn-kernels-backend/app/infrastructure/repos/product"
	stockrepo "github.com/jiradeto/corn-kernels-backend/app/infrastructure/repos/stock"
	"github.com/jiradeto/corn-kernels-backend/app/utils/gerrors"
	"github.com/pkg/errors"
)

// Validate is a function to validate function input
func (c *CreateOneStockMovementInput) Validate() error {
	const errLocation = "movementUsecase/CreateOneStockMovementInput/Validate: %s"
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

// CreateOneStockMovementInput is an input for CreateOneProduct
type CreateOneStockMovementInput struct {
	ProductID    *string    `json:"productID" validate:"required,min=1"`
	Description  *string    `json:"description" validate:"required,min=1"`
	MovementType *string    `json:"movementType" validate:"required,min=1"`
	Quantity     int        `json:"quantity" validate:"required"`
	CreatedAt    *time.Time `json:"createdAt" `
}

func (uc *useCase) CreateOneStockMovement(ctx context.Context, input CreateOneStockMovementInput) (*entities.StockMovement, error) {
	const errLocation = "movementUsecase/CreateOneStockMovement/Validate: %s"
	if err := input.Validate(); err != nil {
		return nil, err
	}

	product, err := uc.ProductRepo.FindOneProduct(nil, productrepo.FindOneProductInput{
		ID: input.ProductID,
	})
	if err != nil {
		return nil, gerrors.InternalError{
			Code:    constants.StatusCodeDatabaseError,
			Message: constants.ErrorMessageDatabaseError,
		}.Wrap(errors.Wrapf(err, errLocation, "unable to find movement"))
	}

	if product == nil {
		return nil, gerrors.RecordNotFoundError{
			Code:    constants.StatusCodeEntryNotFound,
			Message: constants.ErrorMessageNotFound,
		}.Wrap(errors.Errorf(errLocation, "not found movement"))
	}

	if strings.ToUpper(*input.MovementType) != "IN" && strings.ToUpper(*input.MovementType) != "OUT" {
		return nil, gerrors.ParameterError{
			Code:    constants.StatusCodeInvalidParameters,
			Message: fmt.Sprintf(constants.ErrorMessageFmtInvalidFormat, "parameter(s)"),
		}.Wrap(errors.Wrap(err, fmt.Sprintf(errLocation, "unable to process parameter(s)")))
	}

	stockMovementEntity := &entities.StockMovement{
		ProductID:    product.ID,
		Description:  input.Description,
		MovementType: input.MovementType,
		Quantity:     input.Quantity,
	}

	if input.CreatedAt != nil {
		stockMovementEntity.CreatedAt = input.CreatedAt
	}

	stockMovement, err := uc.StockMovementRepo.CreateOneStockMovement(nil, movementrepo.CreateOneStockMovementInput{
		StockMovementEntity: stockMovementEntity,
	})
	if err != nil {
		return nil, err
	}

	// update product stock
	productStock, err := uc.StockRepo.FindOneProductStock(nil, stockrepo.FindOneProductStockInput{ID: product.ID})
	if err != nil {
		return nil, err
	}

	if strings.ToUpper(*input.MovementType) == "IN" {
		productStock.Quantity += input.Quantity
	} else if strings.ToUpper(*input.MovementType) == "OUT" {
		productStock.Quantity -= input.Quantity
	}
	err = uc.StockRepo.UpdateOneProductStock(nil, stockrepo.UpdateOneProductStockInput{
		ProductStockEntity: productStock,
	})

	if err != nil {
		return nil, err
	}

	return stockMovement, nil
}
