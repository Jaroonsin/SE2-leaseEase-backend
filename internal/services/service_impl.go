package services

import (
	"LeaseEase/internal/repositories"

	"go.uber.org/zap"
)

type service struct {
	PropertyService    PropertyService
	AuthService        AuthService
	ReservationService ReservationService
	ReviewService      ReviewService
	PaymentService     PaymentService
}

func NewService(repo repositories.Repository, logger *zap.Logger) Service {
	return &service{
		PropertyService:    NewPropertyService(repo.Property(), logger),
		AuthService:        NewAuthService(repo.Auth(), logger),
		ReservationService: NewReservationService(repo.Reservation(), logger),
		ReviewService:      NewReviewService(repo.Review(), logger),
		PaymentService:     NewPaymentService(repo.Payment(), logger),
	}
}

func (s *service) Property() PropertyService {
	return s.PropertyService
}

func (s *service) Auth() AuthService {
	return s.AuthService
}

func (s *service) Reservation() ReservationService {
	return s.ReservationService
}

func (s *service) Review() ReviewService {
	return s.ReviewService
}

func (s *service) Payment() PaymentService {
	return s.PaymentService
}
