package repositories

import "LeaseEase/internal/models"

type PropertyRepository interface {
	CreateProperty(property *models.Property) error
	UpdateProperty(property *models.Property) error
	DeleteProperty(id uint) error
	GetAllProperty(lessorID uint) ([]models.Property, error)
	GetPaginatedProperty(lessorID uint, limit, offset int) ([]models.Property, error)
	GetPropertyById(propertyID uint) (*models.Property, error)
	SearchProperty(query map[string]string) ([]models.Property, error)
	AutoComplete(query string) ([]string, error)
}
