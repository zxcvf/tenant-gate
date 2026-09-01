package v1

import (
	"github.com/gofiber/fiber/v2"
)

func errorResponse(ctx *fiber.Ctx, status int, message string) error {
	return ctx.Status(status).JSON(fiber.Map{
		"error": message,
	})
}

var (
	errInternalServerErrorCode    = fiber.StatusInternalServerError
	errInternalServerErrorMessage = "Internal server error"

	errInvalidRequestCode    = fiber.StatusBadRequest
	errInvalidRequestMessage = "Invalid request body"

	errUnauthorizedCode    = fiber.StatusUnauthorized
	errUnauthorizedMessage = "invalid credentials"
)
