package services

import (
	"LeaseEase/internal/repositories"

	"go.uber.org/zap"
)

type service struct {
	PropertyService PropertyService
	AuthService     AuthService
	RequestService  RequestService
	ReviewService   ReviewService
}

func NewService(repo repositories.Repository, logger *zap.Logger) Service {
	return &service{
		PropertyService: NewPropertyService(repo.Property(), logger),
		AuthService:     NewAuthService(repo.Auth(), logger),
		RequestService:  NewRequestService(repo.Request(), logger),
		ReviewService:   NewReviewService(repo.Review(), logger),
	}
}

func (s *service) Property() PropertyService {
	return s.PropertyService
}

func (s *service) Auth() AuthService {
	return s.AuthService
}

func (s *service) Request() RequestService {
	return s.RequestService
}

func (s *service) Review() ReviewService {
	return s.ReviewService
}
