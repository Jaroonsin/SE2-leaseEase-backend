package services

import (
	"LeaseEase/internal/dtos"
	"LeaseEase/internal/models"
	"LeaseEase/internal/repositories"
	"LeaseEase/utils/constant"
	"errors"
	"strings"
	"time"

	"go.uber.org/zap"
)

type reviewService struct {
	reviewRepo repositories.ReviewRepository
	logger     *zap.Logger
}

func NewReviewService(reviewRepo repositories.ReviewRepository, logger *zap.Logger) ReviewService {
	return &reviewService{
		reviewRepo: reviewRepo,
		logger:     logger,
	}
}

func (s *reviewService) CreateReview(dto *dtos.CreateReviewDTO, lesseeID uint) error {
	logger := s.logger.Named("CreateReview")

	review := &models.Review{
		ReviewMessage: dto.ReviewMessage,
		Rating:        dto.Rating,
		TimeStamp:     time.Now(),
	}

	propertyReview := &models.PropertyReview{
		LesseeID:   lesseeID,
		PropertyID: dto.PropertyID,
	}

	// if err := s.reviewRepo.CreateReview(review, propertyReview); err != nil {
	// 	logger.Error("Error in create review", zap.Error(err))
	// 	return err
	// }
	err := s.reviewRepo.CreateReview(review, propertyReview)
	if err != nil {
		// Handle Foreign Key Error
		if strings.Contains(err.Error(), "foreign key constraint fails") {
			logger.Error("Property does not exist", zap.Error(err))
			return errors.New("property does not exist")
		}
		logger.Error("Error in create review", zap.Error(err))
		return err
	}

	logger.Info(constant.SuccesCreateReview, zap.String("Review Message", review.ReviewMessage))
	return nil
}

func (s *reviewService) UpdateReview(reviewID uint, dto *dtos.UpdateReviewDTO, lesseeID uint) error {
	logger := s.logger.Named("UpdateReview")

	// updates := make(map[string]interface{})

	// if dto.ReviewMessage != nil {
	// 	updates["review_message"] = *dto.ReviewMessage
	// }
	// if dto.Rating != nil {
	// 	updates["rating"] = *dto.Rating
	// }

	// if len(updates) == 0 {
	// 	return errors.New("no fields to update")
	// }

	err := s.reviewRepo.UpdateReview(reviewID, lesseeID, &models.Review{
		ReviewMessage: dto.ReviewMessage,
		Rating:        dto.Rating,
	})
	if err != nil {
		logger.Error("Error updating review", zap.Error(err))
		return err
	}

	logger.Info("Review updated successfully", zap.Uint("ReviewID", reviewID))
	return nil
}
