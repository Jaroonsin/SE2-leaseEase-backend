package services

import (
	"LeaseEase/internal/dtos"
	"LeaseEase/internal/repositories"
	"LeaseEase/utils"

	"go.uber.org/zap"
)

type lessorService struct {
	lessorRepo repositories.LessorRepository
	logger     *zap.Logger
}

func NewLessorService(lessorRepo repositories.LessorRepository, logger *zap.Logger) LessorService {
	return &lessorService{
		lessorRepo: lessorRepo,
		logger:     logger,
	}
}

func (s *lessorService) AcceptReservation(reservationID uint, lessorID uint) (*dtos.ReservationResponseDTO, error) {
	req, ID, err := s.lessorRepo.AcceptReservation(reservationID, lessorID)
	if err != nil {
		s.logger.Error("failed to accept reservation", zap.Uint("reservationID", reservationID), zap.Error(err))
		return nil, err
	}
	err = utils.SendLessorAcceptanceEmail(req)
	if err != nil {
		s.logger.Error("failed to send acceptance email", zap.Uint("reservationID", reservationID), zap.Error(err))
		return nil, err
	}

	reservationResponse := &dtos.ReservationResponseDTO{
		ID: ID,
	}

	return reservationResponse, nil
}

func (s *lessorService) DeclineReservation(reservationID uint, lessorID uint) (*dtos.ReservationResponseDTO, error) {
	req, ID, err := s.lessorRepo.DeclineReservation(reservationID, lessorID)
	if err != nil {
		s.logger.Error("failed to decline reservation", zap.Uint("reservationID", reservationID), zap.Error(err))
		return nil, err
	}
	err = utils.SendLessorDeclineEmail(req.LesseeEmail, req.PropertyName)
	if err != nil {
		s.logger.Error("failed to send decline email", zap.Uint("reservationID", reservationID), zap.Error(err))
		return nil, err
	}

	reservationResponse := &dtos.ReservationResponseDTO{
		ID: ID,
	}

	return reservationResponse, nil
}

func (s *lessorService) GetReservationsByPropertyID(propertyID uint, limit int, offset int) ([]dtos.GetPropReservationDTO, error) {
	reservations, err := s.lessorRepo.GetReservationByPropertiesID(propertyID, limit, offset)
	if err != nil {
		s.logger.Error("failed to get reservations by property ID", zap.Uint("propertyID", propertyID), zap.Error(err))
		return nil, err
	}
	var GetPropReservationDTOs []dtos.GetPropReservationDTO
	for _, reservation := range reservations {
		GetPropReservationDTO := dtos.GetPropReservationDTO{
			ID:              reservation.ID,
			Purpose:         reservation.Purpose,
			ProposedMessage: reservation.ProposedMessage,
			Question:        reservation.Question,
			Status:          reservation.Status,
			PropertyID:      reservation.InterestedProperty,
			LesseeID:        reservation.LesseeID,
			LesseeName:      reservation.Lessee.Name,
			ImageURL:        reservation.Lessee.ImageURL,
			PropertyName:    reservation.Property.Name,
			LastModified:    reservation.UpdatedAt.Format("2006-01-02T15:04:05Z"),
		}
		GetPropReservationDTOs = append(GetPropReservationDTOs, GetPropReservationDTO)
	}

	return GetPropReservationDTOs, nil
}
