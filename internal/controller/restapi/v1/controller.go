package v1

import (
	usecase "tenant-gate/internal/usecase"
	"tenant-gate/pkg/jwt"
	"tenant-gate/pkg/logger"

	"github.com/go-playground/validator/v10"
)

type Controller struct {
	um     usecase.Manager
	jwt    *jwt.Manager
	logger logger.Interface

	validator *validator.Validate
}
