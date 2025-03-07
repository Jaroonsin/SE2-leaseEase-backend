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

func (r *userRepository) UpdateUser(user *models.User) error {
	result := r.db.Model(&user).Updates(*user)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (r *userRepository) GetUserByID(userID uint) (*models.User, error) {
	user := models.User{}
	err := r.db.Where("id = ?", userID).First(&user).Error
	return &user, err
}
