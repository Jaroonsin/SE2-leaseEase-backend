package services

import "LeaseEase/internal/dtos"

type ReviewService interface {
	CreateReview(dto *dtos.CreateReviewDTO, lesseeID uint) error
	UpdateReview(reviewID uint, dto *dtos.UpdateReviewDTO, lesseeID uint) error
}
