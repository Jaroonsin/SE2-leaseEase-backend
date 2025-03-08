package repositories

import (
	"LeaseEase/internal/models"
)

type LessorRepository interface {
	AcceptReservation(reservationID uint, lessorID uint) error
	DeclineReservation(reservationID uint, lessorID uint) error
	GetReservationByPropertiesID(propertyID uint, page int, pageSize int) ([]models.Reservation, error)
}
