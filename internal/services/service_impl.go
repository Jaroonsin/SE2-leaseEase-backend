package services

import (
	"LeaseEase/internal/repositories"

	"go.uber.org/zap"
)

type service struct {
	PropertyService PropertyService
	AuthService     AuthService
	RequestService  RequestService
}

func NewService(repo repositories.Repository, logger *zap.Logger) Service {
	return &service{
		PropertyService: NewPropertyService(repo.Property(), logger),
		AuthService:     NewAuthService(repo.Auth(), logger),
		RequestService:  NewRequestService(repo.Request(), logger),
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
