package handlers

import (
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
		Secure:   false, // Requires HTTPS ? true for Prod
		SameSite: "Strict",
		Path:     "/",
		Expires:  time.Now().Add(time.Hour * 3),
	})

	return utils.SuccessResponse(c, fiber.StatusCreated, "User login successfully", nil)
}

// AuthCheck godoc
// @Summary Check authentication status
// @Description Validates JWT token and returns authentication status.
// @ID 3
// @Tags auth
// @Produce json
// @Success 200 {object} utils.Response "User is authenticated"
// @Failure 401 {object} utils.Response "Unauthorized - Invalid token"
// @Router /auth/check [get]
func (h *authHandler) AuthCheck(c *fiber.Ctx) error {
	// Retrieve the JWT token from cookies
	token := c.Cookies("auth_token")
	if token == "" {
		return utils.ErrorResponse(c, fiber.StatusUnauthorized, "No token provided")
	}

	// Validate token using the AuthCheck service
	claims, err := h.authService.AuthCheck(token)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusUnauthorized, err.Error())
	}

	// Return success response with user details
	return utils.SuccessResponse(c, fiber.StatusOK, "User is authenticated", claims)
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

	c.Cookie(&fiber.Cookie{
		Name:     "auth_token",
		Value:    "",
		HTTPOnly: true,
		Secure:   false, // Requires HTTPS ? true for Prod
		SameSite: "Strict",
		Path:     "/",
		Expires:  time.Now(),
	})

	return utils.SuccessResponse(c, fiber.StatusCreated, "User logout successfully", nil)
}
