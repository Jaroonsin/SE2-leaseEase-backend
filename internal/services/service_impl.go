package services

import (
	"LeaseEase/internal/repositories"

	"go.uber.org/zap"
)

type service struct {
	PropertyService PropertyService
	AuthService     AuthService
	LesseeService   LesseeService
	ReviewService   ReviewService
	PaymentService  PaymentService
}

func NewService(repo repositories.Repository, logger *zap.Logger) Service {
	return &service{
		PropertyService: NewPropertyService(repo.Property(), logger),
		AuthService:     NewAuthService(repo.Auth(), logger),
		LesseeService:   NewLesseeService(repo.Lessee(), logger),
		ReviewService:   NewReviewService(repo.Review(), logger),
		PaymentService:  NewPaymentService(repo.Payment(), logger),
	}
}

func (s *service) Property() PropertyService {
	return s.PropertyService
}

func (s *service) Auth() AuthService {
	return s.AuthService
}

func (s *service) Lessee() LesseeService {
	return s.LesseeService
}

func (s *service) Review() ReviewService {
	return s.ReviewService
}

func (s *service) Payment() PaymentService {
	return s.PaymentService
}
