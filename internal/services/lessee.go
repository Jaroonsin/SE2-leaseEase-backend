package services

import "LeaseEase/internal/dtos"

type LesseeService interface {
	CreateReservation(reservationDTO *dtos.CreateReservationDTO, lesseeID uint) (dtos.ReservationResponseDTO, error)
	UpdateReservation(reservationDTO *dtos.UpdateReservationDTO, reservationID uint, lesseeID uint) (dtos.ReservationResponseDTO, error)
	DeleteReservation(reservationID uint, lesseeID uint) (dtos.ReservationResponseDTO, error)
	GetReservationsByLesseeID(lesseeID uint, limit int, offset int) ([]dtos.GetReservationDTO, error)
}
