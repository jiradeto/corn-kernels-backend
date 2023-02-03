package healthcheckhttp

import (
	"github.com/gin-gonic/gin"
	"github.com/jiradeto/corn-kernels-backend/app/utils/response"
)

// CheckServiceVersion ..
func (handler *httpHandler) CheckServiceVersion(c *gin.Context) {
	checkServiceVersion, err := handler.Usecase.CheckServiceVersion(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseSuccess(c, checkServiceVersion)
}
