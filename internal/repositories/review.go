package repositories

import "LeaseEase/internal/models"

type ReviewRepository interface {
	CreateReview(review *models.Review, propertyReview *models.PropertyReview) error
	UpdateReview(reviewID uint, lesseeID uint, updates *models.Review) error
	DeleteReview(reviewID uint, lesseeID uint) error
	GetAllReviews(propertyID uint) ([]models.PropertyReview, error)
	GetPaginatedReviews(propertyID uint, limit, offset int) ([]models.PropertyReview, error)
	CountReviewsByProperty(propertyID uint, totalRecords *int64) error

	// GetAllReviewsForAdmin(queryString string, sortParam string, direction string) ([]models.PropertyReview, error)
	CountReviewsByPropertyForAdmin(queryString string, totalRecords *int64) error
	// GetPaginatedReviewsForAdmin(page int, pageSize int, queryString string, sortParam string, direction string) ([]models.PropertyReview, error)

	GetAllReviewsSortedByPropertyName(query, direction string) ([]models.PropertyReview, error)
	GetAllReviewsSortedByReviewer(query, direction string) ([]models.PropertyReview, error)
	GetAllReviewsSortedByTime(query, direction string) ([]models.PropertyReview, error)

	GetPaginatedReviewsSortedByPropertyName(limit, offset int, query, direction string) ([]models.PropertyReview, error)
	GetPaginatedReviewsSortedByReviewer(limit, offset int, query, direction string) ([]models.PropertyReview, error)
	GetPaginatedReviewsSortedByTime(limit, offset int, query, direction string) ([]models.PropertyReview, error)
}
