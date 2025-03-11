package services

import (
	"LeaseEase/internal/models"
	"LeaseEase/internal/repositories"
	"LeaseEase/utils"

	"github.com/omise/omise-go"
	"github.com/omise/omise-go/operations"
	"go.uber.org/zap"
)

type paymentService struct {
	logger      *zap.Logger
	paymentRepo repositories.PaymentRepository
}

func NewPaymentService(paymentRepo repositories.PaymentRepository, logger *zap.Logger) PaymentService {
	return &paymentService{
		logger:      logger,
		paymentRepo: paymentRepo,
	}
}

func (s *paymentService) ProcessPayment(userID uint, amount int64, currency, token string, reservationID uint) error {
	logger := s.logger.Named("ProcessPayment")
	charge := &omise.Charge{}
	client, err := utils.NewOmiseClient()
	if err != nil {
		logger.Error("Fail to create omise client", zap.Error(err))
	}

	err = client.Do(charge, &operations.CreateCharge{
		Amount:   amount,
		Currency: currency,
		Card:     token,
	})

	if err != nil {
		logger.Error("Payment processing failed",
			zap.Uint("userID", userID),
			zap.Error(err),
		)
		return err
	}

	logger.Info("Payment successful",
		zap.String("chargeID", charge.ID),
		zap.String("status", string(charge.Status)),
	)

	payment := &models.Payment{
		ID:       charge.ID,
		UserID:   userID,
		Amount:   amount,
		Currency: currency,
		Status:   string(charge.Status),
	}

	if err := s.paymentRepo.CreatePayment(payment); err != nil {
		logger.Error("Failed to save payment record",
			zap.String("chargeID", charge.ID),
			zap.Error(err),
		)
		return err
	}

	logger.Info("Payment record saved successfully waiting for reservation update",
		zap.String("chargeID", charge.ID),
	)

	if err := s.paymentRepo.UpdatePaymentStatus(reservationID, "accept"); err != nil {
		logger.Error("Failed to update reservation",
			zap.String("chargeID", charge.ID),
			zap.Error(err),
		)
		return err
	}

	logger.Info("Reservation updated successfully",
		zap.Uint("reservationID", reservationID),
	)

	return nil
}
