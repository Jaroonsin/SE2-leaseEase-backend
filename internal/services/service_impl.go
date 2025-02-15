package services

import "LeaseEase/internal/repositories"

type service struct {
	PropertyService PropertyService
	AuthService     AuthService
	RequestService  RequestService
}

func NewService(repo repositories.Repository) Service {
	return &service{
		PropertyService: NewPropertyService(repo.Property()),
		AuthService:     NewAuthService(repo.Auth()),
		RequestService:  NewRequestService(repo.Request()),
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