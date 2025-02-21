package services

import "LeaseEase/internal/dtos"

type ReviewService interface {
	CreateReview(dto *dtos.CreateReviewDTO, lesseeID uint) error
	UpdateReview(reviewID uint, dto *dtos.UpdateReviewDTO, lesseeID uint) error
	DeleteReview(reviewID uint, lesseeID uint) error
	GetAllReviews(propertyID uint, page, pageSize int) (*dtos.GetReviewPaginatedDTO, error)
}
