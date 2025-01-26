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

func (r *propertyRepository) GetAllProperty() ([]models.MarketSlot, error) {
	var properties []models.MarketSlot
	///
	return properties, nil
}

func (r *propertyRepository) GetPropertyById(propertyID uint) (*models.MarketSlot, error) {
	var property models.MarketSlot
	///
	return &property, nil
}

//func pagination
