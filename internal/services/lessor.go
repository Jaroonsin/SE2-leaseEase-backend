package services

import "LeaseEase/internal/dtos"

type LessorService interface {
	AcceptReservation(reservationID uint, req *dtos.ApprovalReservationDTO, lessorID uint) (*dtos.ReservationResponseDTO, error)
	DeclineReservation(reservationID uint, req *dtos.ApprovalReservationDTO, lessorID uint) (*dtos.ReservationResponseDTO, error)
	GetReservationsByPropertyID(propertyID uint, page int, pageSize int) ([]dtos.GetReservationDTO, error)
}
