package repositories

import "LeaseEase/internal/models"

type PropertyRepository interface {
	CreateProperty(property *models.Property) error
	UpdateProperty(property *models.Property) error
	DeleteProperty(id uint) error
	GetAllProperty(limit, offset int) ([]models.Property, error)
}
