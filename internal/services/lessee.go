package services

import "LeaseEase/internal/dtos"

type LesseeService interface {
	CreateReservation(reservationDTO *dtos.CreateReservation, lesseeId uint) error
	UpdateReservation(reservationDTO *dtos.UpdateReservation, reservationID uint) error
	DeleteReservation(reservationID uint) error
	ApproveReservation(status string, reservationID uint) error
}
