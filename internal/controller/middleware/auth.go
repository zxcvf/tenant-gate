package middleware

import (
	"net/http"
	"strings"

	"tenant-gate/pkg/jwt"

	"github.com/gofiber/fiber/v2"
)

const _bearerParts = 2
const PayloadKey = "jwt"

type errorResponse struct {
	Error string `json:"error"`
}

// Auth returns a JWT authentication middleware for Fiber.
func Auth(jwtManager *jwt.Manager) func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		header := ctx.Get("Authorization")
		if header == "" {
			return ctx.Status(http.StatusUnauthorized).JSON(errorResponse{Error: "missing authorization header"})
		}

		parts := strings.SplitN(header, " ", _bearerParts)
		if len(parts) != _bearerParts || parts[0] != "Bearer" {
			return ctx.Status(http.StatusUnauthorized).JSON(errorResponse{Error: "invalid authorization header format"})
		}

		payload, err := jwtManager.ParseToken(parts[1])
		if err != nil {
			return ctx.Status(http.StatusUnauthorized).JSON(errorResponse{Error: "invalid or expired token"})
		}

		ctx.Locals(PayloadKey, payload)
		return ctx.Next()
	}
}
