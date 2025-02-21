package handlers

import (
	"LeaseEase/internal/dtos"
	"LeaseEase/internal/services"
	"LeaseEase/utils"
	"LeaseEase/utils/constant"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

// reviewHandler handles review endpoints.
type reviewHandler struct {
	reviewService services.ReviewService
}

// NewReviewHandler creates a new review handler.
func NewReviewHandler(reviewService services.ReviewService) *reviewHandler {
	return &reviewHandler{
		reviewService: reviewService,
	}
}

func (h *reviewHandler) CreateReview(c *fiber.Ctx) error {
	var req dtos.CreateReviewDTO

	if err := c.BodyParser(&req); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, constant.ErrParsebody)
	}

	LesseeID := uint(c.Locals("user").(jwt.MapClaims)["user_id"].(float64))

	err := h.reviewService.CreateReview(&req, LesseeID)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return utils.SuccessResponse(c, fiber.StatusCreated, constant.SuccesCreateReview, nil)
}

func (h *reviewHandler) UpdateReview(c *fiber.Ctx) error {
	reviewID, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "invalid review ID")
	}

	var req dtos.UpdateReviewDTO
	if err := c.BodyParser(&req); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, constant.ErrParsebody)
	}

	LesseeID := uint(c.Locals("user").(jwt.MapClaims)["user_id"].(float64))

	err = h.reviewService.UpdateReview(uint(reviewID), &req, LesseeID)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return utils.SuccessResponse(c, fiber.StatusOK, "Review updated successfully", nil)
}

func (h *reviewHandler) DeleteReview(c *fiber.Ctx) error {
	reviewID, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "invalid review ID")
	}

	lesseeID := uint(c.Locals("user").(jwt.MapClaims)["user_id"].(float64))

	err = h.reviewService.DeleteReview(uint(reviewID), lesseeID)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return utils.SuccessResponse(c, fiber.StatusOK, "Review deleted successfully", nil)
}
