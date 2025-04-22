package services

import (
	"LeaseEase/internal/dtos"
	"context"
)

type ImageService interface {
	UploadImage(ctx context.Context, imageRequest *dtos.ImageUploadRequestDTO) (*dtos.ImageUploadResponseDTO, error)
}
