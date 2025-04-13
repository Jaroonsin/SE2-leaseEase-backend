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
// @Param        name     query     string  false "Property name to filter reviews"
// @Param        sort     query     string  false "Sort parameter"
// @Param        dir      query     string  false "Sort direction (asc/desc)"
// @Success      200      {object}  map[string]interface{} "Reviews retrieved successfully"
// @Failure      400      {object}  map[string]string      "Invalid pagination parameters"
// @Failure      500      {object}  map[string]string      "Internal server error"
// @Router       /admin/get-reviews [get]
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

// GetAllUsers godoc
// @Summary      Retrieve all users
// @Description  Get all users with pagination support.
// @Tags         User
// @Produce      json
// @Param        page     query     int     false "Page number for pagination"
// @Param        pageSize query     int     false "Page size for pagination"
// @Success      200      {object}  map[string]interface{} "Users retrieved successfully"
// @Failure      400      {object}  map[string]string      "Invalid pagination parameters"
// @Failure      500      {object}  map[string]string      "Internal server error"
// @Router       /admin/get-users [get]
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

// ManageUserStatus godoc
// @Summary      Update user status
// @Description  Update the status of a user by their ID.
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        id      path      int                     true  "User ID"
// @Param        status  body      dtos.UpdateUserStatusDTO true  "Updated status"
// @Success      200     {object}  map[string]string       "User status updated successfully"
// @Failure      400     {object}  map[string]string       "Invalid user ID or request body"
// @Failure      500     {object}  map[string]string       "Internal server error"
// @Router       /admin/manage-user/{id} [patch]
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

// DeleteReview godoc
// @Summary      Delete a review
// @Description  Delete a review by its ID.
// @Tags         Review
// @Produce      json
// @Param        id   path      int     true  "Review ID"
// @Success      200  {object}  map[string]string "Review deleted successfully"
// @Failure      400  {object}  map[string]string "Invalid review ID"
// @Failure      500  {object}  map[string]string "Internal server error"
// @Router       /admin/delete-review/{id} [delete]
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
