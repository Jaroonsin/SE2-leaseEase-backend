package repositories

import "LeaseEase/internal/models"

type ReviewRepository interface {
	CreateReview(review *models.Review, propertyReview *models.PropertyReview) error
	UpdateReview(reviewID uint, lesseeID uint, updates *models.Review) error
	DeleteReview(reviewID uint, lesseeID uint) error
	GetAllReviews(propertyID uint) ([]models.PropertyReview, error)
	GetPaginatedReviews(propertyID uint, limit, offset int) ([]models.PropertyReview, error)
	CountReviewsByProperty(propertyID uint, totalRecords *int64) error
}
