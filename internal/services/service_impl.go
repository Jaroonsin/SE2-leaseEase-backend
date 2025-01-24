package services

import "LeaseEase/internal/repositories"

type service struct {
	UserService UserService
}

func NewService(repo repositories.Repository) Service {
	return &service{
		UserService: NewUserService(repo.User()),
	}
}

func (s *service) User() UserService {
	return s.UserService
}
