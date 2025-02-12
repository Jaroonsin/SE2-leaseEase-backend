package handlers

import (
	"LeaseEase/internal/dtos"
	"LeaseEase/internal/services"
	"LeaseEase/utils"
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
// @Tags auth
// @Accept json
// @Produce json
// @Param register body dtos.RegisterDTO true "Register request payload"
// @Success 201  "User registered successfully"
// @Failure 400  "Invalid request payload"
// @Failure 500  "Internal server error"
// @Router /auth/register [post]
func (h *authHandler) Register(c *fiber.Ctx) error {

	var req dtos.RegisterDTO
	if err := c.BodyParser(&req); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Failed to parse request body")
	}

	err := h.authService.Register(&req)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return utils.SuccessResponse(c, fiber.StatusCreated, "User registered successfully", nil)
}

// Login godoc
// @Summary Login an existing user
// @Description Authenticate user and set an authentication cookie.
// @Tags auth
// @Accept json
// @Produce json
// @Param login body dtos.LoginDTO true "Login request payload"
// @Success 201 {array} utils.Response "User login successfully"
// @Failure 400  "Invalid request payload"
// @Failure 500  "Internal server error"
// @Router /auth/login [post]
func (h *authHandler) Login(c *fiber.Ctx) error {

	var req dtos.LoginDTO
	if err := c.BodyParser(&req); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Failed to parse request body")
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
	
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User login successfully",
		"token":   token,
	})
}
