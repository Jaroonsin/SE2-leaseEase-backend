package handlers

import (
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

func (h *propertyHandler) ListAllProperties(c *fiber.Ctx) error {
	properties, err := h.propertyService.ListAllProperties()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(properties)
}

func (h *propertyHandler) FindPropertyByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid property ID"})
	}

	property, err := h.propertyService.FindPropertyByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(property)
}

func (h *propertyHandler) ListPropertiesWithPagination(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))    // Default page = 1
	limit, _ := strconv.Atoi(c.Query("limit", "10")) // Default limit = 10

	response, err := h.propertyService.ListPropertiesWithPagination(page, limit)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(response)
}
