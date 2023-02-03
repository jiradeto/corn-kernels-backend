package producthttp

import (
	"github.com/gin-gonic/gin"
	productusecase "github.com/jiradeto/corn-kernels-backend/app/usecases/product"
)

// HTTPHandler ...
type HTTPHandler interface {
	CreateOneProduct(c *gin.Context)
	GetProductList(c *gin.Context)
	GetDashboardData(c *gin.Context)
}

type httpHandler struct {
	ProductUsecase productusecase.UseCase
}

// New is a constructor method of HTTPHandler
func New(
	productUsecase productusecase.UseCase,
) HTTPHandler {
	return &httpHandler{
		ProductUsecase: productUsecase,
	}
}
