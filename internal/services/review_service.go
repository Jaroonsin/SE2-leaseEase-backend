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

func (s *reviewService) DeleteReview(reviewID uint, lesseeID uint) error {
	logger := s.logger.Named("DeleteReview")

	err := s.reviewRepo.DeleteReview(reviewID, lesseeID)
	if err != nil {
		logger.Error("Error deleting review", zap.Error(err))
		return err
	}

	logger.Info("Review deleted successfully", zap.Uint("ReviewID", reviewID))
	return nil
}

func (s *reviewService) GetAllReviews(propertyID uint, page, pageSize int) (*dtos.GetReviewPaginatedDTO, error) {
	logger := s.logger.Named("GetAllReviews")
	var propertyReviews []models.PropertyReview
	var totalRecords int64
	var err error

	// Case 1: Fetch all reviews (when no pagination)
	if page == 0 || pageSize == 0 {
		propertyReviews, err = s.reviewRepo.GetAllReviews(propertyID)
		if err != nil {
			logger.Error("Failed to fetch all reviews", zap.Error(err))
			return nil, err
		}
		totalRecords = int64(len(propertyReviews))
	} else {
		// Case 2: Apply pagination
		err = s.reviewRepo.CountReviewsByProperty(propertyID, &totalRecords)
		if err != nil {
			return nil, err
		}

		offset := (page - 1) * pageSize
		propertyReviews, err = s.reviewRepo.GetPaginatedReviews(propertyID, pageSize, offset)
		if err != nil {
			logger.Error("Failed to fetch paginated reviews", zap.Int("page", page), zap.Int("pageSize", pageSize), zap.Error(err))
			return nil, err
		}
	}

	totalPages := 1
	if pageSize > 0 {
		totalPages = int((totalRecords + int64(pageSize) - 1) / int64(pageSize))
	}

	// Convert to DTO
	var reviewDTOs []dtos.GetReviewDTO
	for _, pr := range propertyReviews {
		reviewDTOs = append(reviewDTOs, dtos.GetReviewDTO{
			ReviewID:      pr.Review.ID,
			ReviewMessage: pr.Review.ReviewMessage,
			Rating:        pr.Review.Rating,
			TimeStamp:     pr.Review.TimeStamp,
			LesseeName:    pr.Lessee.Name,
		})
	}

	logger.Info("Success fetching reviews", zap.Int("count", len(reviewDTOs)))
	return &dtos.GetReviewPaginatedDTO{
		Reviews:      reviewDTOs,
		TotalRecords: int(totalRecords),
		TotalPages:   totalPages,
		CurrentPage:  page,
		PageSize:     pageSize,
	}, nil
}
