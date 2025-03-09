package repositories

import (
	"LeaseEase/internal/dtos"
	"LeaseEase/internal/models"
)

type LessorRepository interface {
	AcceptReservation(reservationID uint, lessorID uint) (*dtos.ApprovalReservationDTO, uint, error)
	DeclineReservation(reservationID uint, lessorID uint) (*dtos.ApprovalReservationDTO, uint, error)
	GetReservationByPropertiesID(propertyID uint, limit int, offset int) ([]models.Reservation, error)
}
