package services

import (
	"LeaseEase/internal/dtos"
	"LeaseEase/internal/models"
	"LeaseEase/internal/repositories"
	"LeaseEase/utils"
	"errors"
	"time"
)

type authService struct {
	authRepo repositories.AuthRepository
}

func NewAuthService(authRepo repositories.AuthRepository) AuthService {
	return &authService{
		authRepo: authRepo,
	}
}

func (s *authService) Register(registerDTO *dtos.RegisterDTO) error {
	hashedPassword, err := utils.HashPassword(registerDTO.Password)
	if err != nil {
		return err
	}

	user := &models.User{
		Email:    registerDTO.Email,
		Password: hashedPassword,
		Name:     registerDTO.Name,
		Address:  registerDTO.Address,
		Birthday: time.Now(),
		UserType: registerDTO.Role,
	}

	return s.authRepo.CreateUser(user)
}

func (s *authService) Login(loginDTO *dtos.LoginDTO) (string, error) {
	user, err := s.authRepo.GetUserByEmail(loginDTO.Email)
	if err != nil {
		return "", errors.New("user not found")
	}

	if !utils.CheckPasswordHash(loginDTO.Password, user.Password) {
		return "", errors.New("invalid password")
	}

	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}
