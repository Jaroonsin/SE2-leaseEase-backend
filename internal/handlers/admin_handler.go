package handlers

import (
	"LeaseEase/internal/dtos"
	"LeaseEase/internal/services"
	"LeaseEase/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// fix path for swagger
type adminHandler struct {
	adminService services.AdminService
}

func NewAdminHandler(adminService services.AdminService) *adminHandler {
	return &adminHandler{
		adminService: adminService,
	}
}

// GetAllReviewsForAdmin godoc
// @Summary      Retrieve all reviews for admin
// @Description  Get all reviews for admin. Supports pagination through query parameters.
// @Tags         Review
// @Produce      json
// @Param        page     query     int     false "Page number for pagination"
// @Param        pageSize query     int     false "Page size for pagination"
// @Success      200      {object}  map[string]string"Reviews retrieved successfully"
// @Failure      400      {object}  map[string]string"Invalid pagination parameters"
// @Failure      500      {object}  map[string]string"Internal server error"
// @Router       /propertyReview/get [get]
func (h *adminHandler) GetAllReviewsForAdmin(c *fiber.Ctx) error {
	pageStr := c.Query("page", "")
	pageSizeStr := c.Query("pageSize", "")
	propName := c.Query("name", "")
	sortParameter := c.Query("sort", "")
	direction := c.Query("dir", "")

	if pageStr == "" && pageSizeStr == "" {
		reviews, err := h.adminService.GetAllReviewsForAdmin(0, 0, propName, sortParameter, direction)
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

	reviews, err := h.adminService.GetAllReviewsForAdmin(page, pageSize, propName, sortParameter, direction)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}
	return utils.SuccessResponse(c, fiber.StatusOK, "Success", reviews)
}

func (h *adminHandler) GetAllUsers(c *fiber.Ctx) error {
	page := c.Query("page", "1")
	pageSize := c.Query("pageSize", "10")

	pageInt, err := strconv.Atoi(page)
	if err != nil || pageInt < 1 {
		pageInt = 1
	}

	pageSizeInt, err := strconv.Atoi(pageSize)
	if err != nil || pageSizeInt < 1 {
		pageSizeInt = 10
	}

	users, err := h.adminService.GetAllUsers(pageInt, pageSizeInt)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}
	return utils.SuccessResponse(c, fiber.StatusOK, "Success", users)
}

func (h *adminHandler) ManageUserStatus(c *fiber.Ctx) error {
	idStr := c.Params("id")
	if idStr == "" {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid user ID")
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid user ID")
	}

	var updatedStatus dtos.UpdateUserStatusDTO
	err = c.BodyParser(&updatedStatus)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid request body")
	}

	err = h.adminService.ManageUserStatus(id, updatedStatus.Status)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}
	return utils.SuccessResponse(c, fiber.StatusOK, "User status updated successfully", nil)
}

func (h *adminHandler) DeleteReview(c *fiber.Ctx) error {
	idStr := c.Params("id")
	if idStr == "" {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid ID")
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Cannot convert ID to integer")
	}

	err = h.adminService.DeleteReview(id)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}
	return utils.SuccessResponse(c, fiber.StatusOK, "Review deleted successfully", nil)
}
