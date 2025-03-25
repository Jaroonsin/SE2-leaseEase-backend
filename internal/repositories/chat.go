package repositories

import "LeaseEase/internal/models"

type ChatRepository interface {
	StoreMessage(msg models.Message) error
	GetUndeliveredMessages(userID string, receiverID string) ([]models.Message, error)
	GetHistoryMessages(userID string, receiverID string) ([]models.Message, error)
	MarkMessagesAsDelivered(userID string, receiverID string) error
}
