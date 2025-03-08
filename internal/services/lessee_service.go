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

func (r *lesseeService) CreateReservation(reservationDTO *dtos.CreateReservationDTO, lesseeID uint) error {
	reservation := &models.Reservation{
		LesseeID:           lesseeID,
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

func (r *lesseeService) GetReservationsByLesseeID(lesseeID uint, limit int, offset int) ([]dtos.GetReservationDTO, error) {
	reservations, err := r.lesseeRepo.GetReservationByLesseeID(lesseeID, limit, offset)
	if err != nil {
		r.logger.Error("failed to get reservations by lessee ID", zap.Error(err))
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

	return GetReservationDTOs, nil
}
