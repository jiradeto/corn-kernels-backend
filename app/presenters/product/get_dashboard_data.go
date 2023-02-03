package producthttp

import (
	"github.com/gin-gonic/gin"
	"github.com/jiradeto/corn-kernels-backend/app/utils/response"
)

func (handler *httpHandler) GetDashboardData(c *gin.Context) {

	dashboardData, err := handler.ProductUsecase.GetDashboardData(c.Request.Context())
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	response.ResponseSuccess(c, dashboardData)
}
