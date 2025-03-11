package repositories

import "LeaseEase/internal/models"

type PaymentRepository interface {
	CreatePayment(payment *models.Payment) error
	UpdatePaymentStatus(id uint, status string) error
}
