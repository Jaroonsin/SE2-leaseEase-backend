package services

import (
	"LeaseEase/internal/models"
	"LeaseEase/internal/repositories"
	"LeaseEase/utils"
	"errors"
)

type userService struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (s *userService) Register(email, password, role string) error {
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return err
	}

	user := &models.User{
		Email:    email,
		Password: hashedPassword,
		Role:     role,
	}

	return s.userRepo.CreateUser(user)
}

func (s *userService) Login(email, password string) (string, error) {
	user, err := s.userRepo.GetUserByEmail(email)
	if err != nil {
		return "", errors.New("user not found")
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return "", errors.New("invalid password")
	}

	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}
