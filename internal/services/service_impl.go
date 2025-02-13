package services

import "LeaseEase/internal/repositories"

type service struct {
	PropertyService PropertyService
	AuthService     AuthService
}

func NewService(repo repositories.Repository) Service {
	return &service{
		PropertyService: NewPropertyService(repo.Property()),
		AuthService:     NewAuthService(repo.Auth()),
	}
}


func (s *service) Property() PropertyService {
	return s.PropertyService
}

func (s *service) Auth() AuthService {
	return s.AuthService
}
