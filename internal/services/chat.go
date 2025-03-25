package services

import (
	"LeaseEase/internal/dtos"
)

type ChatService interface {
	ProcessMessage(msg dtos.SendMessageRequest, isReceiverOnline bool) error
	DeliverOfflineMessages(userID string, receiverID string) ([]dtos.MessageResponse, error)
	DeliverHistoryMessages(userID string, receiverID string) ([]dtos.MessageResponse, error)
}
