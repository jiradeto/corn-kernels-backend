package healthcheckhttp

import (
	"github.com/gin-gonic/gin"
	"github.com/jiradeto/corn-kernels-backend/app/utils/response"
)

func (handler *httpHandler) CheckLiveness(c *gin.Context) {
	response.ResponseSuccess(c, nil)
}
