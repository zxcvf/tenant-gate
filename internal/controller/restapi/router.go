package restapi

import (
	"net/http"
	"tenant-gate/config"
	"tenant-gate/internal/controller/restapi/middleware"
	v1 "tenant-gate/internal/controller/restapi/v1"
	"tenant-gate/internal/usecase"
	"tenant-gate/pkg/jwt"
	"tenant-gate/pkg/logger"

	filber "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func NewRouter(app *filber.App, cfg *config.Config, usecases *usecase.Manager, jwtManager *jwt.Manager, l logger.Interface) {
	app.Use(cors.New())
	app.Use(middleware.Logger(l))
	app.Use(middleware.Recovery(l))
	// prometheus metrics
	//	app.Use(middleware.Prometheus(cfg, l))

	// swagger docs
	//	app.Get("/swagger/*", swagger.HandlerDefault) // default

	// health check
	app.Get("/health", func(c *filber.Ctx) error {
		return c.SendString("OK")
	})
	// k8s health check
	app.Get("/health", func(c *filber.Ctx) error {
		return c.SendStatus(http.StatusOK)
	})

	appV1Group := app.Group("/v1")
	{
		// tracing
		v1.NewRoutes(appV1Group, usecases, jwtManager, l)
	}
}
