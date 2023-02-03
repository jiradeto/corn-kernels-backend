package movementhttp

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jiradeto/corn-kernels-backend/app/constants"
	movementusecase "github.com/jiradeto/corn-kernels-backend/app/usecases/movement"
	"github.com/jiradeto/corn-kernels-backend/app/utils/gerrors"
	"github.com/jiradeto/corn-kernels-backend/app/utils/response"
	"github.com/pkg/errors"
)

// CreateOneMovementRequest request body for CreateOneMovement
type CreateOneMovementRequest struct {
	ProductID    *string    `json:"productID"`
	Description  *string    `json:"description"`
	MovementType *string    `json:"movementType"`
	Quantity     *int       `json:"quantity"`
	CreatedAt    *time.Time `json:"createdAt"`
}

func (handler *httpHandler) CreateOneMovement(c *gin.Context) {
	errLocation := "movementHTTP/CreateOneMovement: %s"

	var req CreateOneMovementRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ResponseError(c, gerrors.ParameterError{
			Code:    constants.StatusCodeInvalidParameters,
			Message: fmt.Sprintf(constants.ErrorMessageFmtInvalidFormat, "parameter(s)"),
		}.Wrap(errors.Wrap(err, fmt.Sprintf(errLocation, "unable to process parameter(s)"))))
		return
	}

	movement, err := handler.MovementUsecase.CreateOneStockMovement(c.Request.Context(), movementusecase.CreateOneStockMovementInput{
		ProductID:    req.ProductID,
		Description:  req.Description,
		MovementType: req.MovementType,
		Quantity:     *req.Quantity,
		CreatedAt:    req.CreatedAt,
	})
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	response.ResponseSuccess(c, movement)
}
