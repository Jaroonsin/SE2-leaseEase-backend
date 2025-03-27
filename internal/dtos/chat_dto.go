package dtos

import "time"

type SendMessageRequest struct {
	SenderID   uint   `json:"sender_id" validate:"required"`
	ReceiverID uint   `json:"receiver_id" validate:"required"`
	Content    string `json:"content" validate:"required"`
}

type MessageResponse struct {
	ID         uint      `json:"id"`
	SenderID   uint      `json:"sender_id"`
	ReceiverID uint      `json:"receiver_id"`
	Content    string    `json:"content"`
	Timestamp  time.Time `json:"timestamp"`
	Delivered  bool      `json:"delivered"`
}
