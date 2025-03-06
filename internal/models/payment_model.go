package models

import (
	"time"
)

// Payment struct to store transaction details
type Payment struct {
	ID        string    `gorm:"primaryKey"`
	UserID    uint      `gorm:"not null;index"`              // Foreign key
	User      User      `gorm:"constraint:OnDelete:CASCADE"` // Establish relationship with User
	Amount    int64     // Stored in satangs (Omise)
	Currency  string    // Example: "THB", "USD"
	Status    string    // "pending", "successful", "failed"
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
