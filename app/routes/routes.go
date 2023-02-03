package routes

import (
	"github.com/gin-gonic/gin"
	healthcheckhttp "github.com/jiradeto/corn-kernels-backend/app/presenters/health_check"
	movementhttp "github.com/jiradeto/corn-kernels-backend/app/presenters/movement"
	producthttp "github.com/jiradeto/corn-kernels-backend/app/presenters/product"
)

// HTTPRoutes ...
type HTTPRoutes struct {
	HealthCheck  healthcheckhttp.HTTPHandler
	ProductHTTP  producthttp.HTTPHandler
	MovementHTTP movementhttp.HTTPHandler
}

// RegisterHealthCheckRoutes represents routing for healthcheck group
func RegisterHealthCheckRoutes(r *gin.Engine, httpRoutes *HTTPRoutes) {
	apiRoute := r.Group("/health")
	{
		apiRoute.GET("/check", httpRoutes.HealthCheck.CheckLiveness)
		apiRoute.GET("/version", httpRoutes.HealthCheck.CheckServiceVersion)
	}
}

// RegisterAPIRoutes represents routing for api group
func RegisterAPIRoutes(r *gin.Engine, httpRoutes *HTTPRoutes) {
	apiRoute := r.Group("/api")
	v1GroupRoute := apiRoute.Group("v1") // , middlewares.JWTAuth.MiddlewareFunc())
	{
		// stock
		{
			v1GroupRoute.POST("/stock", httpRoutes.MovementHTTP.CreateOneMovement)
			v1GroupRoute.GET("/stock/list", httpRoutes.MovementHTTP.GetMovementList)
		}
		// product
		{
			v1GroupRoute.GET("/dashboard", httpRoutes.ProductHTTP.GetDashboardData)
			v1GroupRoute.POST("/product", httpRoutes.ProductHTTP.CreateOneProduct)
			v1GroupRoute.GET("/product/list", httpRoutes.ProductHTTP.GetProductList)
		}
	}
}
