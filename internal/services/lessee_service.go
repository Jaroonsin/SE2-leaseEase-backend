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

func (r *lesseeService) CreateReservation(reservationDTO *dtos.CreateReservationDTO, lesseeID uint) (*dtos.ReservationResponseDTO, error) {
	logger := r.logger.Named("CreateReservation")
	reservation := &models.Reservation{
		LesseeID:           lesseeID,
		Purpose:            reservationDTO.Purpose,
		ProposedMessage:    reservationDTO.ProposedMessage,
		Status:             "pending",
		Question:           reservationDTO.Question,
		InterestedProperty: reservationDTO.InterestedProperty,
	}

	id, err := r.lesseeRepo.CreateReservation(reservation)
	if err != nil {
		logger.Error("Failed to create reservation", zap.Error(err))
		return nil, err
	}
	logger.Info("Reservation created successfully", zap.Uint("reservationID", id))

	CreateReservationResponseDTO := &dtos.ReservationResponseDTO{
		ID: id,
	}
	return CreateReservationResponseDTO, nil
}

func (r *lesseeService) UpdateReservation(reservationDTO *dtos.UpdateReservationDTO, reservationID uint, lesseeID uint) (*dtos.ReservationResponseDTO, error) {
	logger := r.logger.Named("UpdateReservation")
	reservation := &models.Reservation{
		ID:              reservationID,
		Purpose:         reservationDTO.Purpose,
		ProposedMessage: reservationDTO.ProposedMessage,
		Question:        reservationDTO.Question,
	}
	id, err := r.lesseeRepo.UpdateReservation(reservation, lesseeID)
	if err != nil {
		logger.Error("Failed to update reservation", zap.Error(err))
		return nil, err
	}
	logger.Info("Reservation updated successfully", zap.Uint("reservationID", id))

	updateReservationResponseDTO := &dtos.ReservationResponseDTO{
		ID: id,
	}

	return updateReservationResponseDTO, nil
}

func (r *lesseeService) DeleteReservation(reservationID uint, lesseeID uint) (*dtos.ReservationResponseDTO, error) {
	logger := r.logger.Named("DeleteReservation")
	id, err := r.lesseeRepo.DeleteReservation(reservationID, lesseeID)
	if err != nil {
		logger.Error("Failed to delete reservation", zap.Error(err))
		return nil, err
	}
	logger.Info("Reservation deleted successfully", zap.Uint("reservationID", id))

	deleteReservationResponseDTO := &dtos.ReservationResponseDTO{
		ID: id,
	}
	return deleteReservationResponseDTO, nil
}

func (r *lesseeService) GetReservationsByLesseeID(lesseeID uint, limit int, offset int) ([]dtos.GetReservationDTO, error) {
	logger := r.logger.Named("GetReservationsByLesseeID")
	reservations, err := r.lesseeRepo.GetReservationByLesseeID(lesseeID, limit, offset)
	if err != nil {
		logger.Error("Failed to get reservations by lessee ID", zap.Error(err))
		return nil, err
	}

	var GetReservationDTOs []dtos.GetReservationDTO
	for _, reservation := range reservations {
		reservationDTO := dtos.GetReservationDTO{
			ID:              reservation.ID,
			LesseeID:        reservation.LesseeID,
			Purpose:         reservation.Purpose,
			ProposedMessage: reservation.ProposedMessage,
			Status:          reservation.Status,
			Question:        reservation.Question,
			PropertyID:      reservation.InterestedProperty,
			PropertyName:    reservation.Property.Name,
			LastModified:    reservation.UpdatedAt.Format("2006-01-02T15:04:05Z"),
		}
		GetReservationDTOs = append(GetReservationDTOs, reservationDTO)
	}

	logger.Info("Reservations fetched successfully", zap.Int("count", len(GetReservationDTOs)))
	return GetReservationDTOs, nil
}
