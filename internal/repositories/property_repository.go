package repositories

import (
	"LeaseEase/internal/models"

	"gorm.io/gorm"
)

type propertyRepository struct {
	db *gorm.DB
}

func NewPropertyRepository(db *gorm.DB) PropertyRepository {
	return &propertyRepository{
		db: db,
	}
}

func (r *propertyRepository) CreateProperty(property *models.Property) error {
	return r.db.Create(property).Error
}

func (r *propertyRepository) UpdateProperty(property *models.Property) error {
	return r.db.Model(&property).Updates(*property).Error
}

func (r *propertyRepository) DeleteProperty(id uint) error {
	return r.db.Delete(&models.Property{}, id).Error
}

func (r *propertyRepository) GetAllProperty(limit, offset int) ([]models.Property, error) {
	var properties []models.Property
	err := r.db.Limit(limit).Offset(offset).Find(&properties).Error
	if err != nil {
		return nil, err
	}
	return properties, nil
}
