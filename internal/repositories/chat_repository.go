package repositories

import (
	"LeaseEase/internal/models"
	"time"

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

// Chatroom Management
func (r *chatRepository) CreateChatroom(name string, isPrivate bool) (*models.Chatroom, error) {
	chatroom := &models.Chatroom{
		Name:      name,
		IsPrivate: isPrivate,
	}

	err := r.db.Create(chatroom).Error
	if err != nil {
		return nil, err
	}

	return chatroom, nil
}

func (r *chatRepository) GetChatroomByID(chatroomID uint) (*models.Chatroom, error) {
	var chatroom models.Chatroom
	err := r.db.First(&chatroom, chatroomID).Error
	if err != nil {
		return nil, err
	}
	return &chatroom, nil
}

func (r *chatRepository) GetChatroomsByUser(userID uint, limit int, offset int) ([]models.Chatroom, error) {
	var chatrooms []models.Chatroom
	err := r.db.Model(&models.Chatroom{}).
		Joins("JOIN chatroom_members ON chatrooms.chatroom_id = chatroom_members.chatroom_id").
		Where("chatroom_members.user_id = ?", userID).
		Order("chatrooms.created_at DESC").
		Limit(limit).   // Apply limit for pagination
		Offset(offset). // Apply offset for pagination
		Find(&chatrooms).Error

	if err != nil {
		return nil, err
	}
	return chatrooms, nil
}

func (r *chatRepository) UpdateLastMessage(chatroomID uint, messageID uint) error {
	err := r.db.Model(&models.Chatroom{}).
		Where("chatroom_id = ?", chatroomID).
		Update("last_message_id", messageID).Error

	if err != nil {
		return err
	}
	return nil
}

func (r *chatRepository) AddMemberToChatroom(chatroomID uint, userID uint) error {
	err := r.db.Create(&models.ChatroomMember{
		ChatroomID: chatroomID,
		UserID:     userID,
		JoinedAt:   time.Now(),
	}).Error

	if err != nil {
		return err
	}
	return nil
}

func (r *chatRepository) RemoveMemberFromChatroom(chatroomID uint, userID uint) error {
	err := r.db.Where("chatroom_id = ? AND user_id = ?", chatroomID, userID).
		Delete(&models.ChatroomMember{}).Error

	if err != nil {
		return err
	}
	return nil
}

func (r *chatRepository) GetMembersOfChatroom(chatroomID uint) ([]models.User, error) {
	var members []models.User
	err := r.db.Model(&models.ChatroomMember{}).
		Joins("JOIN users ON chatroom_members.user_id = users.id").
		Where("chatroom_members.chatroom_id = ?", chatroomID).
		Select("users.*"). // Select all fields from the 'users' table
		Find(&members).Error

	if err != nil {
		return nil, err
	}
	return members, nil
}

// Message Management
func (r *chatRepository) CreateMessage(chatroomID uint, senderID uint, content string) (*models.Message, error) {
	message := &models.Message{
		ChatroomID: chatroomID,
		SenderID:   senderID,
		Content:    content,
	}

	err := r.db.Create(message).Error
	if err != nil {
		return nil, err
	}
	return message, nil
}

func (r *chatRepository) GetMessagesByChatroom(chatroomID uint, limit, offset int) ([]models.Message, error) {
	var messages []models.Message
	err := r.db.Where("chatroom_id = ?", chatroomID).
		Limit(limit).Offset(offset).
		Order("sent_at DESC").
		Find(&messages).Error

	if err != nil {
		return nil, err
	}
	return messages, nil
}

func (r *chatRepository) GetLastMessageInChatroom(chatroomID uint) (*models.Message, error) {
	var message models.Message
	err := r.db.Where("chatroom_id = ?", chatroomID).
		Order("sent_at DESC").
		First(&message).Error

	if err != nil {
		return nil, err
	}
	return &message, nil
}

func (r *chatRepository) GetMessageByID(messageID uint) (*models.Message, error) {
	err := r.db.Where("id = ?", messageID).
		First(&models.Message{}).Error

	if err != nil {
		return nil, err
	}
	return &models.Message{}, nil
}

func (r *chatRepository) MarkMessageAsRead(messageID uint, userID uint) error {
	err := r.db.Create(&models.MessageRead{
		MessageID: messageID,
		UserID:    userID,
		ReadAt:    time.Now(),
	}).Error

	if err != nil {
		return err
	}
	return nil
}

func (r *chatRepository) GetMessageReadStatus(messageID uint, userID uint) (*models.MessageRead, error) {
	var messageRead models.MessageRead
	err := r.db.Where("message_id = ? AND user_id = ?", messageID, userID).
		First(&messageRead).Error

	if err != nil {
		return nil, err
	}
	return &messageRead, nil
}

func (r *chatRepository) GetUnreadMessagesForUserInChatroom(userID uint, chatroomID uint) ([]models.Message, error) {
	var messages []models.Message
	err := r.db.Model(&models.Message{}).
		Joins("LEFT JOIN message_reads ON messages.id = message_reads.message_id AND message_reads.user_id = ?", userID).
		Where("messages.chatroom_id = ? AND message_reads.user_id IS NULL", chatroomID).
		Find(&messages).Error

	if err != nil {
		return nil, err
	}
	return messages, nil
}

// utility Methods
func (r *chatRepository) IsUserInChatroom(chatroomID uint, userID uint) (bool, error) {
	var count int64
	err := r.db.Model(&models.ChatroomMember{}).
		Where("chatroom_id = ? AND user_id = ?", chatroomID, userID).
		Count(&count).Error

	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *chatRepository) GetUnreadMessageCount(userID uint, chatroomID uint) (int, error) {
	var count int64
	err := r.db.Model(&models.Message{}).
		Joins("LEFT JOIN message_reads ON messages.id = message_reads.message_id AND message_reads.user_id = ?", userID).
		Where("messages.chatroom_id = ? AND message_reads.user_id IS NULL", chatroomID).
		Count(&count).Error

	if err != nil {
		return 0, err
	}
	return int(count), nil
}

func (r *chatRepository) GetHistoryMessages(chatroomID string, limit int, offset int) ([]models.Message, error) {
	// Example: Retrieve chat history from the database (messages table) by chatroomID
	// You can use GORM or any other database query approach

	var messages []models.Message
	if err := r.db.Where("chatroom_id = ?", chatroomID).
		Order("Timestamp ASC").
		Limit(limit).
		Offset(offset).
		Find(&messages).Error; err != nil {
		return nil, err
	}

	return messages, nil
}
