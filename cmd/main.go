package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jiradeto/corn-kernels-backend/app/constants"
	"github.com/jiradeto/corn-kernels-backend/app/environments"
	"github.com/jiradeto/corn-kernels-backend/app/infrastructure/interfaces/connectors"
	"github.com/jiradeto/corn-kernels-backend/app/infrastructure/migrations"
	movementrepo "github.com/jiradeto/corn-kernels-backend/app/infrastructure/repos/movement"
	productrepo "github.com/jiradeto/corn-kernels-backend/app/infrastructure/repos/product"
	stockrepo "github.com/jiradeto/corn-kernels-backend/app/infrastructure/repos/stock"
	healthcheckhttp "github.com/jiradeto/corn-kernels-backend/app/presenters/health_check"
	movementhttp "github.com/jiradeto/corn-kernels-backend/app/presenters/movement"
	producthttp "github.com/jiradeto/corn-kernels-backend/app/presenters/product"
	"github.com/jiradeto/corn-kernels-backend/app/routes"
	healthcheckusecase "github.com/jiradeto/corn-kernels-backend/app/usecases/health_check"
	movementusecase "github.com/jiradeto/corn-kernels-backend/app/usecases/movement"
	productusecase "github.com/jiradeto/corn-kernels-backend/app/usecases/product"
	"github.com/jiradeto/corn-kernels-backend/app/utils/cors"
	"github.com/jiradeto/corn-kernels-backend/app/utils/gerrors"
	"github.com/jiradeto/corn-kernels-backend/app/utils/loggers"
	"github.com/jiradeto/corn-kernels-backend/app/utils/response"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "app",
		Short: "Start application",
		Run:   startApp,
	}
	rootCmd.Flags().Bool("production", false, "whether an app running in production")
	rootCmd.Flags().Bool("check-migration", false, "check and run migratioh if necessary")
	rootCmd.AddCommand(&cobra.Command{
		Use:   "migrate",
		Short: "Migrate database schema",
		Run:   runMigration,
	})
	rootCmd.AddCommand(&cobra.Command{
		Use:   "rollbackmigration",
		Short: "Rollback Last migration",
		Run:   rollbackMigration,
	})

	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
	}
}

func startApp(cmd *cobra.Command, wow []string) {
	useProductionEnv, err := cmd.Flags().GetBool("production")
	if err != nil {
		log.Fatal("Error parsing 'production' flag")
	}
	environments.Init(useProductionEnv)
	if useProductionEnv {
		gin.SetMode(gin.ReleaseMode)
	}
	runMigration(nil, nil)

	middlewareLog := loggers.New()
	loggers.JSON.Info("Starting corn-kernels api...")
	middlewareCORS := cors.New()

	db := connectors.ConnectSqliteDB()

	// repos
	productRepo := productrepo.New(db)
	movementRepo := movementrepo.New(db)
	stockRepo := stockrepo.New(db)

	// usecases
	healthcheckUseCase := healthcheckusecase.New()
	productUsecase := productusecase.New(productRepo, stockRepo)
	movementUsecase := movementusecase.New(movementRepo, productRepo, stockRepo)

	// http
	healthcheckHTTP := healthcheckhttp.New(healthcheckUseCase)
	productHTTP := producthttp.New(productUsecase)
	movementHTTP := movementhttp.New(movementUsecase)

	app := gin.Default()
	app.Use(middlewareLog)
	app.Use(middlewareCORS)
	app.NoRoute(func(c *gin.Context) {
		response.ResponseError(c, gerrors.RecordNotFoundError{
			Code:    constants.StatusCodeEntryNotFound,
			Message: "endpoint not found",
		}.Wrap(errors.New("the requested endpoint is not registered")))
	})

	routes.RegisterHealthCheckRoutes(app, &routes.HTTPRoutes{
		HealthCheck: healthcheckHTTP,
	})

	routes.RegisterAPIRoutes(app, &routes.HTTPRoutes{
		ProductHTTP:  productHTTP,
		MovementHTTP: movementHTTP,
	})

	var port = os.Getenv("PORT")
	if port == "" {
		port = "8003"
	}

	server := &http.Server{
		Addr:    ":" + port,
		Handler: app,
	}

	var exit = make(chan os.Signal)
	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		// service connections
		loggers.JSON.Info(fmt.Sprintf("Listening and serving HTTP on %s\n", port))
		if err := server.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			loggers.JSON.Error(fmt.Sprintf("Listening and serving HTTP Error: %s\n", err))
		}
	}()

	<-exit
	loggers.JSON.Info("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		loggers.JSON.Error(fmt.Sprint("Server forced to shutdown:", err))
	}
	loggers.JSON.Info("Server exiting")
}

func runMigration(_ *cobra.Command, _ []string) {
	environments.Init(false)
	loggers.JSON.Info("Running migration...")
	err := migrations.Migrate()
	if err != nil {
		loggers.JSON.Error(err.Error())
		panic(err)
	}
}

// rollback last migration
func rollbackMigration(_ *cobra.Command, _ []string) {
	environments.Init(false)
	loggers.JSON.Info("Rollback last migration...")

	err := migrations.RollbackLast()
	if err != nil {
		loggers.JSON.Error(err.Error())
		panic(err)
	}
}
