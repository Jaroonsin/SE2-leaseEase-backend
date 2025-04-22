package repositories

import "LeaseEase/internal/models"

type ChatRepository interface {
	// Chatroom Management
	CreateChatroom(name string, isPrivate bool) (*models.Chatroom, error)
	GetChatroomByID(chatroomID uint) (*models.Chatroom, error)
	GetChatroomsByUser(userID uint, limit int, offset int) ([]models.Chatroom, error)
	UpdateLastMessageID(chatroomID uint, messageID uint) error
	UpdateLastMessageContent(chatroomID uint, content string) error

	// Chatroom Members Management
	AddMemberToChatroom(chatroomID uint, userID uint) error
	RemoveMemberFromChatroom(chatroomID uint, userID uint) error
	GetMembersOfChatroom(chatroomID uint) ([]models.User, error)
	GetLastReadMessageID(chatroomID uint, userID uint) (uint, error)
	UpdateLastReadMessageID(chatroomID uint, userID uint, messageID uint) error

	// Message Management
	CreateMessage(chatroomID uint, senderID uint, content string) (*models.Message, error)
	GetMessagesByChatroom(chatroomID uint, limit, offset int) ([]models.Message, error)
	GetLastMessageInChatroom(chatroomID uint) (*models.Message, error)
	GetMessageByID(messageID uint) (*models.Message, error)

	// Message Read Status
	MarkMessageAsRead(messageID uint, userID uint) error
	GetMessageReadStatus(messageID uint, userID uint) (*models.MessageRead, error)
	GetUnreadMessagesForUserInChatroom(userID uint, chatroomID uint) ([]models.Message, error)

	// Utility Methods
	IsUserInChatroom(chatroomID uint, userID uint) (bool, error)
	GetUnreadMessageCount(userID uint, chatroomID uint) (int, error)
	GetHistoryMessages(chatroomID string, limit int, offset int) ([]models.Message, error)
}
