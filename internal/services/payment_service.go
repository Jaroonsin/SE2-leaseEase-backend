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

func (s *paymentService) ProcessPayment(userID uint, amount int64, currency, token string) (*models.Payment, error) {
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
		return nil, err
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
		return nil, err
	}

	logger.Info("Payment record saved successfully",
		zap.String("chargeID", charge.ID),
	)
	return payment, nil
}
