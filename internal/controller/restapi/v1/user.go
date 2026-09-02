package v1

import (
	"net/http"
	"tenant-gate/internal/controller/middleware"
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

	token, err := r.um.User.Login(ctx.UserContext(), body.TenantName, body.Email, body.Password)
	if err != nil {
		r.logger.Error("Login failed: %v", err)
		return errorResponse(ctx, errUnauthorizedCode, errUnauthorizedMessage)
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Login successful",
		"token":   token,
	})

}

// @Summary     Get User Profile
// @Description Get the profile of the authenticated user
// @ID          profile
// @Tags        user
// @Produce     json
// @Success     200     {object} response.UserProfile
// @Failure     401     {object} response.Error
// @Failure     500     {object} response.Error
// @Router      /user/profile [get]
func (r *Controller) profile(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Profile fetched successful",
		"user":    ctx.Locals(middleware.PayloadKey),
	})
}
