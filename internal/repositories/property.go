package repositories

import "LeaseEase/internal/models"

type PropertyRepository interface {
	GetAllProperty() ([]models.MarketSlot, error)
	GetPropertyById(propertyID uint) (*models.MarketSlot, error)
	GetPaginatedProperties(page, limit int) ([]models.MarketSlot, int64, error)
}
