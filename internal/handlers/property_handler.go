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
	pageStr := c.Query("page", "")
	pageSizeStr := c.Query("pageSize", "")

	// Case 1: No pagination parameters → Fetch all properties
	if pageStr == "" && pageSizeStr == "" {
		properties, err := h.propertyService.GetAllProperty(0, 0) // 0 means fetch all
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		return c.Status(fiber.StatusOK).JSON(properties)
	}

	// Parse page number (default to 1)
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	// Parse page size (default to 10)
	pageSize := 10
	if pageSizeStr != "" {
		pageSize, err = strconv.Atoi(pageSizeStr)
		if err != nil || pageSize < 1 {
			pageSize = 10
		}
	}

	properties, err := h.propertyService.GetAllProperty(page, pageSize)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(properties)
}

func (h *propertyHandler) GetPropertyByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid property ID"})
	}

	property, err := h.propertyService.GetPropertyByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(property)
}
