package handlers

import (
	"LeaseEase/internal/dtos"
	"LeaseEase/internal/services"
	"time"

	"github.com/gofiber/fiber/v2"
)

type userHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) *userHandler {
	return &userHandler{
		userService: userService,
	}
}

func (h *userHandler) Register(c *fiber.Ctx) error {

	var req dtos.RegisterDTO
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Failed to parse request body"})
	}

	err := h.userService.Register(&req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "User registered successfully"})
}

func (h *userHandler) Login(c *fiber.Ctx) error {

	var req dtos.LoginDTO
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Failed to parse request body"})
	}

	token,err := h.userService.Login(&req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
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
	

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "User login successfully"})
}
