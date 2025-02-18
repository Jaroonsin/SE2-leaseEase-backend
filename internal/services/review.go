package services

import "LeaseEase/internal/dtos"

type ReviewService interface {
	CreateReview(dto *dtos.CreateReviewDTO, lesseeID uint) error
}
