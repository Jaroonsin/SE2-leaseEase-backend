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

func (r *lesseeRepository) UpdateReservation(reservation *models.Reservation) error {
	result := r.db.Model(&reservation).Updates(reservation)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (r *lesseeRepository) DeleteReservation(reservationID uint) error {
	result := r.db.Delete(&models.Reservation{}, reservationID)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (r *lesseeRepository) AcceptReservation(reservationID uint) error {

	var reservation models.Reservation
	result := r.db.First(&reservation, reservationID)
	if result.Error != nil {
		return result.Error
	}

	reservation.Status = "waiting"
	//some logic
	if err := r.db.Save(&reservation).Error; err != nil {
		return err
	}

	return nil
}

func (r *lesseeRepository) DeclineReservation(reservationID uint) error {

	var reservation models.Reservation
	result := r.db.First(&reservation, reservationID)
	if result.Error != nil {
		return result.Error
	}

	reservation.Status = "cancel"
	if err := r.db.Save(&reservation).Error; err != nil {
		return err
	}

	return nil
}
