package handlers

import (
	"LeaseEase/internal/dtos"
	"LeaseEase/internal/services"
	"LeaseEase/utils"
	"LeaseEase/utils/constant"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
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
// @Security cookieAuth
// @Param request body dtos.PropertyDTO true "Property Data"
// @Success 201 {array} utils.Response "Property created successfully"
// @Failure 400 {array} utils.Response "Bad Request"
// @Failure 500 {array} utils.Response"Internal Server Error"
// @Router /properties/create [post]
func (h *propertyHandler) CreateProperty(c *fiber.Ctx) error {
	var req dtos.PropertyDTO
	if err := c.BodyParser(&req); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, constant.ErrParsebody)
	}
	LessorID := uint(c.Locals("user").(jwt.MapClaims)["user_id"].(float64))

	err := h.propertyService.CreateProperty(&req, LessorID)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return utils.SuccessResponse(c, fiber.StatusCreated, constant.SuccessCreateProp, nil)
}

// UpdateProperty godoc
// @Summary Update a property
// @Description Update existing property data
// @Tags Property
// @Accept json
// @Produce json
// @Security cookieAuth
// @Param id path int true "Property ID"
// @Param request body dtos.PropertyDTO true "Updated property data"
// @Success 200 {array} utils.Response "Property updated successfully"
// @Failure 400 {array} utils.Response "Bad Request"
// @Failure 404 {array} utils.Response "Property not found"
// @Failure 500 {array} utils.Response "Internal Server Error"
// @Router /properties/update/{id} [put]
func (h *propertyHandler) UpdateProperty(c *fiber.Ctx) error {
	PropertyID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid property ID")
	}

	var req dtos.PropertyDTO
	if err := c.BodyParser(&req); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Failed to parse request body")
	}

	err = h.propertyService.UpdateProperty(&req, uint(PropertyID))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return utils.ErrorResponse(c, fiber.StatusNotFound, "Property not found")
		}
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
// @Security cookieAuth
// @Param id path uint true "Property ID"
// @Success 200 {array} utils.Response "Property deleted successfully"
// @Failure 400 {array} utils.Response "Bad Request"
// @Failure 404 {array} utils.Response "Property not found"
// @Failure 500 {array} utils.Response "Internal Server Error"
// @Router /properties/delete/{id} [delete]
func (h *propertyHandler) DeleteProperty(c *fiber.Ctx) error {
	propertyID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid property ID")
	}

	err = h.propertyService.DeleteProperty(uint(propertyID))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return utils.ErrorResponse(c, fiber.StatusNotFound, "Property not found")
		}
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
// @Security cookieAuth
// @Param page query int false "Page number" default(1)
// @Param pageSize query int false "Page size" default(10)
// @Success 200 {array} utils.Response{data=[]dtos.GetPropertyDTO} "Properties retrieved successfully"
// @Failure 500 {array} utils.Response "Internal Server Error"
// @Router /properties/get [get]
func (h *propertyHandler) GetAllProperty(c *fiber.Ctx) error {
	pageStr := c.Query("page", "")
	pageSizeStr := c.Query("pageSize", "")

	LessorID := uint(c.Locals("user").(jwt.MapClaims)["user_id"].(float64))

	// Case 1: No pagination parameters → Fetch all properties
	if pageStr == "" && pageSizeStr == "" {
		properties, err := h.propertyService.GetAllProperty(LessorID, 0, 0) // 0 means fetch all
		if err != nil {
			return utils.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
		}
		return utils.SuccessResponse(c, fiber.StatusOK, constant.SuccessGetAllProp, properties)
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

	properties, err := h.propertyService.GetAllProperty(LessorID, page, pageSize)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}
	return utils.SuccessResponse(c, fiber.StatusOK, constant.SuccessGetAllProp, properties)
}

// GetPropertyByID godoc
// @Summary Get property by ID
// @Description Retrieve property details by its ID
// @Tags Property
// @Accept json
// @Produce json
// @Security cookieAuth
// @Param id path int true "Property ID"
// @Success 200 {array} utils.Response{data=dtos.GetPropertyDTO} "Property retrieved successfully"
// @Failure 400 {array} utils.Response "Bad Request"
// @Failure 404 {array} utils.Response "Not Found"
// @Router /properties/get/{id} [get]
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

// SearchProperty godoc
// @Summary Search properties
// @Description Search properties using query parameters
// @Tags Property
// @Accept json
// @Produce json
// @Security cookieAuth
// @Param page query int true "Page number"
// @Param pagesize query int true "Page size"
// @Param name query string false "Property name keyword"
// @Param minprice query number false "Minimum price"
// @Param maxprice query number false "Maximum price"
// @Param minsize query number false "Minimum size"
// @Param maxsize query number false "Maximum size"
// @Param sortby query string false "Order field (price or size)"
// @Param order query string false "Order direction (asc or desc)"
// @Param availability query bool false "Availability status"
// @Success 200 {array} utils.Response{data=[]dtos.GetPropertyDTO} "Properties retrieved successfully"
// @Failure 400 {array} utils.Response "Bad Request"
// @Failure 500 {array} utils.Response "Internal Server Error"
// @Router /properties/search [get]
func (h *propertyHandler) SearchProperty(c *fiber.Ctx) error {

	query := c.Queries()

	if query["page"] == "" && query["pagesize"] == "" {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Page and pageSize parameters are required")
	}

	if query == nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Query parameter is required")
	}

	properties, err := h.propertyService.SearchProperty(query)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return utils.SuccessResponse(c, fiber.StatusOK, "Properties retrieved successfully", properties)
}

// Auto Complete:
//
//	@Summary Auto complete property search
//	@Description Retrieve property suggestions based on a partial search query
//	@Tags Property
//	@Accept json
//	@Produce json
//	@Security cookieAuth
//	@Param query query string true "Partial property name"
//	@Success 200 {object} utils.Response "Properties retrieved successfully"
//	@Failure 400 {object} utils.Response "Bad Request"
//	@Failure 500 {object} utils.Response "Internal Server Error"
//	@Router /properties/autocomplete [get]
func (h *propertyHandler) AutoComplete(c *fiber.Ctx) error {
	query := c.Query("query")
	if query == "" {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Query parameter is required")
	}

	properties, err := h.propertyService.AutoComplete(query)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return utils.SuccessResponse(c, fiber.StatusOK, "Properties retrieved successfully", properties)
}
