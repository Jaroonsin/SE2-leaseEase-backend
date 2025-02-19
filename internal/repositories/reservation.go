package repositories

import "LeaseEase/internal/models"

type ReservationRepository interface {
	CreateReservation(reservation *models.Reservation) error
	UpdateReservation(reservation *models.Reservation) error
	DeleteReservation(reservationID uint) error
}