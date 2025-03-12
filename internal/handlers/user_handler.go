package handlers

import (
	"LeaseEase/internal/dtos"
	"LeaseEase/internal/services"
	"LeaseEase/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type userHandler struct {
	UserService services.UserService
}

func NewUserHandler(service services.UserService) *userHandler {
	return &userHandler{
		UserService: service,
	}
}

// UpdateUser godoc
// @Summary      Update user information
// @Description  Updates user details for the authenticated user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        body  body  dtos.UpdateUserDTO  true  "User update information"
// @Success      200  {object}  utils.Response  "User updated successfully"
// @Failure      400  {object}  utils.Response  "Invalid request"
// @Failure      500  {object}  utils.Response  "Failed to update user"
// @Router       /user/user [put]
// @Security     CookieAuth
func (h *userHandler) UpdateUser(c *fiber.Ctx) error {
	userID := uint(c.Locals("user").(jwt.MapClaims)["user_id"].(float64))
	var user dtos.UpdateUserDTO
	if err := c.BodyParser(&user); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid request")
	}
	if err := h.UserService.UpdateUser(userID, user); err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to update user")
	}
	return utils.SuccessResponse(c, fiber.StatusOK, "User updated successfully", nil)
}

// UpdateImage godoc
// @Summary      Update user profile image
// @Description  Updates the profile image for the authenticated user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        body  body  dtos.UpdateImageDTO  true  "Image update information"
// @Success      200  {object}  utils.Response  "Image updated successfully"
// @Failure      400  {object}  utils.Response  "Invalid request"
// @Failure      500  {object}  utils.Response  "Failed to update image"
// @Router       /user/image [put]
// @Security     CookieAuth
func (h *userHandler) UpdateImage(c *fiber.Ctx) error {
	userID := uint(c.Locals("user").(jwt.MapClaims)["user_id"].(float64))
	var image dtos.UpdateImageDTO
	if err := c.BodyParser(&image); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid request")
	}
	if err := h.UserService.UpdateImage(userID, image); err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to update image")
	}
	return utils.SuccessResponse(c, fiber.StatusOK, "Image updated successfully", nil)
}

// CheckUser godoc
// @Summary      Verify user authentication
// @Description  Validates the auth token from cookies and returns user information if authenticated
// @Tags         users
// @Accept       json
// @Produce      json
// @Success      200  {object}  utils.Response  "User is authenticated with claims data"
// @Failure      401  {object}  utils.Response  "Unauthorized - No token provided or invalid token"
// @Router       /user/check [post]
// @Security     CookieAuth
func (h *userHandler) CheckUser(c *fiber.Ctx) error {
	// Retrieve the JWT token from cookies
	token := c.Cookies("auth_token")
	if token == "" {
		return utils.ErrorResponse(c, fiber.StatusUnauthorized, "No token provided")
	}

	// Validate token using the AuthCheck service
	claims, err := h.UserService.CheckUser(token)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusUnauthorized, err.Error())
	}

	// Return success response with user details
	return utils.SuccessResponse(c, fiber.StatusOK, "User is authenticated", claims)
}
