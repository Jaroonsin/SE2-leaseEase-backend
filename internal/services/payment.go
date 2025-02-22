package services

import "LeaseEase/internal/models"

type PaymentService interface {
	ProcessPayment(userID uint, amount int64, currency, token string) (*models.Payment, error)
}
