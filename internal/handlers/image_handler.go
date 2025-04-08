package handlers

import (
	"LeaseEase/internal/dtos"
	"LeaseEase/internal/services"
	"LeaseEase/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type imageHandler struct {
	ImageService services.ImageService
}

func NewImageHandler(imageService services.ImageService) *imageHandler {
	return &imageHandler{
		ImageService: imageService,
	}
}

// UploadImage handles requests to upload images to the system.
// @Summary Upload an image
// @Description Upload an image to the system
// @Tags images
// @Accept mpfd
// @Produce mpfd
// @Param id formData string true "ID of the related entity"
// @Param category formData string true "Image category (profiles or properties)"
// @Param image formData file true "Image file"
// @Success 200 {object} utils.Response "Image uploaded successfully"
// @Failure 400 {object} utils.Response "Invalid request payload"
// @Failure 500 {object} utils.Response "Failed to upload image"
// @Router /images/upload [post]
func (h *imageHandler) UploadImage(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.FormValue("id"))
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid ID")
	}

	category := c.FormValue("category")

	if category != "profiles" && category != "properties" {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid image category")
	}

	fileHeader, err := c.FormFile("image")

	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Failed to retrieve image file")
	}

	imageRequest := dtos.ImageUploadRequestDTO{
		ID:            uint(id),
		ImageCategory: category,
		Image:         fileHeader,
	}

	// 4. Pass the DTO to your service
	response, err := h.ImageService.UploadImage(c.Context(), &imageRequest)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to upload image")
	}

	return utils.SuccessResponse(c, fiber.StatusOK, "Image uploaded successfully", response)
}
