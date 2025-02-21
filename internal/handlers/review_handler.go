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

func (h *reviewHandler) GetAllReviews(c *fiber.Ctx) error {
	pageStr := c.Query("page", "")
	pageSizeStr := c.Query("pageSize", "")

	propertyID, err := strconv.Atoi(c.Params("propertyID"))
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid property ID")
	}

	if pageStr == "" && pageSizeStr == "" {
		reviews, err := h.reviewService.GetAllReviews(uint(propertyID), 0, 0)
		if err != nil {
			return utils.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
		}
		return utils.SuccessResponse(c, fiber.StatusOK, "Success", reviews)
	}

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	pageSize := 10
	if pageSizeStr != "" {
		pageSize, err = strconv.Atoi(pageSizeStr)
		if err != nil || pageSize < 1 {
			pageSize = 10
		}
	}

	reviews, err := h.reviewService.GetAllReviews(uint(propertyID), page, pageSize)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}
	return utils.SuccessResponse(c, fiber.StatusOK, "Success", reviews)
}
