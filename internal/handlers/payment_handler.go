package handlers

import (
	"LeaseEase/internal/dtos"
	"LeaseEase/internal/services"
	"LeaseEase/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type paymentHandler struct {
	paymentService services.PaymentService
}

func NewPaymentHandler(paymentService services.PaymentService) *paymentHandler {
	return &paymentHandler{paymentService: paymentService}
}

// HandlePayment godoc
// @Summary Process a payment
// @Description This endpoint processes a payment using the provided user ID, amount, currency, and card token.
// @Tags payments
// @Accept json
// @Produce json
// @Param payment body dtos.PaymentDTO true "Payment details"
// @Success 200 {object} utils.Response "Payment successful"
// @Failure 400 {object} utils.Response "Invalid request body"
// @Failure 500 {object} utils.Response "Payment process failed"
// @Router /payments/process [post]
func (h *paymentHandler) HandlePayment(c *fiber.Ctx) error {
	var req dtos.PaymentDTO

	if err := c.BodyParser(&req); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid request body")
	}

	lesseeID := uint(c.Locals("user").(jwt.MapClaims)["user_id"].(float64))
	err := h.paymentService.ProcessPayment(lesseeID, req.Currency, req.Token, req.ReservationID)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Payment process failed")
	}

	return utils.SuccessResponse(c, fiber.StatusOK, "Payment successful", nil)
}
