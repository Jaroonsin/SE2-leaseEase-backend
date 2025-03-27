package repositories

import "LeaseEase/internal/models"

type UserRepository interface {
	UpdateUser(User *models.User) error
	GetUserByID(userID uint) (*models.User, error)
}
