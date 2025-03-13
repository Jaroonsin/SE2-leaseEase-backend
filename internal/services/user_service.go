package services

import (
	"LeaseEase/internal/dtos"
	"LeaseEase/internal/models"
	"LeaseEase/internal/repositories"
	"LeaseEase/utils"
	"errors"
	"log"

	"go.uber.org/zap"
)

type userService struct {
	UserRepo repositories.UserRepository
	logger   *zap.Logger
}

func NewUserService(repo repositories.UserRepository, logger *zap.Logger) UserService {
	return &userService{
		UserRepo: repo,
		logger:   logger,
	}
}

func (s *userService) UpdateUser(userID uint, User dtos.UpdateUserDTO) error {

	user := models.User{
		ID:      userID,
		Name:    User.Name,
		Address: User.Address,
	}

	return s.UserRepo.UpdateUser(&user)
}

func (s *userService) UpdateImage(userID uint, Image dtos.UpdateImageDTO) error {
	logger := s.logger.Named("UpdateImage")
	if Image.ImageURL == "" {
		logger.Error("no image URL provided")
		return errors.New("no image URL provided")
	}
	log.Print("Image URL: ", Image.ImageURL)

	user := models.User{
		ID:       userID,
		ImageURL: Image.ImageURL,
	}

	logger.Info("updating user image", zap.Uint("UserID", userID), zap.String("ImageURL", Image.ImageURL))
	return s.UserRepo.UpdateUser(&user)
}

func (s *userService) CheckUser(token string) (*dtos.CheckUserDTO, error) {
	logger := s.logger.Named("CheckUser")
	if token == "" {
		logger.Error("no token provided")
		return nil, errors.New("no token provided")
	}

	claims, err := utils.ParseJWT(token)
	if err != nil {
		logger.Error("invalid or expired token", zap.Error(err))
		return nil, errors.New("invalid or expired token")
	}

	userIDFloat, ok1 := claims["user_id"].(float64)
	role, ok2 := claims["role"].(string)
	if !ok1 || !ok2 {
		logger.Error("invalid token payload")
		return nil, errors.New("invalid token payload")
	}
	if role != "lessor" && role != "lessee" {
		logger.Error("invalid role")
		return nil, errors.New("invalid role")
	}

	user, err := s.UserRepo.GetUserByID(uint(userIDFloat))
	if err != nil {
		logger.Error("user not found", zap.Error(err))
		return nil, errors.New("user not found")
	}

	userID := uint(userIDFloat)
	logger.Info("user authenticated", zap.Uint("UserID", userID), zap.String("Role", role))
	return &dtos.CheckUserDTO{
		UserID:   userID,
		Role:     role,
		Email:    user.Email,
		Name:     user.Name,
		Address:  user.Address,
		ImageURL: user.ImageURL,
	}, nil
}

func (s *userService) GetUser(userID uint) (*dtos.GetUserDTO, error) {
	logger := s.logger.Named("GetUser")

	user, err := s.UserRepo.GetUserByID(userID)
	if err != nil {
		logger.Error("user not found", zap.Error(err))
		return nil, errors.New("user not found")
	}

	logger.Info("user found", zap.Uint("UserID", userID))
	return &dtos.GetUserDTO{
		Name:     user.Name,
		Address:  user.Address,
		ImageURL: user.ImageURL,
	}, nil
}
