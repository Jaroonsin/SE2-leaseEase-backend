package services

import (
	"LeaseEase/internal/dtos"
)

// type ChatService interface {
// 	ProcessMessage(msg dtos.SendMessageRequest, isReceiverOnline bool) error
// 	DeliverOfflineMessages(userID string, receiverID string) ([]dtos.MessageResponse, error)
// 	DeliverHistoryMessages(userID string, receiverID string) ([]dtos.MessageResponse, error)
// }

type ChatService interface {
	// CreateMessage handles saving a new message to the database.
	CreateMessage(chatroomID, senderID, content string) error

	// GetChatroomMembers retrieves a list of members in the specified chatroom.
	GetChatroomMembers(chatroomID string) ([]string, error)

	// MarkMessageAsRead marks a message as read by the specified user.
	MarkMessageAsRead(messageID, userID string) error

	// JoinChatroom adds a user to the specified chatroom.
	JoinChatroom(userID, chatroomID string) error

	// LeaveChatroom removes a user from the specified chatroom.
	LeaveChatroom(userID, chatroomID string) error

	// GetChatHistory retrieves the chat history for the specified chatroom with pagination (limit and offset).
	GetChatHistory(chatroomID string, limit int, offset int) ([]dtos.MessageDTO, error)

	// CreateChatroom creates a new chatroom with the specified name and privacy setting.
	CreateChatroom(name string, members []string, isPrivate bool) (string, error)

	// GetChatroomByUserID retrieves the chatroom ID for a user.
	GetChatroomsForUser(userID string, limit int, offset int) ([]dtos.ChatroomDTO, error)
}
