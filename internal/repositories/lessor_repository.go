package repositories

import (
	"LeaseEase/internal/dtos"
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

func (r *lessorRepository) AcceptReservation(reservationID uint, lessorID uint) (*dtos.ApprovalReservationDTO, uint, error) {

	var reservation models.Reservation
	result := r.db.Preload("Property").Preload("Lessee").First(&reservation, reservationID)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, 0, gorm.ErrRecordNotFound
	} else if result.Error != nil {
		return nil, 0, result.Error
	}

	if reservation.Property.LessorID != lessorID {
		return nil, 0, gorm.ErrRecordNotFound
	}

	if reservation.Status != "pending" {
		return nil, 0, errors.New("reservation can only be accepted if it is pending")
	}

	reservation.Status = "waiting"
	if err := r.db.Save(&reservation).Error; err != nil {
		return nil, 0, err
	}

	approvalDTO := &dtos.ApprovalReservationDTO{
		LesseeEmail:  reservation.Lessee.Email,
		PropertyName: reservation.Property.Name,
	}
	return approvalDTO, reservation.ID, nil
}

func (r *lessorRepository) DeclineReservation(reservationID uint, lessorID uint) (*dtos.ApprovalReservationDTO, uint, error) {

	var reservation models.Reservation
	result := r.db.Preload("Property").Preload("Lessee").First(&reservation, reservationID)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, 0, gorm.ErrRecordNotFound
	} else if result.Error != nil {
		return nil, 0, result.Error
	}

	if reservation.Property.LessorID != lessorID {
		return nil, 0, gorm.ErrRecordNotFound
	}

	if reservation.Status != "pending" {
		return nil, 0, errors.New("reservation can only be declined if it is pending")
	}

	reservation.Status = "cancel"
	if err := r.db.Save(&reservation).Error; err != nil {
		return nil, 0, err
	}

	approvalDTO := &dtos.ApprovalReservationDTO{
		LesseeEmail:  reservation.Lessee.Email,
		PropertyName: reservation.Property.Name,
	}
	return approvalDTO, reservation.ID, nil
}

func (r *lessorRepository) GetReservationByPropertiesID(propertyID uint, limit int, offset int) ([]models.Reservation, error) {
	var reservations []models.Reservation
	result := r.db.Preload("Property").Preload("Lessee").Where("interested_property = ?", propertyID).Offset(offset).Limit(limit).Find(&reservations)
	if result.Error != nil {
		return nil, result.Error
	}
	return reservations, nil
}
