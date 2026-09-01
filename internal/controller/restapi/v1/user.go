package v1

import (
	"net/http"
	"tenant-gate/internal/controller/restapi/v1/request"

	"github.com/gofiber/fiber/v2"
)

// @Summary     Login
// @Description Authenticate user and get JWT token
// @ID          login
// @Tags        auth
// @Accept      json
// @Produce     json
// @Param       request body     request.Login true "Login credentials"
// @Success     200     {object} response.Token
// @Failure     400     {object} response.Error
// @Failure     401     {object} response.Error
// @Failure     500     {object} response.Error
// @Router      /auth/login [post]
func (r *Controller) login(ctx *fiber.Ctx) error {
	var body request.Login
	if err := ctx.BodyParser(&body); err != nil {
		r.logger.Error("Failed to parse request body: %v", err)
		return errorResponse(ctx, errInvalidRequestCode, errInvalidRequestMessage)
	}

	if err := r.validator.Struct(body); err != nil {
		r.logger.Error("Validation failed: %v", err)
		return errorResponse(ctx, errInvalidRequestCode, errInvalidRequestMessage)
	}

	r.um.User.Login(ctx.UserContext(), body.TenantName, body.Email, body.Password)

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Login successful",
		"token":   "your-jwt-token", // Replace with actual token generation logic
	})

}
