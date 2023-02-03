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

type GetMovementListRequest struct {
	Name            *string    `form:"name"`
	ProductID       *string    `form:"productID"`
	Type            *string    `form:"type"`
	URL             *string    `form:"url"`
	Limit           *int       `form:"limit,default=20"`
	FromCreatedDate *time.Time `form:"from"`
	ToCreatedDate   *time.Time `form:"to"`
}

func (handler *httpHandler) GetMovementList(c *gin.Context) {
	errLocation := "movementHTTP/GetMovementList: %s"
	var req GetMovementListRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		response.ResponseError(c, gerrors.ParameterError{
			Code:    constants.StatusCodeInvalidParameters,
			Message: constants.ErrorMessageParameterInvalid,
		}.Wrap(errors.Wrap(err, fmt.Sprintf(errLocation, "unable to process parameter(s)"))))
		return
	}

	movements, err := handler.MovementUsecase.GetMovementList(c.Request.Context(), movementusecase.GetMovementListInput{
		Name:            req.Name,
		Type:            req.Type,
		ProductID:       req.ProductID,
		FromCreatedDate: req.FromCreatedDate,
		ToCreatedDate:   req.ToCreatedDate,
		Limit:           req.Limit,
	})
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	response.ResponseSuccess(c, movements)
}
