package services

import "LeaseEase/internal/dtos"

type LessorService interface {
	AcceptReservation(reservationID uint, req *dtos.ApprovalReservationDTO, lessorID uint) error
	DeclineReservation(reservationID uint, req *dtos.ApprovalReservationDTO, lessorID uint) error
}
