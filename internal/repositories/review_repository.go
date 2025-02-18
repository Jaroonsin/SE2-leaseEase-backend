package repositories

import (
	"LeaseEase/internal/models"

	"gorm.io/gorm"
)

type reviewRepository struct {
	db *gorm.DB
}

func NewReviewRepository(db *gorm.DB) ReviewRepository {
	return &reviewRepository{
		db: db,
	}
}

func (r *reviewRepository) CreateReview(review *models.Review, propertyReview *models.PropertyReview) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		// Insert Review
		if err := tx.Create(review).Error; err != nil {
			return err
		}

		// Insert PropertyReview
		propertyReview.ReviewID = review.ID
		if err := tx.Create(propertyReview).Error; err != nil {
			return err
		}

		return nil
	})
}
