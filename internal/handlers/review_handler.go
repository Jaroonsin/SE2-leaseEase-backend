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

// CreateReview godoc
// @Summary      Create a new review
// @Description  Create a new review for a property by the authenticated lessee.
// @Tags         Review
// @Accept       json
// @Produce      json
// @Param        review  body      dtos.CreateReviewDTO  true  "Review Data"
// @Success      201     {object}  map[string]string            "Review created successfully"
// @Failure      400     {object}  map[string]string            "Invalid request body"
// @Failure      500     {object}  map[string]string            "Internal server error"
// @Router       /propertyReview/create/ [post]
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

// UpdateReview godoc
// @Summary      Update an existing review
// @Description  Update a review by its ID for the authenticated lessee.
// @Tags         Review
// @Accept       json
// @Produce      json
// @Param        id      path      uint                  true  "Review ID"
// @Param        review  body      dtos.UpdateReviewDTO  true  "Updated Review Data"
// @Success      200     {object}  map[string]string            "Review updated successfully"
// @Failure      400     {object}  map[string]string            "Invalid review ID or body"
// @Failure      500     {object}  map[string]string            "Internal server error"
// @Router       /propertyReview/update/{id} [put]
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

// DeleteReview godoc
// @Summary      Delete a review
// @Description  Delete a review by its ID for the authenticated lessee.
// @Tags         Review
// @Produce      json
// @Param        id   path      uint      true  "Review ID"
// @Success      200  {object}  map[string]string"Review deleted successfully"
// @Failure      400  {object}  map[string]string"Invalid review ID"
// @Failure      500  {object}  map[string]string"Internal server error"
// @Router       /propertyReview/delete/{id} [delete]
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

// GetAllReviews godoc
// @Summary      Retrieve all reviews for a property
// @Description  Get all reviews for a specific property. Supports pagination through query parameters.
// @Tags         Review
// @Produce      json
// @Param        propertyID  path      int     true  "Property ID"
// @Param        page        query     int     false "Page number for pagination"
// @Param        pageSize    query     int     false "Page size for pagination"
// @Success      200         {object}  map[string]string"Reviews retrieved successfully"
// @Failure      400         {object}  map[string]string"Invalid property ID or pagination parameters"
// @Failure      500         {object}  map[string]string"Internal server error"
// @Router       /propertyReview/get/{propertyID} [get]
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
