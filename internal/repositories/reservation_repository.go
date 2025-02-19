package repositories

import (
	"LeaseEase/internal/models"

	"gorm.io/gorm"
)

type reservationRepository struct {
	db *gorm.DB
}

func NewReservationRepository(db *gorm.DB) ReservationRepository {
	return &reservationRepository{
		db: db,
	}
}

func (r *reservationRepository) CreateReservation(reservation *models.Reservation) error {
	return r.db.Create(reservation).Error
}

func (r *reservationRepository) UpdateReservation(reservation *models.Reservation) error {
	result := r.db.Model(&reservation).Updates(reservation)

	if result.Error != nil {
		return result.Error 
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound 
	}

	return nil
}

func (r *reservationRepository) DeleteReservation(reservationID uint) error {
	result := r.db.Delete(&models.Reservation{}, reservationID)
	if result.Error != nil {
		return result.Error 
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound 
	}

	return nil
}