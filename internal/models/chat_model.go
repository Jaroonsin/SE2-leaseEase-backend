package models

import "time"

type Message struct {
	ID         uint      `gorm:"primaryKey"`
	SenderID   uint      `gorm:"not null"`
	ReceiverID uint      `gorm:"not null"`
	Content    string    `gorm:"type:text;not null"`
	Timestamp  time.Time `gorm:"autoCreateTime"`
	Delivered  bool      `gorm:"default:false"`

	Sender   User `gorm:"foreignKey:SenderID;references:ID"`
	Receiver User `gorm:"foreignKey:ReceiverID;references:ID"`
}
