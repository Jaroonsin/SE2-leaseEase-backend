package repositories

import "LeaseEase/internal/models"

type PropertyRepository interface {
	CreateProperty(property *models.Property) error
	UpdateProperty(property *models.Property) error
	DeleteProperty(id uint) error
	GetAllProperty(lessorID uint) ([]models.Property, error)
	GetPaginatedProperty(lessorID uint, limit, offset int) ([]models.Property, error)
	GetPropertyById(propertyID uint) (*models.Property, error)
	SearchProperty(query map[string]string) ([]models.Property, uint, error)
	AutoComplete(query string) ([]string, error)
	CountPropertiesByLessor(lessorID uint, totalRecords *int64) error
	GetPropertyReviewsData(properties []models.Property) ([]float64, []int, [][]uint, error)
	GetPropertyReviewDataByID(propertyID uint) (float64, int, []uint, error)
}
