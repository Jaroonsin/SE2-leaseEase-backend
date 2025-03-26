package handlers

import (
	"LeaseEase/config"
	"LeaseEase/internal/dtos"
	"LeaseEase/internal/services"
	"LeaseEase/utils"
	"LeaseEase/utils/constant"
	"time"

	"github.com/gofiber/fiber/v2"
)

type authHandler struct {
	authService services.AuthService
}

func NewAuthHandler(authService services.AuthService) *authHandler {
	return &authHandler{
		authService: authService,
	}
}

// Register godoc
// @Summary Register a new user
// @Description Register a new user account with the provided details.
// @ID 1
// @Tags auth
// @Accept json
// @Produce json
// @Param register body dtos.RegisterDTO true "Register request payload"
// @Success 201 {object} utils.Response "User registered successfully"  example({"staus_code": 201, "message": "User registered successfully", "data": nil})
// @Failure 400 {array} utils.Response "Invalid request payload"
// @Failure 500 {array} utils.Response "Internal server error"
// @Router /auth/register [post]
func (h *authHandler) Register(c *fiber.Ctx) error {

	var req dtos.RegisterDTO
	if err := c.BodyParser(&req); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, constant.ErrParsebody)
	}

	err := h.authService.Register(&req)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return utils.SuccessResponse(c, fiber.StatusCreated, constant.SuccessRegister, nil)
}

// Login godoc
// @Summary Login an existing user
// @Description Authenticate user and set an authentication cookie.
// @ID 2
// @Tags auth
// @Accept json
// @Produce json
// @Param login body dtos.LoginDTO true "Login request payload"
// @Success 201 {array} utils.Response "User login successfully"
// @Failure 400 {array} utils.Response "Invalid request payload"
// @Failure 500 {array} utils.Response "Internal server error"
// @Router /auth/login [post]
func (h *authHandler) Login(c *fiber.Ctx) error {
	// Secure := config.LoadEnv() == "production" || config.LoadEnv() == "staging"

	var req dtos.LoginDTO
	if err := c.BodyParser(&req); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, constant.ErrParsebody)
	}

	token, err := h.authService.Login(&req)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	c.Cookie(&fiber.Cookie{
		Name:     "auth_token",
		Value:    token,
		HTTPOnly: true,
		// Secure:   Secure, // Requires HTTPS ? true for Prod
		SameSite: fiber.CookieSameSiteStrictMode,
		//Domain:   config.LoadConfig().ClientURL,
		Path:    "/",
		Expires: time.Now().Add(time.Hour * 3),
	})

	return utils.SuccessResponse(c, fiber.StatusCreated, "User login successfully", fiber.Map{"token": token})
}

// Login godoc
// @Summary Logout an existing user
// @Description Clear the authentication cookie to logout the user.
// @Tags auth
// @Accept json
// @Produce json
// @Success 201 {array} utils.Response "User logout successfully"
// @Router /auth/logout [post]
func (h *authHandler) Logout(c *fiber.Ctx) error {
	Secure := config.LoadEnv() == "production" || config.LoadEnv() == "staging"
	c.Cookie(&fiber.Cookie{
		Name:     "auth_token",
		Value:    "delete",
		HTTPOnly: true,
		Secure:   Secure, // Requires HTTPS ? true for Prod
		SameSite: fiber.CookieSameSiteNoneMode,
		Path:     "/",
		//Domain:   config.LoadConfig().ClientURL,
		Expires: time.Now().Add(time.Second * -3),
	})

	return utils.SuccessResponse(c, fiber.StatusCreated, "User logout successfully", nil)
}

// API: Request OTP for registration
// RequestOTP godoc
// @Summary Request OTP for authentication
// @Description Sends a one-time password (OTP) to the user's contact information provided.
// @Tags auth
// @Accept  json
// @Produce  json
// @Param  request body dtos.RequestOTPDTO true "Request payload containing user identifier"
// @Success 201 {object} utils.Response{data=any} "OTP sent successfully"
// @Failure 400 {object} utils.Response "Bad Request - Unable to parse request body"
// @Failure 500 {object} utils.Response "Internal Server Error - Failed to process OTP request"
// @Router /auth/request-otp [post]
func (h *authHandler) RequestOTP(c *fiber.Ctx) error {
	var req dtos.RequestOTPDTO
	if err := c.BodyParser(&req); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, constant.ErrParsebody)
	}

	if err := h.authService.RequestOTP(&req); err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return utils.SuccessResponse(c, fiber.StatusCreated, "OTP sent successfully", nil)
}

// VerifyOTP godoc
// @Summary Verify provided OTP
// @Description Validates the OTP provided by the user for authentication.
// @Tags auth
// @Accept  json
// @Produce  json
// @Param  request body dtos.VerifyOTPDTO true "Request payload containing OTP and user identifier"
// @Success 200 {object} utils.Response{data=any} "OTP verification successful"
// @Failure 400 {object} utils.Response "Bad Request - Invalid OTP payload or incorrect OTP"
// @Failure 500 {object} utils.Response "Internal Server Error - Failed to verify OTP"
// @Router /auth/verify-otp [post]
func (h *authHandler) VerifyOTP(c *fiber.Ctx) error {

	var req dtos.VerifyOTPDTO
	if err := c.BodyParser(&req); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, constant.ErrParsebody)
	}

	if err := h.authService.VerifyOTP(&req); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	return utils.SuccessResponse(c, fiber.StatusOK, "User registration verified", nil)
}

// ResetPasswordRequest godoc
// @Summary Forgot a password
// @Description Generates and sends a password reset link to the provided email address if the user exists.
// @Tags auth
// @Accept  json
// @Produce  json
// @Param  request body dtos.ResetPassRequestDTO true "Request payload containing the user's email"
// @Success 200 {object} utils.Response "Reset link sent successfully"
// @Failure 400 {object} utils.Response "Bad Request - Invalid request payload"
// @Failure 404 {object} utils.Response "Not Found - Email not associated with any account"
// @Failure 500 {object} utils.Response "Internal Server Error - Failed to send reset email"
// @Router /auth/forgot-password [post]
func (h *authHandler) ResetPasswordRequest(c *fiber.Ctx) error {
	var req dtos.ResetPassRequestDTO
	if err := c.BodyParser(&req); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, constant.ErrParsebody)
	}

	resetLink, err := h.authService.RequestPasswordReset(&req)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to create reset link")
	}

	if err := utils.SendPasswordResetEmail(&req, resetLink); err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to send reset email")
	}

	if config.LoadEnv() == "development" {
		return utils.SuccessResponse(c, fiber.StatusOK, "Reset link sent", fiber.Map{"reset_link": resetLink})
	}
	return utils.SuccessResponse(c, fiber.StatusOK, "Reset link sent", nil)
}

// ResetPassword godoc
// @Summary Reset user password
// @Description Resets the user's password using the provided reset token and new password.
// @Tags auth
// @Accept  json
// @Produce  json
// @Param  request body dtos.ResetPassDTO true "Request payload containing the reset token and new password"
// @Success 200 {object} utils.Response "Password reset successful"
// @Failure 400 {object} utils.Response "Bad Request - Invalid request payload"
// @Failure 401 {object} utils.Response "Unauthorized - Invalid or expired reset token"
// @Failure 500 {object} utils.Response "Internal Server Error - Unable to reset password"
// @Router /auth/reset-password [post]
func (h *authHandler) ResetPassword(c *fiber.Ctx) error {
	var req dtos.ResetPassDTO
	if err := c.BodyParser(&req); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, constant.ErrParsebody)
	}

	err := h.authService.ResetPassword(&req)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusUnauthorized, "Failed to reset password")
	}

	return utils.SuccessResponse(c, fiber.StatusOK, "Password reset successful", nil)
}
