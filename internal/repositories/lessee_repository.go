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

func (r *lesseeRepository) CreateReservation(reservation *models.Reservation) (uint, error) {
	if err := r.db.Create(reservation).Error; err != nil {
		return 0, err
	}
	return reservation.ID, nil
}

func (r *lesseeRepository) UpdateReservation(reservation *models.Reservation, lesseeID uint) (uint, error) {
	var existingReservation models.Reservation
	result := r.db.First(&existingReservation, reservation.ID)
	if result.Error != nil {
		return 0, result.Error
	}

	if existingReservation.LesseeID != lesseeID {
		return 0, gorm.ErrRecordNotFound
	}

	result = r.db.Model(&existingReservation).Updates(reservation)
	if result.Error != nil {
		return 0, result.Error
	}

	if result.RowsAffected == 0 {
		return 0, gorm.ErrRecordNotFound
	}

	return existingReservation.ID, nil
}

func (r *lesseeRepository) DeleteReservation(reservationID uint, lesseeID uint) (uint, error) {
	var reservation models.Reservation
	result := r.db.First(&reservation, reservationID)
	if result.Error != nil {
		return 0, result.Error
	}

	if reservation.LesseeID != lesseeID {
		return 0, gorm.ErrRecordNotFound
	}

	result = r.db.Delete(&reservation)
	if result.Error != nil {
		return 0, result.Error
	}

	if result.RowsAffected == 0 {
		return 0, gorm.ErrRecordNotFound
	}

	return reservation.ID, nil
}

func (r *lesseeRepository) GetReservationByLesseeID(lesseeID uint, limit int, offset int) ([]models.Reservation, error) {
	var reservations []models.Reservation
	result := r.db.Preload("Property").Where("lessee_id = ?", lesseeID).Limit(limit).Offset(offset).Find(&reservations)
	if result.Error != nil {
		return nil, result.Error
	}
	return reservations, nil
}
