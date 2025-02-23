package repositories

import (
	"LeaseEase/internal/models"
	"errors"

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

func (r *lessorRepository) AcceptReservation(reservationID uint, lessorID uint) error {

	var reservation models.Reservation
	result := r.db.Preload("Property").First(&reservation, reservationID)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return gorm.ErrRecordNotFound
	} else if result.Error != nil {
		return result.Error
	}

	if reservation.Property.LessorID != lessorID {
		return gorm.ErrRecordNotFound
	}

	if reservation.Status != "pending" {
		return errors.New("reservation can only be accepted if it is pending")
	}

	reservation.Status = "waiting"
	if err := r.db.Save(&reservation).Error; err != nil {
		return err
	}

	return nil
}

func (r *lessorRepository) DeclineReservation(reservationID uint, lessorID uint) error {

	var reservation models.Reservation
	result := r.db.Preload("Property").First(&reservation, reservationID)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return gorm.ErrRecordNotFound
	} else if result.Error != nil {
		return result.Error
	}

	if reservation.Property.LessorID != lessorID {
		return gorm.ErrRecordNotFound
	}

	if reservation.Status != "pending" {
		return errors.New("reservation can only be declined if it is pending")
	}

	reservation.Status = "cancel"
	if err := r.db.Save(&reservation).Error; err != nil {
		return err
	}

	return nil
}
