package repositories

import (
	"LeaseEase/internal/models"
)

type LessorRepository interface {
	AcceptReservation(reservationID uint, lessorID uint) (string, uint, error)
	DeclineReservation(reservationID uint, lessorID uint) (string, uint, error)
	GetReservationByPropertiesID(propertyID uint, limit int, offset int) ([]models.Reservation, error)
}
