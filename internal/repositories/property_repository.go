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

func (r *propertyRepository) CreateProperty(property *models.MarketSlot) error {
	return r.db.Create(property).Error
}

func (r *propertyRepository) UpdateProperty(property *models.MarketSlot) error {
	return r.db.Save(property).Error
}

func (r *propertyRepository) DeleteProperty(id uint) error {
	return r.db.Delete(&models.MarketSlot{}, id).Error
}
