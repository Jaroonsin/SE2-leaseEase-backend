package services

import (
	"LeaseEase/internal/dtos"
	"LeaseEase/internal/models"
	"LeaseEase/internal/repositories"

	"go.uber.org/zap"
)

type lesseeService struct {
	lesseeRepo repositories.LesseeRepository
	logger     *zap.Logger
}

func NewLesseeService(lesseeRepo repositories.LesseeRepository, logger *zap.Logger) LesseeService {
	return &lesseeService{
		lesseeRepo: lesseeRepo,
		logger:     logger,
	}
}

func (r *lesseeService) CreateReservation(reservationDTO *dtos.CreateReservationDTO, lesseeId uint) error {
	reservation := &models.Reservation{
		LesseeID:           lesseeId,
		Purpose:            reservationDTO.Purpose,
		ProposedMessage:    reservationDTO.ProposedMessage,
		Status:             "pending",
		Question:           reservationDTO.Question,
		InterestedProperty: reservationDTO.InterestedProperty,
	}

	return r.lesseeRepo.CreateReservation(reservation)
}

func (r *lesseeService) UpdateReservation(reservationDTO *dtos.UpdateReservationDTO, reservationID uint, lesseeID uint) error {
	reservation := &models.Reservation{
		ID:              reservationID,
		Purpose:         reservationDTO.Purpose,
		ProposedMessage: reservationDTO.ProposedMessage,
		Question:        reservationDTO.Question,
	}
	return r.lesseeRepo.UpdateReservation(reservation, lesseeID)
}

func (r *lesseeService) DeleteReservation(reservationID uint, lesseeID uint) error {
	return r.lesseeRepo.DeleteReservation(reservationID, lesseeID)
}
