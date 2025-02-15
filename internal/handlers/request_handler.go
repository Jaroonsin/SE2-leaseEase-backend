package handlers

import (
	"LeaseEase/internal/dtos"
	"LeaseEase/internal/services"
	"LeaseEase/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

type requestHandler struct {
	requestService services.RequestService
}

func NewRequestHandler(requestService services.RequestService) *requestHandler {
	return &requestHandler{
		requestService: requestService,
	}
}

// CreateRequest godoc
// @Summary      Create a New Lease Request
// @Description  Parses the request body and creates a new lease request using the lessee ID from the JWT token.
// @Tags         Request
// @Accept       json
// @Produce      json
// @Param        request  body      dtos.CreateRequest  true  "Lease Request Data"
// @Success      201      {object}  utils.Response  "Request created successfully"
// @Failure      400      {object}  utils.Response    "Failed to parse request body"
// @Failure      500      {object}  utils.Response    "Internal server error"
// @Router       /request/create [post]
func (h *requestHandler) CreateRequest(c *fiber.Ctx) error {
	var req dtos.CreateRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Failed to parse request body")
	}
	
	lesseeID := uint(c.Locals("user").(jwt.MapClaims)["user_id"].(float64))

	err := h.requestService.CreateRequest(&req, lesseeID)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}
	return utils.SuccessResponse(c, fiber.StatusCreated, "Request created successfully", nil)
}

// UpdateRequest godoc
// @Summary      Update an Existing Lease Request
// @Description  Parses the request body and updates an existing lease request identified by its ID.
// @Tags         Request
// @Accept       json
// @Produce      json
// @Param        id       path      int                 true  "Request ID"
// @Param        request  body      dtos.UpdateRequest  true  "Lease Request Update Data"
// @Success      200      {object}  utils.Response  "Request updated successfully"
// @Failure      400      {object}  utils.Response    "Failed to parse request body or invalid request ID"
// @Failure      404      {object}  utils.Response    "Request not found"
// @Failure      500      {object}  utils.Response    "Internal server error"
// @Router       /requests/update/{id} [put]
func (h *requestHandler) UpdateRequest(c *fiber.Ctx) error {
	var req dtos.UpdateRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Failed to parse request body")
	}
	requestID , err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid request ID")
	}

	err = h.requestService.UpdateRequest(&req, uint(requestID))
	if err != nil {
		if err == gorm.ErrRecordNotFound{
			return utils.ErrorResponse(c, fiber.StatusNotFound, "Request not found")
		}
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}
	return utils.SuccessResponse(c, fiber.StatusOK, "Request updated successfully", nil)
}


// DeleteRequest godoc
// @Summary      Delete a Lease Request
// @Description  Deletes a lease request using the request ID provided in the URL.
// @Tags         Request
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Request ID"
// @Success      200  {object}  utils.Response  "Request deleted successfully"
// @Failure      400  {object}  utils.Response    "Invalid request ID"
// @Failure      404  {object}  utils.Response    "Request not found"
// @Failure      500  {object}  utils.Response    "Internal server error"
// @Router       /requests/delete/{id} [delete]
func (h *requestHandler) DeleteRequest(c *fiber.Ctx) error {
	requestID , err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid request ID")
	}

	err = h.requestService.DeleteRequest(uint(requestID))
	if err != nil {
		if err == gorm.ErrRecordNotFound{
			return utils.ErrorResponse(c, fiber.StatusNotFound, "Request not found")
		}
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}
	
	return utils.SuccessResponse(c, fiber.StatusOK, "Request deleted successfully", nil)
}