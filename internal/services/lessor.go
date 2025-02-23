package services

import "LeaseEase/internal/dtos"

type LessorService interface {
	AcceptReservation(reservationID uint, req *dtos.AcceptReservationDTO) error
	DeclineReservation(reservationID uint) error
}
