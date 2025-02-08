package handlers

import (
	"LeaseEase/internal/dtos"
	"LeaseEase/internal/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type propertyHandler struct {
	propertyService services.PropertyService
}

func NewPropertyHandler(propertyService services.PropertyService) *propertyHandler {
	return &propertyHandler{
		propertyService: propertyService,
	}
}

func (h *propertyHandler) CreateProperty(c *fiber.Ctx) error {
	var req dtos.CreateDTO
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Failed to parse request body"})
	}

	err := h.propertyService.CreateProperty(&req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Property registered successfully"})
}

func (h *propertyHandler) UpdateProperty(c *fiber.Ctx) error {
	PropertyID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	var req dtos.UpdateDTO
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Failed to parse request body"})
	}
	req.PropertyID = uint(PropertyID)
	err = h.propertyService.UpdateProperty(&req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Property updated successfully"})
}

func (h *propertyHandler) DeleteProperty(c *fiber.Ctx) error {
	var req dtos.DeleteDTO
	propertyID, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	req.PropertyID = uint(propertyID)
	err = h.propertyService.DeleteProperty(&req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Property deleted successfully"})
}

func (h *propertyHandler) GetAllProperty(c *fiber.Ctx) error {
	page, err := strconv.Atoi(c.Query("page", "1")) // Default to page 1
	if err != nil || page < 1 {
		page = 1
	}

	pageSize, err := strconv.Atoi(c.Query("pageSize", "10")) // Default to 10 items per page
	if err != nil || pageSize < 1 {
		pageSize = 10
	}

	properties, err := h.propertyService.GetAllProperty(page, pageSize)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(properties)
}
