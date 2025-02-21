package repositories

import (
	"LeaseEase/internal/models"
	"errors"

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

func (r *reviewRepository) UpdateReview(reviewID uint, lesseeID uint, updates *models.Review) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		var propertyReview models.PropertyReview

		// Check if the review exists and belongs to the lessee
		if err := tx.Where("review_id = ? AND lessee_id = ?", reviewID, lesseeID).
			First(&propertyReview).Error; err != nil {
			return errors.New("review not found or unauthorized")
		}

		// Update the review
		// if err := tx.Model(&models.Review{}).
		// 	Where("id = ?", reviewID).
		// 	Updates(updates).Error; err != nil {
		// 	return err
		// }
		if err := tx.Model(&models.Review{}).
			Where("id = ?", reviewID).
			Select("ReviewMessage", "Rating").
			Updates(map[string]interface{}{
				"review_message": updates.ReviewMessage,
				"rating":         updates.Rating,
			}).Error; err != nil {
			return err
		}

		return nil
	})
}
