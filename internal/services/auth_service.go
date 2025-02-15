package services

import (
	"LeaseEase/internal/dtos"
	"LeaseEase/internal/models"
	"LeaseEase/internal/repositories"
	"LeaseEase/utils"
	"LeaseEase/utils/constant"
	"errors"
	"time"

	"go.uber.org/zap"
)

type authService struct {
	authRepo repositories.AuthRepository
	logger   *zap.Logger
}

func NewAuthService(authRepo repositories.AuthRepository, logger *zap.Logger) AuthService {
	return &authService{
		authRepo: authRepo,
		logger:   logger,
	}
}

func (s *authService) Register(registerDTO *dtos.RegisterDTO) error {
	logger := s.logger.Named("Register")

	hashedPassword, err := utils.HashPassword(registerDTO.Password)
	if err != nil {
		logger.Error("cannot hash password", zap.Error(err))
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

	logger.Info(constant.SuccessRegister, zap.String("Email", user.Email))
	return s.authRepo.CreateUser(user)
}

func (s *authService) Login(loginDTO *dtos.LoginDTO) (string, error) {
	logger := s.logger.Named("Login")

	user, err := s.authRepo.GetUserByEmail(loginDTO.Email)
	if err != nil {
		logger.Error(constant.ErrUserNotFound, zap.Error(err))
		return "", errors.New(constant.ErrUserNotFound)
	}

	if !utils.CheckPasswordHash(loginDTO.Password, user.Password) {
		logger.Error(constant.ErrPassNotMatch, zap.Error(err))
		return "", errors.New(constant.ErrPassNotMatch)
	}
	JWTDTO := dtos.JWTDTO{
		UserID: user.ID,
		Role:   user.UserType,
	}

	token, err := utils.GenerateJWT(JWTDTO)
	if err != nil {
		logger.Error("cannot generate JWT", zap.Error(err))
		return "", err
	}

	logger.Info(constant.SuccessLogin, zap.String("token", token))
	return token, nil
}
