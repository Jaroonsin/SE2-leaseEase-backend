package repositories

import "LeaseEase/internal/models"

type PaymentRepository interface {
	CreatePayment(payment *models.Payment) error
	GetAmountByReservationID(reservationID uint) (float64, error)
	UpdatePaymentStatus(id uint, status string) error
}
