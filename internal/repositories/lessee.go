package repositories

import "LeaseEase/internal/models"

type LesseeRepository interface {
	CreateReservation(reservation *models.Reservation) error
	UpdateReservation(reservation *models.Reservation, lesseeID uint) error
	DeleteReservation(reservationID uint, lesseeID uint) error
	GetReservationByLesseeID(lesseeID uint, limit int, offset int) ([]models.Reservation, error)
}
