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

func (r *reviewRepository) DeleteReview(reviewID uint, lesseeID uint) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		var propertyReview models.PropertyReview

		if err := tx.Where("review_id = ? AND lessee_id = ?", reviewID, lesseeID).
			First(&propertyReview).Error; err != nil {
			return errors.New("review not found or unauthorized")
		}

		// Delete PropertyReview first
		if err := tx.Where("review_id = ?", reviewID).
			Delete(&models.PropertyReview{}).Error; err != nil {
			return err
		}

		// Delete Review
		if err := tx.Where("id = ?", reviewID).
			Delete(&models.Review{}).Error; err != nil {
			return err
		}

		return nil
	})
}

func (r *reviewRepository) GetAllReviews(propertyID uint) ([]models.PropertyReview, error) {
	var propertyReviews []models.PropertyReview
	err := r.db.Preload("Review").Preload("Lessee").
		Where("property_id = ?", propertyID).
		Find(&propertyReviews).Error
	if err != nil {
		return nil, err
	}
	return propertyReviews, nil
}

func (r *reviewRepository) GetPaginatedReviews(propertyID uint, limit, offset int) ([]models.PropertyReview, error) {
	var propertyReviews []models.PropertyReview
	err := r.db.Preload("Review").Preload("Lessee").
		Where("property_id = ?", propertyID).
		Limit(limit).Offset(offset).
		Find(&propertyReviews).Error
	if err != nil {
		return nil, err
	}
	return propertyReviews, nil
}

func (r *reviewRepository) CountReviewsByProperty(propertyID uint, totalRecords *int64) error {
	return r.db.Model(&models.PropertyReview{}).
		Where("property_id = ?", propertyID).
		Count(totalRecords).Error
}
