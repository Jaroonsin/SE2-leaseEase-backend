package repositories

import (
	"LeaseEase/internal/models"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (u *userRepository) UpdateUser(user *models.User) error {
	return u.db.Save(user).Error
}

func (u *userRepository) GetUserByID(userID uint) (*models.User, error) {
	user := models.User{}
	err := u.db.Where("id = ?", userID).First(&user).Error
	return &user, err
}
