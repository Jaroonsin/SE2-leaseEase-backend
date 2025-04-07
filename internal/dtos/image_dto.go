package dtos

import "mime/multipart"

type ImageUploadRequestDTO struct {
	ImageKey string                `json:"image_key" validate:"required"`
	Image    *multipart.FileHeader `json:"image"`
}

type ImageUploadResponseDTO struct {
	ImageURL string `json:"image_url"`
}
