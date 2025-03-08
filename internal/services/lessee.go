package services

import "LeaseEase/internal/dtos"

type LesseeService interface {
	CreateReservation(reservationDTO *dtos.CreateReservationDTO, lesseeId uint) error
	UpdateReservation(reservationDTO *dtos.UpdateReservationDTO, reservationID uint, lesseeID uint) error
	DeleteReservation(reservationID uint, lesseeID uint) error
	GetReservationsByLesseeID(lesseeID uint, limit int, offset int) ([]dtos.GetReservationDTO, error)
}
