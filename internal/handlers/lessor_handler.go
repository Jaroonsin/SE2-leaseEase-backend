package handlers

import (
	"LeaseEase/internal/services"
	"LeaseEase/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
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
// @Success 200 {object} utils.Response "Reservation accepted successfully"
// @Failure 400 {object} utils.Response "Invalid reservation ID"
// @Failure 500 {object} utils.Response "Failed to accept reservation"
// @Router /lessor/accept/{id} [post]
func (h *lessorHandler) AcceptReservation(c *fiber.Ctx) error {
	reservationID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid reservation ID")
	}

	err = h.lessorService.AcceptReservation(uint(reservationID))
	if err != nil {
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
// @Success 200 {object} utils.Response "Reservation declined successfully"
// @Failure 400 {object} utils.Response "Invalid reservation ID"
// @Failure 500 {object} utils.Response "Failed to decline reservation"
// @Router /lessor/decline/{id} [post]
func (h *lessorHandler) DeclineReservation(c *fiber.Ctx) error {
	reservationID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid reservation ID")
	}

	err = h.lessorService.DeclineReservation(uint(reservationID))
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to decline reservation")
	}

	return utils.SuccessResponse(c, fiber.StatusOK, "Reservation declined successfully", nil)
}
