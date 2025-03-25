package repositories

import (
	"LeaseEase/internal/models"

	"gorm.io/gorm"
)

type chatRepository struct {
	db *gorm.DB
}

func NewChatRepository(db *gorm.DB) ChatRepository {
	return &chatRepository{
		db: db,
	}
}

// Store a new message
func (r *chatRepository) StoreMessage(msg models.Message) error {
	err := r.db.Create(&msg).Error
	return err
}

// GetUndeliveredMessages fetches undelivered messages between a specific sender and receiver
func (r *chatRepository) GetUndeliveredMessages(userID string, receiverID string) ([]models.Message, error) {
	var messages []models.Message

	err := r.db.Where("sender_id = ? AND receiver_id = ? AND delivered = ?", userID, receiverID, false).
		Find(&messages).Error

	if err != nil {
		return nil, err
	}

	return messages, nil
}

// MarkMessagesAsDelivered updates message status to delivered for specific sender and receiver
func (r *chatRepository) MarkMessagesAsDelivered(userID string, receiverID string) error {
	return r.db.Model(&models.Message{}).
		Where("sender_id = ? AND receiver_id = ? AND delivered = ?", userID, receiverID, false).
		Update("delivered", true).Error
}

// Get message history between two users
func (r *chatRepository) GetHistoryMessages(userID string, receiverID string) ([]models.Message, error) {
	var messages []models.Message

	err := r.db.Where("(sender_id = ? AND receiver_id = ?) OR (sender_id = ? AND receiver_id = ?)",
		userID, receiverID, receiverID, userID).
		Order("timestamp ASC"). // Ensures messages are sorted in chronological order
		Find(&messages).Error

	if err != nil {
		return nil, err
	}

	return messages, nil
}
