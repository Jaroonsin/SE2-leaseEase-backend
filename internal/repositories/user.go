package repositories

import "LeaseEase/internal/models"

type UserRepository interface {
	CreateUser(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
}
