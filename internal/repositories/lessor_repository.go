package repositories

import (
	"LeaseEase/internal/models"

	"gorm.io/gorm"
)

type lessorRepository struct {
	db *gorm.DB
}

func NewLessorRepository(db *gorm.DB) LessorRepository {
	return &lessorRepository{
		db: db,
	}
}

func (r *lessorRepository) AcceptReservation(reservationID uint) error {

	var reservation models.Reservation
	result := r.db.First(&reservation, reservationID)
	if result.Error != nil {
		return result.Error
	}

	reservation.Status = "waiting"
	//add some logics
	if err := r.db.Save(&reservation).Error; err != nil {
		return err
	}

	return nil
}

func (r *lessorRepository) DeclineReservation(reservationID uint) error {

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
