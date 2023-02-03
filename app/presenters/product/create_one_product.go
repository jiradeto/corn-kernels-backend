package producthttp

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jiradeto/corn-kernels-backend/app/constants"
	productusecase "github.com/jiradeto/corn-kernels-backend/app/usecases/product"
	"github.com/jiradeto/corn-kernels-backend/app/utils/gerrors"
	"github.com/jiradeto/corn-kernels-backend/app/utils/response"
	"github.com/pkg/errors"
)

// CreateOneMovementRequest request body for CreateOneProduct
type CreateOneMovementRequest struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
}

func (handler *httpHandler) CreateOneProduct(c *gin.Context) {
	errLocation := "movementHTTP/CreateOneProduct: %s"

	var req CreateOneMovementRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ResponseError(c, gerrors.ParameterError{
			Code:    constants.StatusCodeInvalidParameters,
			Message: fmt.Sprintf(constants.ErrorMessageFmtInvalidFormat, "parameter(s)"),
		}.Wrap(errors.Wrap(err, fmt.Sprintf(errLocation, "unable to process parameter(s)"))))
		return
	}

	movement, err := handler.ProductUsecase.CreateOneProduct(c.Request.Context(), productusecase.CreateOneProductInput{
		Name:        req.Name,
		Description: req.Description,
	})
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	response.ResponseSuccess(c, movement)
}
