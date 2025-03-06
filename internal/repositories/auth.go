package repositories

import (
	"LeaseEase/internal/models"
	"time"
)

type AuthRepository interface {
	FindEmailExisted(email string) bool
	CreateUser(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
	SaveTempUser(user models.TempUser)
	FindByEmailTempUser(email string) (models.TempUser, bool)
	DeleteTempUser(email string)
	SaveOTP(otp models.OTP)
	FindByEmailOTP(email string) (models.OTP, bool)
	DeleteOTP(email string)
	FindByEmail(email string) (*models.User, error)
	SaveResetToken(user *models.User, token string, expiry time.Time) error
	UpdateUserPassword(user *models.User, hashedPassword string) error
}
