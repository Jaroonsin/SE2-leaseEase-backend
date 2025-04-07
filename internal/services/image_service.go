package services

import (
	"LeaseEase/internal/dtos"
	"LeaseEase/internal/repositories"
	"context"

	"go.uber.org/zap"
)

type imageService struct {
	ImageRepo repositories.ImageRepository
	logger    *zap.Logger
}

func NewImageService(imageRepo repositories.ImageRepository, logger *zap.Logger) ImageService {
	return &imageService{
		ImageRepo: imageRepo,
		logger:    logger,
	}
}

// UploadImage handles the image upload process
func (s *imageService) UploadImage(ctx context.Context, imageRequest *dtos.ImageUploadRequestDTO) (*dtos.ImageUploadResponseDTO, error) {
	logger := s.logger.Named("UploadImage")
	logger.Info("Uploading image", zap.String("imagePath", imageRequest.ImageKey))


	// Call the repository to store the image
	ImageURL, err := s.ImageRepo.UploadFile(ctx, imageRequest.ImageKey, imageRequest.Image)
	if err != nil {
		logger.Error("Failed to upload image", zap.Error(err))
		return nil, err
	}

	response := &dtos.ImageUploadResponseDTO{
		ImageURL: ImageURL,
	}
	logger.Info("Image uploaded successfully", zap.String("imageURL", ImageURL))

	return response, nil
}
