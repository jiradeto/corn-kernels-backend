package movementhttp

import (
	"github.com/gin-gonic/gin"
	movementusecase "github.com/jiradeto/corn-kernels-backend/app/usecases/movement"
)

// HTTPHandler ...
type HTTPHandler interface {
	CreateOneMovement(c *gin.Context)
	GetMovementList(c *gin.Context)
}

type httpHandler struct {
	MovementUsecase movementusecase.UseCase
}

// New is a constructor method of HTTPHandler
func New(
	movementUsecase movementusecase.UseCase,
) HTTPHandler {
	return &httpHandler{
		MovementUsecase: movementUsecase,
	}
}
