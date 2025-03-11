package repositories

import (
	"LeaseEase/internal/models"

	"gorm.io/gorm"
)

type paymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) PaymentRepository {
	return &paymentRepository{db: db}
}

func (r *paymentRepository) CreatePayment(payment *models.Payment) error {
	return r.db.Create(payment).Error
}

func (r *paymentRepository) UpdatePaymentStatus(id uint, status string) error {
	err := r.db.Model(&models.Reservation{}).Where("id = ?", id).Update("status", status).Error
	if err != nil {
		return err
	}
	return nil
}
