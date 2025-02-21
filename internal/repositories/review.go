package repositories

import "LeaseEase/internal/models"

type ReviewRepository interface {
	CreateReview(review *models.Review, propertyReview *models.PropertyReview) error
	UpdateReview(reviewID uint, lesseeID uint, updates *models.Review) error
}
