package services

import (
	"LeaseEase/internal/dtos"
	"LeaseEase/internal/models"
	"LeaseEase/internal/repositories"

	"go.uber.org/zap"
)

type reservationService struct {
	reservationRepo repositories.ReservationRepository
	logger          *zap.Logger
}

func NewReservationService(reservationRepo repositories.ReservationRepository, logger *zap.Logger) ReservationService {
	return &reservationService{
		reservationRepo: reservationRepo,
		logger:          logger,
	}
}

func (r *reservationService) CreateReservation(reservationDTO *dtos.CreateReservation, lesseeId uint) error {
	reservation := &models.Reservation{
		LesseeID:           lesseeId,
		Purpose:            reservationDTO.Purpose,
		ProposedMessage:    reservationDTO.ProposedMessage,
		Status:             "pending",
		Question:           reservationDTO.Question,
		InterestedProperty: reservationDTO.InterestedProperty,
	}

	return r.reservationRepo.CreateReservation(reservation)
}

func (r *reservationService) UpdateReservation(reservationDTO *dtos.UpdateReservation, reservationID uint) error {
	reservation := &models.Reservation{
		ID:              reservationID,
		Purpose:         reservationDTO.Purpose,
		ProposedMessage: reservationDTO.ProposedMessage,
		Question:        reservationDTO.Question,
	}
	return r.reservationRepo.UpdateReservation(reservation)
}

func (r *reservationService) DeleteReservation(reservationID uint) error {
	return r.reservationRepo.DeleteReservation(reservationID)
}

func (r *reservationService) ApproveReservation(status string, reservationID uint) error {
	reservation := &models.Reservation{
		ID:     reservationID,
		Status: status,
	}
	return r.reservationRepo.UpdateReservation(reservation)
}
