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
		print(review.ID)
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

// func (r *reviewRepository) GetAllReviewsForAdmin(queryString string, sortParam string, direction string) ([]models.PropertyReview, error) {
// 	var propertyReviews []models.PropertyReview
// 	err := r.db.Preload("Review").Preload("Lessee").
// 		Where("property_id = ?", queryString).
// 		Order(sortParam + " " + direction).
// 		Find(&propertyReviews).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	return propertyReviews, nil
// }

func (r *reviewRepository) CountReviewsByPropertyForAdmin(queryString string, totalRecords *int64) error {
	return r.db.Model(&models.PropertyReview{}).
		Joins("JOIN properties ON properties.id = property_reviews.property_id").
		Where("properties.name ILIKE ?", "%"+queryString+"%").
		Count(totalRecords).Error
}

// func (r *reviewRepository) GetPaginatedReviewsForAdmin(page int, pageSize int, queryString string, sortParam string, direction string) ([]models.PropertyReview, error) {
// 	var propertyReviews []models.PropertyReview
// 	err := r.db.Preload("Review").Preload("Lessee").
// 		Where("property_id = ?", queryString).
// 		Order(sortParam + " " + direction).
// 		Limit(pageSize).Offset((page - 1) * pageSize).
// 		Find(&propertyReviews).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	return propertyReviews, nil
// }

func (r *reviewRepository) GetAllReviewsSortedByPropertyName(query, direction string) ([]models.PropertyReview, error) {
	var propertyReviews []models.PropertyReview
	if direction != "ASC" && direction != "DESC" {
		direction = "ASC"
	}

	err := r.db.
		Joins("JOIN properties ON properties.id = property_reviews.property_id").
		Preload("Review").
		Preload("Lessee").
		Where("properties.name ILIKE ?", "%"+query+"%").
		Order("properties.name " + direction).
		Find(&propertyReviews).Error

	return propertyReviews, err
}

func (r *reviewRepository) GetAllReviewsSortedByReviewer(query, direction string) ([]models.PropertyReview, error) {
	var propertyReviews []models.PropertyReview
	if direction != "ASC" && direction != "DESC" {
		direction = "ASC"
	}

	err := r.db.
		Joins("JOIN users ON users.id = property_reviews.lessee_id").
		Preload("Review").
		Preload("Lessee").
		Preload("Property").
		Where("users.name ILIKE ?", "%"+query+"%").
		Order("users.name " + direction).
		Find(&propertyReviews).Error

	return propertyReviews, err
}

func (r *reviewRepository) GetAllReviewsSortedByTime(query, direction string) ([]models.PropertyReview, error) {
	var propertyReviews []models.PropertyReview

	if direction != "ASC" && direction != "DESC" {
		direction = "ASC"
	}

	err := r.db.
		Joins("JOIN properties ON properties.id = property_reviews.property_id").
		Joins("JOIN reviews ON reviews.id = property_reviews.review_id").
		Preload("Review").
		Preload("Lessee").
		Preload("Property").
		Where("properties.name ILIKE ?", "%"+query+"%").
		Order("reviews.time_stamp " + direction).
		Find(&propertyReviews).Error

	return propertyReviews, err
}

func (r *reviewRepository) GetPaginatedReviewsSortedByPropertyName(limit, offset int, query, direction string) ([]models.PropertyReview, error) {
	var propertyReviews []models.PropertyReview
	if direction != "ASC" && direction != "DESC" {
		direction = "ASC"
	}

	err := r.db.
		Joins("JOIN properties ON properties.id = property_reviews.property_id").
		Preload("Review").
		Preload("Lessee").
		Where("properties.name ILIKE ?", "%"+query+"%").
		Order("properties.name " + direction).
		Limit(limit).
		Offset(offset).
		Find(&propertyReviews).Error

	return propertyReviews, err
}

func (r *reviewRepository) GetPaginatedReviewsSortedByReviewer(limit, offset int, query, direction string) ([]models.PropertyReview, error) {
	var propertyReviews []models.PropertyReview
	if direction != "ASC" && direction != "DESC" {
		direction = "ASC"
	}

	err := r.db.
		Joins("JOIN lessees ON users.id = property_reviews.lessee_id").
		Preload("Review").
		Preload("Lessee").
		Where("users.name ILIKE ?", "%"+query+"%").
		Order("users.name " + direction).
		Limit(limit).
		Offset(offset).
		Find(&propertyReviews).Error

	return propertyReviews, err
}

func (r *reviewRepository) GetPaginatedReviewsSortedByTime(limit, offset int, query, direction string) ([]models.PropertyReview, error) {
	var propertyReviews []models.PropertyReview
	if direction != "ASC" && direction != "DESC" {
		direction = "ASC"
	}

	err := r.db.
		Joins("JOIN properties ON properties.id = property_reviews.property_id").
		Joins("JOIN reviews ON reviews.id = property_reviews.review_id").
		Preload("Review").
		Preload("Lessee").
		Preload("Property").
		Where("properties.name ILIKE ?", "%"+query+"%").
		Order("reviews.time_stamp " + direction).
		Limit(limit).
		Offset(offset).
		Find(&propertyReviews).Error

	return propertyReviews, err
}
