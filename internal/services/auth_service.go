package services

import (
	"LeaseEase/config"
	"LeaseEase/internal/dtos"
	"LeaseEase/internal/models"
	"LeaseEase/internal/repositories"
	"LeaseEase/utils"
	"LeaseEase/utils/constant"
	"errors"
	"fmt"
	"regexp"
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
			Email:     registerDTO.Email,
			Password:  registerDTO.Password,
			Name:      registerDTO.Name,
			Address:   registerDTO.Address,
			CreatedAt: time.Now(),
			UserType:  registerDTO.Role,
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

	re := regexp.MustCompile(`^john\.doe(?:[1-9][0-9]?)?@example\.com$`)
	dev := config.LoadConfig().ServerEnv == "development"
	if dev && re.MatchString(requestOTPDTO.Email) {
		otp = "123456"
	}

	s.authRepo.SaveOTP(models.OTP{
		Email:    requestOTPDTO.Email,
		OTP:      otp,
		ExpireAt: expiry,
	})

	if !dev || !re.MatchString(requestOTPDTO.Email) {
		if err := utils.SendOTP(requestOTPDTO.Email, otp); err != nil {
			logger.Error("failed to send OTP", zap.Error(err))
			return errors.New("failed to send OTP")
		}
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

func (s *authService) RequestPasswordReset(resetPassRequestDTO *dtos.ResetPassRequestDTO) (string, error) {
	logger := s.logger.Named("RequestPasswordReset")
	email := resetPassRequestDTO.Email

	user, err := s.authRepo.FindByEmail(email)
	if err != nil {
		logger.Warn("User not found", zap.String("email", email))
		return "", errors.New("user not found")
	}

	token, err := utils.GenerateSecureToken()
	if err != nil {
		logger.Error("Failed to generate reset token", zap.Error(err))
		return "", errors.New("could not generate token")
	}

	expiry := time.Now().Add(15 * time.Minute)
	if err := s.authRepo.SaveResetToken(user, token, expiry); err != nil {
		logger.Error("Failed to save reset token", zap.Error(err))
		return "", errors.New("could not save reset token")
	}

	resetPassURL := config.LoadConfig().ResetPassURL
	resetLink := fmt.Sprintf("%s?email=%s&token=%s", resetPassURL, email, token)
	logger.Info("Generated reset link", zap.String("email", email), zap.String("reset_link", resetLink))

	return resetLink, nil
}

func (s *authService) ResetPassword(resetPassDTO *dtos.ResetPassDTO) error {
	logger := s.logger.Named("ResetPassword")
	email := resetPassDTO.Email
	token := resetPassDTO.Token
	newPassword := resetPassDTO.Password

	user, err := s.authRepo.FindByEmail(email)
	if err != nil {
		logger.Warn("Invalid reset attempt - user not found", zap.String("email", email))
		return errors.New("invalid request")
	}

	if time.Now().After(user.TokenExpiry) {
		logger.Warn("Expired reset token", zap.String("email", email))
		return errors.New("token expired")
	}

	hashedInputToken := utils.HashToken(token)
	if hashedInputToken != user.ResetToken {
		logger.Warn("Invalid reset token", zap.String("email", email))
		return errors.New("invalid token")
	}

	hashedPassword, err := utils.HashPassword(newPassword)
	if err != nil {
		logger.Error("Failed to hash new password", zap.Error(err))
		return errors.New("could not hash password")
	}

	if err := s.authRepo.UpdateUserPassword(user, hashedPassword); err != nil {
		logger.Error("Failed to update user password", zap.String("email", email), zap.Error(err))
		return errors.New("could not update password")
	}

	logger.Info("Password reset successful", zap.String("email", email))
	return nil
}
