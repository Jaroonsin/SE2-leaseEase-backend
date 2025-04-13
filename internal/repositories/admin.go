package repositories

import "LeaseEase/internal/models"

type AdminRepository interface {
	GetAllUsers(limit, offset int) ([]models.User, error)
	GetUserByID(id int) (models.User, error)

	CountReviewsByPropertyForAdmin(queryString string, totalRecords *int64) error

	GetAllReviewsSortedByPropertyName(query, direction string) ([]models.PropertyReview, error)
	GetAllReviewsSortedByReviewer(query, direction string) ([]models.PropertyReview, error)
	GetAllReviewsSortedByTime(query, direction string) ([]models.PropertyReview, error)

	GetPaginatedReviewsSortedByPropertyName(limit, offset int, query, direction string) ([]models.PropertyReview, error)
	GetPaginatedReviewsSortedByReviewer(limit, offset int, query, direction string) ([]models.PropertyReview, error)
	GetPaginatedReviewsSortedByTime(limit, offset int, query, direction string) ([]models.PropertyReview, error)

	UpdateUserStatus(id int, status string) error

	DeleteReviewByID(id int) error
}
