package services

import "LeaseEase/internal/dtos"

type LessorService interface {
	AcceptReservation(reservationID uint, lessorID uint) (*dtos.ReservationResponseDTO, error)
	DeclineReservation(reservationID uint, lessorID uint) (*dtos.ReservationResponseDTO, error)

	GetReservationsByPropertyID(propertyID uint, page int, pageSize int) ([]dtos.GetPropReservationDTO, error)
}
