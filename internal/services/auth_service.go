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
	logger := s.logger.Named("Register tempUsers")

	if s.authRepo.FindEmailExisted(registerDTO.Email) {
		logger.Error("Email already existed", zap.String("Email", registerDTO.Email))
		return errors.New("email already existed")
	}

	hashedPassword, err := utils.HashPassword(registerDTO.Password)
	if err != nil {
		logger.Error("cannot hash password", zap.Error(err))
		return err
	}
	registerDTO.Password = hashedPassword

	s.authRepo.SaveTempUser(models.TempUser{
		User: &models.User{
			Email:    registerDTO.Email,
			Password: registerDTO.Password,
			Name:     registerDTO.Name,
			Address:  registerDTO.Address,
			Birthday: time.Now(),
			UserType: registerDTO.Role,
		},
		ExpireAt: time.Now().Add(30 * time.Minute),
	})

	logger.Info("Temporary user created", zap.String("Email", registerDTO.Email))
	return nil
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

func (s *authService) AuthCheck(token string) (*dtos.AuthCheckDTO, error) {
	logger := s.logger.Named("AuthCheck")
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

	userID := uint(userIDFloat)
	logger.Info("user authenticated", zap.Uint("UserID", userID), zap.String("Role", role))
	return &dtos.AuthCheckDTO{
		UserID: userID,
		Role:   role,
	}, nil
}

func (s *authService) RequestOTP(requestOTPDTO *dtos.RequestOTPDTO) error {
	logger := s.logger.Named("RequestOTP")
	otp := utils.GenerateOTP()
	expiry := time.Now().Add(3 * time.Minute)

	s.authRepo.SaveOTP(models.OTP{
		Email:    requestOTPDTO.Email,
		OTP:      otp,
		ExpireAt: expiry,
	})

	if err := utils.SendOTP(requestOTPDTO.Email, otp); err != nil {
		logger.Error("failed to send OTP", zap.Error(err))
		return errors.New("failed to send OTP")
	}
	logger.Info("OTP sent", zap.String("Email", requestOTPDTO.Email))
	return nil
}

func (s *authService) VerifyOTP(verifyOTPDTO *dtos.VerifyOTPDTO) error {
	logger := s.logger.Named("VerifyOTP")
	otpData, exists := s.authRepo.FindByEmailOTP(verifyOTPDTO.Email)

	if !exists || time.Now().After(otpData.ExpireAt) {
		s.authRepo.DeleteOTP(verifyOTPDTO.Email)
		logger.Error("OTP is invalid or expired", zap.String("Email", verifyOTPDTO.Email))
		return errors.New("OTP is invalid or expired")
	}

	if otpData.OTP != verifyOTPDTO.OTP {
		s.authRepo.DeleteOTP(verifyOTPDTO.Email)
		logger.Error("invalid OTP", zap.String("Email", verifyOTPDTO.Email))
		return errors.New("invalid OTP")
	}

	s.authRepo.DeleteOTP(verifyOTPDTO.Email)

	tempUser, exists := s.authRepo.FindByEmailTempUser(verifyOTPDTO.Email)
	if !exists {
		logger.Error("Temporary user not found", zap.String("Email", verifyOTPDTO.Email))
		return errors.New("temporary user not found")
	}

	user := tempUser.User

	s.authRepo.DeleteTempUser(verifyOTPDTO.Email)
	logger.Info(constant.SuccessRegister, zap.String("Email", user.Email))
	return s.authRepo.CreateUser(user)
}