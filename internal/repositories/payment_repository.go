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

func (r *paymentRepository) UpdatePaymentStatus(id string, status string) error {
	return r.db.Model(&models.Payment{}).Where("id = ?", id).Update("status", status).Error
}
