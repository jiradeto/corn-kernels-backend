package producthttp

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jiradeto/corn-kernels-backend/app/constants"
	productusecase "github.com/jiradeto/corn-kernels-backend/app/usecases/product"
	"github.com/jiradeto/corn-kernels-backend/app/utils/gerrors"
	"github.com/jiradeto/corn-kernels-backend/app/utils/response"
	"github.com/pkg/errors"
)

type GetProductListRequest struct {
	Query           *string    `form:"q"`
	Name            *string    `form:"name"`
	URL             *string    `form:"url"`
	Limit           *int       `form:"limit,default=20"`
	FromCreatedDate *time.Time `form:"from"`
	ToCreatedDate   *time.Time `form:"to"`
}

func (handler *httpHandler) GetProductList(c *gin.Context) {
	errLocation := "movementHTTP/GetProductList: %s"
	var req GetProductListRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		response.ResponseError(c, gerrors.ParameterError{
			Code:    constants.StatusCodeInvalidParameters,
			Message: constants.ErrorMessageParameterInvalid,
		}.Wrap(errors.Wrap(err, fmt.Sprintf(errLocation, "unable to process parameter(s)"))))
		return
	}

	products, err := handler.ProductUsecase.GetProductList(c.Request.Context(), productusecase.GetProductListInput{
		Query:           req.Query,
		Name:            req.Name,
		URL:             req.URL,
		FromCreatedDate: req.FromCreatedDate,
		ToCreatedDate:   req.ToCreatedDate,
		Limit:           req.Limit,
	})
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	response.ResponseSuccess(c, products)
}
