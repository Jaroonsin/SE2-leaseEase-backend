package services

import "LeaseEase/internal/dtos"

type LesseeService interface {
	CreateReservation(reservationDTO *dtos.CreateReservationDTO, lesseeId uint) error
	UpdateReservation(reservationDTO *dtos.UpdateReservationDTO, reservationID uint) error
	DeleteReservation(reservationID uint) error
	ApproveReservation(status string, reservationID uint) error
}
