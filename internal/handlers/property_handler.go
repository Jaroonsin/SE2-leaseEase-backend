package handlers

import (
	"LeaseEase/internal/dtos"
	"LeaseEase/internal/services"
	"LeaseEase/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

// propertyHandler handles property endpoints.
type propertyHandler struct {
	propertyService services.PropertyService
}

// NewPropertyHandler creates a new property handler.
func NewPropertyHandler(propertyService services.PropertyService) *propertyHandler {
	return &propertyHandler{
		propertyService: propertyService,
	}
}

// CreateProperty godoc
// @Summary Create a property
// @Description Create a new property
// @Tags Property
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param request body dtos.CreateDTO true "Property Data"
// @Success 201 {object} map[string]interface{} "Property created successfully"
// @Failure 400 {object} map[string]interface{} "Bad Request"
// @Failure 500 {object} map[string]interface{} "Internal Server Error"
// @Router /properties [post]
func (h *propertyHandler) CreateProperty(c *fiber.Ctx) error {
	var req dtos.CreateDTO
	if err := c.BodyParser(&req); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Failed to parse request body")
	}
	req.LessorID = uint(c.Locals("user").(jwt.MapClaims)["user_id"].(float64))

	err := h.propertyService.CreateProperty(&req)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}
	
	return utils.SuccessResponse(c, fiber.StatusCreated, "Property created successfully", nil)
}

// UpdateProperty godoc
// @Summary Update a property
// @Description Update existing property data
// @Tags Property
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Property ID"
// @Param request body dtos.UpdateDTO true "Updated property data"
// @Success 200 {object} map[string]interface{} "Property updated successfully"
// @Failure 400 {object} map[string]interface{} "Bad Request"
// @Failure 500 {object} map[string]interface{} "Internal Server Error"
// @Router /properties/{id} [put]
func (h *propertyHandler) UpdateProperty(c *fiber.Ctx) error {
	PropertyID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid property ID")
	}

	var req dtos.UpdateDTO
	if err := c.BodyParser(&req); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Failed to parse request body")
	}
	req.PropertyID = uint(PropertyID)
	err = h.propertyService.UpdateProperty(&req)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return utils.SuccessResponse(c, fiber.StatusOK, "Property updated successfully", nil)
}

// DeleteProperty godoc
// @Summary Delete a property
// @Description Delete a property by ID
// @Tags Property
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Property ID"
// @Success 200 {object} map[string]interface{} "Property deleted successfully"
// @Failure 400 {object} map[string]interface{} "Bad Request"
// @Failure 500 {object} map[string]interface{} "Internal Server Error"
// @Router /properties/{id} [delete]
func (h *propertyHandler) DeleteProperty(c *fiber.Ctx) error {
	var req dtos.DeleteDTO
	propertyID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid property ID")
	}

	req.PropertyID = uint(propertyID)
	err = h.propertyService.DeleteProperty(&req)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}
	return utils.SuccessResponse(c, fiber.StatusOK, "Property deleted successfully", nil)
}

// GetAllProperty godoc
// @Summary Get all properties
// @Description Retrieve list of all properties with pagination
// @Tags Property
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param page query int false "Page number" default(1)
// @Param pageSize query int false "Page size" default(10)
// @Success 200 {object} map[string]interface{} "Properties retrieved successfully"
// @Failure 500 {object} map[string]interface{} "Internal Server Error"
// @Router /properties [get]
func (h *propertyHandler) GetAllProperty(c *fiber.Ctx) error {
	pageStr := c.Query("page", "")
	pageSizeStr := c.Query("pageSize", "")

	// Case 1: No pagination parameters → Fetch all properties
	if pageStr == "" && pageSizeStr == "" {
		properties, err := h.propertyService.GetAllProperty(0, 0) // 0 means fetch all
		if err != nil {
			return utils.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
		}
		return utils.SuccessResponse(c, fiber.StatusOK, "Properties retrieved successfully", properties)
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
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}
	return utils.SuccessResponse(c, fiber.StatusOK, "Properties retrieved successfully", properties)
}

// GetPropertyByID godoc
// @Summary Get property by ID
// @Description Retrieve property details by its ID
// @Tags Property
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Property ID"
// @Success 200 {object} map[string]interface{} "Property retrieved successfully"
// @Failure 400 {object} map[string]interface{} "Bad Request"
// @Failure 404 {object} map[string]interface{} "Not Found"
// @Router /properties/{id} [get]
func (h *propertyHandler) GetPropertyByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid property ID")
	}

	property, err := h.propertyService.GetPropertyByID(uint(id))
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusNotFound, err.Error())
	}

	return utils.SuccessResponse(c, fiber.StatusOK, "Property retrieved successfully", property)
}
