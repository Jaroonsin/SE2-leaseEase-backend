package services

import (
	"LeaseEase/internal/repositories"

	"go.uber.org/zap"
)

type lessorService struct {
	lessorRepo repositories.LessorRepository
	logger     *zap.Logger
}

func NewLessorService(repo repositories.LessorRepository, logger *zap.Logger) LessorService {
	return &lessorService{
		lessorRepo: repo,
		logger:     logger,
	}
}

func (s *lessorService) AcceptReservation(reservationID uint) error {
	err := s.lessorRepo.AcceptReservation(reservationID)
	if err != nil {
		s.logger.Error("failed to accept reservation", zap.Uint("reservationID", reservationID), zap.Error(err))
		return err
	}
	return nil
}

func (s *lessorService) DeclineReservation(reservationID uint) error {
	err := s.lessorRepo.DeclineReservation(reservationID)
	if err != nil {
		s.logger.Error("failed to decline reservation", zap.Uint("reservationID", reservationID), zap.Error(err))
		return err
	}
	return nil
}
