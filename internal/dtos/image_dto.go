package dtos

import "mime/multipart"

type ImageUploadRequestDTO struct {
	ID            uint                  `json:"image_key" validate:"required"`
	ImageCategory string                `json:"image_category" validate:"required"`
	Image         *multipart.FileHeader `json:"image"`
}

type ImageUploadResponseDTO struct {
	ImageURL string `json:"image_url"`
}
