package repositories

import (
	"LeaseEase/internal/models"

	"gorm.io/gorm"
)

type requestRepository struct {
	db *gorm.DB
}

func NewRequestRepository(db *gorm.DB) RequestRepository {
	return &requestRepository{
		db: db,
	}
}

func (r *requestRepository) CreateRequest(request *models.Request) error {
	return r.db.Create(request).Error
}

func (r *requestRepository) UpdateRequest(request *models.Request) error {
	result := r.db.Model(&request).Updates(request)

	if result.Error != nil {
		return result.Error 
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound 
	}

	return nil
}

func (r *requestRepository) DeleteRequest(requestID uint) error {
	result := r.db.Delete(&models.Request{}, requestID)
	if result.Error != nil {
		return result.Error 
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound 
	}

	return nil
}