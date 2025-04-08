package services

import (
	"LeaseEase/internal/dtos"
	"LeaseEase/internal/repositories"
	"context"
	"fmt"

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

	key := fmt.Sprintf("%s/%d", imageRequest.ImageCategory, imageRequest.ID)

	logger.Info("Uploading image", zap.String("imagePath", key))

	// Call the repository to store the image
	ImageURL, err := s.ImageRepo.UploadFile(ctx, key, imageRequest.Image)
	if err != nil {
		logger.Error("Failed to upload image", zap.Error(err))
		return nil, err
	}

	// if imageRequest.ImageCategory == "profiles" {
	// 	user := &models.User{
	// 		ID:       imageRequest.ID,
	// 		ImageURL: ImageURL,
	// 	}

	// } else if imageRequest.ImageCategory == "properties" {
	// 	property := &models.Property{
	// 		ID:       imageRequest.ID,
	// 		ImageURL: ImageURL,
	// 	}
	// }

	response := &dtos.ImageUploadResponseDTO{
		ImageURL: ImageURL,
	}
	logger.Info("Image uploaded successfully", zap.String("imageURL", ImageURL))

	return response, nil
}
