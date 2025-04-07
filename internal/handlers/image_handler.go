package handlers

import (
	"LeaseEase/internal/dtos"
	"LeaseEase/internal/services"
	"LeaseEase/utils"

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
// @Accept json
// @Produce json
// @Param imageRequest body dtos.ImageUploadRequestDTO true "Image data"
// @Success 200 {object} utils.Response "Image uploaded successfully"
// @Failure 400 {object} utils.Response "Invalid request payload"
// @Failure 500 {object} utils.Response "Failed to upload image"
// @Router /images/upload [post]
func (h *imageHandler) UploadImage(c *fiber.Ctx) error {
	key := c.FormValue("key")

	fileHeader, err := c.FormFile("image")
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Failed to retrieve image file")
	}

	err = c.SaveFile(fileHeader, "./uploads/"+fileHeader.Filename)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to save image file")
	}

	imageRequest := dtos.ImageUploadRequestDTO{
		ImageKey: key,
		Image:    fileHeader,
	}

	// 4. Pass the DTO to your service
	response, err := h.ImageService.UploadImage(c.Context(), &imageRequest)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to upload image")
	}

	return utils.SuccessResponse(c, fiber.StatusOK, "Image uploaded successfully", response)
}
