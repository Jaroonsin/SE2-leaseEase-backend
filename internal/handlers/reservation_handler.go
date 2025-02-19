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

type reservationHandler struct {
	reservationService services.ReservationService
}

func NewReservationHandler(reservationService services.ReservationService) *reservationHandler {
	return &reservationHandler{
		reservationService: reservationService,
	}
}

// CreateReservation godoc
// @Summary      Create a New Lease Reservation
// @Description  Parses the reservation body and creates a new lease reservation using the lessee ID from the JWT token.
// @Tags         Reservation
// @Accept       json
// @Produce      json
// @Param        reservation  body      dtos.CreateReservation  true  "Lease Reservation Data"
// @Success      201      {object}  utils.Response  "Reservation created successfully"
// @Failure      400      {object}  utils.Response    "Failed to parse reservation body"
// @Failure      500      {object}  utils.Response    "Internal server error"
// @Router       /reservation/create [post]
func (h *reservationHandler) CreateReservation(c *fiber.Ctx) error {
	var req dtos.CreateReservation
	if err := c.BodyParser(&req); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Failed to parse reservation body")
	}

	lesseeID := uint(c.Locals("user").(jwt.MapClaims)["user_id"].(float64))

	err := h.reservationService.CreateReservation(&req, lesseeID)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}
	return utils.SuccessResponse(c, fiber.StatusCreated, "Reservation created successfully", nil)
}

// UpdateReservation godoc
// @Summary      Update an Existing Lease Reservation
// @Description  Parses the reservation body and updates an existing lease reservation identified by its ID.
// @Tags         Reservation
// @Accept       json
// @Produce      json
// @Param        id       path      int                 true  "Reservation ID"
// @Param        reservation  body      dtos.UpdateReservation  true  "Lease Reservation Update Data"
// @Success      200      {object}  utils.Response  "Reservation updated successfully"
// @Failure      400      {object}  utils.Response    "Failed to parse reservation body or invalid reservation ID"
// @Failure      404      {object}  utils.Response    "Reservation not found"
// @Failure      500      {object}  utils.Response    "Internal server error"
// @Router       /reservations/update/{id} [put]
func (h *reservationHandler) UpdateReservation(c *fiber.Ctx) error {
	var req dtos.UpdateReservation
	if err := c.BodyParser(&req); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Failed to parse reservation body")
	}
	reservationID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid reservation ID")
	}

	err = h.reservationService.UpdateReservation(&req, uint(reservationID))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return utils.ErrorResponse(c, fiber.StatusNotFound, "Reservation not found")
		}
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}
	return utils.SuccessResponse(c, fiber.StatusOK, "Reservation updated successfully", nil)
}

// DeleteReservation godoc
// @Summary      Delete a Lease Reservation
// @Description  Deletes a lease reservation using the reservation ID provided in the URL.
// @Tags         Reservation
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Reservation ID"
// @Success      200  {object}  utils.Response  "Reservation deleted successfully"
// @Failure      400  {object}  utils.Response    "Invalid reservation ID"
// @Failure      404  {object}  utils.Response    "Reservation not found"
// @Failure      500  {object}  utils.Response    "Internal server error"
// @Router       /reservations/delete/{id} [delete]
func (h *reservationHandler) DeleteReservation(c *fiber.Ctx) error {
	reservationID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid reservation ID")
	}

	err = h.reservationService.DeleteReservation(uint(reservationID))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return utils.ErrorResponse(c, fiber.StatusNotFound, "Reservation not found")
		}
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return utils.SuccessResponse(c, fiber.StatusOK, "Reservation deleted successfully", nil)
}

// ApproveReservation godoc
// @Summary      Approve a Lease Reservation
// @Description  Approves a lease reservation using the reservation ID provided in the URL.
// @Tags         Reservation
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Reservation ID"
// @Param        reservation  body      dtos.ApproveReservation  true  "Reservation Status"
// @Success      200  {object}  utils.Response  "Reservation approved successfully"
// @Failure      400  {object}  utils.Response    "Invalid reservation ID"
// @Failure      404  {object}  utils.Response    "Reservation not found"
// @Failure      500  {object}  utils.Response    "Internal server error"
// @Router       /reservations/{id} [put]
func (h *reservationHandler) ApproveReservation(c *fiber.Ctx) error {
	var req dtos.ApproveReservation
	reservationID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid reservation ID")
	}

	if err := c.BodyParser(&req); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Failed to parse reservation body")
	}

	err = h.reservationService.ApproveReservation(req.Status, uint(reservationID))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return utils.ErrorResponse(c, fiber.StatusNotFound, "Reservation not found")
		}
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return utils.SuccessResponse(c, fiber.StatusOK, "Reservation approved successfully", nil)
}
