package repositories

import "LeaseEase/internal/models"

type PropertyRepository interface {
	CreateProperty(property *models.MarketSlot) error
	UpdateProperty(property *models.MarketSlot) error
	DeleteProperty(id uint) error
}
