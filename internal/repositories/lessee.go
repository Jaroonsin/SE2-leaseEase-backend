package repositories

import "LeaseEase/internal/models"

type LesseeRepository interface {
	CreateReservation(reservation *models.Reservation) (uint, error)
	UpdateReservation(reservation *models.Reservation, lesseeID uint) (uint, error)
	DeleteReservation(reservationID uint, lesseeID uint) (uint, error)
	GetReservationByLesseeID(lesseeID uint, limit int, offset int) ([]models.Reservation, error)
}
