package services

import "LeaseEase/internal/repositories"

type service struct {
	UserService     UserService
	PropertyService PropertyService
	AuthService     AuthService
}

func NewService(repo repositories.Repository) Service {
	return &service{
		UserService:     NewUserService(repo.User()),
		PropertyService: NewPropertyService(repo.Property()),
		AuthService:     NewAuthService(repo.Auth()),
	}
}

func (s *service) User() UserService {
	return s.UserService
}

func (s *service) Property() PropertyService {
	return s.PropertyService
}

func (s *service) Auth() AuthService {
	return s.AuthService
}
