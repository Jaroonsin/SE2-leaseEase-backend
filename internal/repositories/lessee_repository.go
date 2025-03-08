package repositories

import (
	"LeaseEase/internal/models"

	"gorm.io/gorm"
)

type lesseeRepository struct {
	db *gorm.DB
}

func NewLesseeRepository(db *gorm.DB) LesseeRepository {
	return &lesseeRepository{
		db: db,
	}
}

func (r *lesseeRepository) CreateReservation(reservation *models.Reservation) error {
	return r.db.Create(reservation).Error
}

func (r *lesseeRepository) UpdateReservation(reservation *models.Reservation, lesseeID uint) error {
	var existingReservation models.Reservation
	result := r.db.First(&existingReservation, reservation.ID)
	if result.Error != nil {
		return result.Error
	}

	if existingReservation.LesseeID != lesseeID {
		return gorm.ErrRecordNotFound
	}

	result = r.db.Model(&existingReservation).Updates(reservation)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (r *lesseeRepository) DeleteReservation(reservationID uint, lesseeID uint) error {
	var reservation models.Reservation
	result := r.db.First(&reservation, reservationID)
	if result.Error != nil {
		return result.Error
	}

	if reservation.LesseeID != lesseeID {
		return gorm.ErrRecordNotFound
	}

	result = r.db.Delete(&reservation)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (r *lesseeRepository) GetReservationByLesseeID(lesseeID uint, limit int, offset int) ([]models.Reservation, error) {
	var reservations []models.Reservation
	result := r.db.Where("lessee_id = ?", lesseeID).Limit(limit).Offset(offset).Find(&reservations)
	if result.Error != nil {
		return nil, result.Error
	}
	return reservations, nil
}
