package v1

import (
	"tenant-gate/internal/controller/middleware"
	"tenant-gate/internal/usecase"
	"tenant-gate/pkg/jwt"
	"tenant-gate/pkg/logger"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func NewRoutes(apiV1Group fiber.Router, um *usecase.Manager, jwtManager *jwt.Manager, l logger.Interface) {
	controller := &Controller{
		um:        *um,
		jwt:       jwtManager,
		validator: validator.New(validator.WithRequiredStructEnabled()),
		logger:    l,
	}

	apiV1Group.Get("/health", func(c *fiber.Ctx) error {
		return c.SendStatus(200)
	})

	// Public routes
	authGroup := apiV1Group.Group("/auth")
	{
		authGroup.Post("/login", controller.login)
	}

	// protected routes
	protected := apiV1Group.Group("", middleware.Auth(jwtManager))

	userGroup := protected.Group("/user")
	{
		userGroup.Get("/profile", controller.profile)
	}

}
