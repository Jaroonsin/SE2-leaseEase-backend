package services

import (
	"LeaseEase/internal/dtos"
	"LeaseEase/internal/models"
	"LeaseEase/internal/repositories"
	"LeaseEase/utils"
	"errors"
	"time"
)

type userService struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (s *userService) Register(registerDTO *dtos.RegisterDTO) error {
	hashedPassword, err := utils.HashPassword(registerDTO.Password)
	if err != nil {
		return err
	}

	user := &models.User{
		ID:       registerDTO.ID,
		Name:     registerDTO.Name,
		Address:  registerDTO.Address,
		Birthday: time.Now(),
		Email:    registerDTO.Email,
		Password: hashedPassword,
		UserType: registerDTO.Role,
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
