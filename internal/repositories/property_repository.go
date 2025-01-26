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

	err := r.db.Find(&properties).Error
	if err != nil {
		return nil, err
	}

	return properties, nil
}

func (r *propertyRepository) GetPropertyById(propertyID uint) (*models.MarketSlot, error) {
	var property models.MarketSlot

	err := r.db.First(&property, propertyID).Error
	if err != nil {
		return nil, err
	}

	return &property, nil
}

func (r *propertyRepository) GetPaginatedProperties(page, limit int) ([]models.MarketSlot, int64, error) {
	var properties []models.MarketSlot
	var totalRecords int64

	err := r.db.Model(&models.MarketSlot{}).Count(&totalRecords).Error
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	err = r.db.Offset(offset).Limit(limit).Find(&properties).Error
	if err != nil {
		return nil, 0, err
	}

	return properties, totalRecords, nil
}
