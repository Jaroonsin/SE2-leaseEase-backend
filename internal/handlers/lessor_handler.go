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
