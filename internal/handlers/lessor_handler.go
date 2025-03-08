package handlers

import (
	"LeaseEase/internal/dtos"
	"LeaseEase/internal/services"
	"LeaseEase/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type lessorHandler struct {
	lessorService services.LessorService
}

func NewLessorHandler(lessorService services.LessorService) *lessorHandler {
	return &lessorHandler{
		lessorService: lessorService,
	}
}

// AcceptReservation godoc
// @Summary Accept a reservation
// @Description Accept a reservation by ID
// @Tags Lessor
// @Accept json
// @Produce json
// @Param id path int true "Reservation ID"
// @Param reservation body dtos.ApprovalReservationDTO true "Reservation details"
// @Success 200 {object} utils.Response "Reservation accepted successfully"
// @Failure 400 {object} utils.Response "Invalid reservation ID"
// @Failure 400 {object} utils.Response "Invalid request body"
// @Failure 500 {object} utils.Response "Failed to accept reservation"
// @Router /lessor/accept/{id} [post]
func (h *lessorHandler) AcceptReservation(c *fiber.Ctx) error {
	reservationID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid reservation ID")
	}
	lesseeID := uint(c.Locals("user").(jwt.MapClaims)["user_id"].(float64))
	var req dtos.ApprovalReservationDTO
	if err := c.BodyParser(&req); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid request body")
	}

	if err := h.lessorService.AcceptReservation(uint(reservationID), &req, lesseeID); err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to accept reservation")
	}

	return utils.SuccessResponse(c, fiber.StatusOK, "Reservation accepted successfully", nil)
}

// DeclineReservation godoc
// @Summary Decline a reservation
// @Description Decline a reservation by ID
// @Tags Lessor
// @Accept json
// @Produce json
// @Param id path int true "Reservation ID"
// @Param reservation body dtos.ApprovalReservationDTO true "Reservation details"
// @Success 200 {object} utils.Response "Reservation declined successfully"
// @Failure 400 {object} utils.Response "Invalid reservation ID"
// @Failure 400 {object} utils.Response "Invalid request body"
// @Failure 500 {object} utils.Response "Failed to decline reservation"
// @Router /lessor/decline/{id} [post]
func (h *lessorHandler) DeclineReservation(c *fiber.Ctx) error {
	reservationID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid reservation ID")
	}
	lesseeID := uint(c.Locals("user").(jwt.MapClaims)["user_id"].(float64))
	var req dtos.ApprovalReservationDTO
	if err := c.BodyParser(&req); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid request body")
	}

	if err := h.lessorService.DeclineReservation(uint(reservationID), &req, lesseeID); err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to decline reservation")
	}

	return utils.SuccessResponse(c, fiber.StatusOK, "Reservation declined successfully", nil)
}

// GetReservationsByPropID godoc
// @Summary Get reservations by property ID
// @Description Get reservations by property ID with pagination
// @Tags Lessor
// @Accept json
// @Produce json
// @Param propID path int true "Property ID"
// @Param page query int false "Page number"
// @Param pageSize query int false "Page size"
// @Success 200 {object} utils.Response "Reservations retrieved successfully"
// @Failure 400 {object} utils.Response "Invalid property ID"
// @Failure 400 {object} utils.Response "Invalid page or pageSize parameter"
// @Failure 500 {object} utils.Response "Failed to retrieve reservations"
// @Router /lessor/reservations/{propID} [get]
func (h *lessorHandler) GetReservationsByPropID(c *fiber.Ctx) error {
	propID, err := strconv.Atoi(c.Params("propID"))
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid property ID")
	}

	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil || page < 1 {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid page parameter")
	}

	pageSize, err := strconv.Atoi(c.Query("pageSize", "10"))
	if err != nil || pageSize < 1 {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid pageSize parameter")
	}

	offset := (page - 1) * pageSize

	reservations, err := h.lessorService.GetReservationsByPropertyID(uint(propID), pageSize, offset)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to retrieve reservations")
	}
	if reservations == nil {
		return utils.ErrorResponse(c, fiber.StatusNotFound, "No reservations found")
	}

	return utils.SuccessResponse(c, fiber.StatusOK, "Reservations retrieved successfully", reservations)
}
