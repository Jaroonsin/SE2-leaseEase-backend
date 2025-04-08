package services

import (
	"LeaseEase/internal/dtos"
	"LeaseEase/internal/repositories"
	"strconv"
	"strings"

	"go.uber.org/zap"
)

type chatService struct {
	chatRepo repositories.ChatRepository
	logger   *zap.Logger
}

func NewChatService(chatRepo repositories.ChatRepository, logger *zap.Logger) ChatService {
	return &chatService{
		chatRepo: chatRepo,
		logger:   logger,
	}
}

// CreateMessage handles saving a new message to the database.
func (s *chatService) CreateMessage(chatroomID, senderID, content string) error {
	logger := s.logger.Named("CreateMessage")
	logger.Info("Creating new message", zap.String("chatroomID", chatroomID), zap.String("senderID", senderID), zap.String("content", content))

	chatroomIDUint, err := strconv.ParseUint(chatroomID, 10, 0)
	if err != nil {
		logger.Error("Invalid chatroom ID", zap.Error(err))
	}

	senderIDUint, err := strconv.ParseUint(senderID, 10, 0)
	if err != nil {
		logger.Error("Invalid sender ID", zap.Error(err))
	}

	// Check if the chatroom exists
	chatroom, err := s.chatRepo.GetChatroomByID(uint(chatroomIDUint))
	if err != nil {
		logger.Error("Chatroom not found", zap.Error(err))
	}
	logger.Info("Chatroom found", zap.String("chatroomID", strconv.FormatUint(uint64(chatroom.ChatroomID), 10)))

	message, err := s.chatRepo.CreateMessage(uint(chatroomIDUint), uint(senderIDUint), content)
	if err != nil {
		logger.Error("Failed to create message", zap.Error(err))
		return err
	}

	// Update the last message in the chatroom
	err = s.chatRepo.UpdateLastMessage(uint(chatroomIDUint), message.MessageID)
	if err != nil {
		logger.Error("Failed to update last message in chatroom", zap.Error(err))
	}
	logger.Info("Last message updated successfully", zap.String("chatroomID", strconv.FormatUint(uint64(chatroomIDUint), 10)))

	logger.Info("Message created successfully", zap.String("messageID", strconv.FormatUint(uint64(message.MessageID), 10)))
	return nil

	// need to handle the case where the chatroom does not exist
	// need to handle the case where the sender is not a member of the chatroom
	// need to handle the case where the message is empty or invalid
	// need to handle the case where the message creation fails
	// need to handle the case where the message is too long or exceeds the limit
}

// Fetch and deliver history messages
func (s *chatService) GetChatHistory(chatroomID string, limit int, offset int) ([]dtos.MessageDTO, error) {
	logger := s.logger.Named("GetChatHistory")
	logger.Info("Fetching chat history", zap.String("chatroomID", chatroomID), zap.Int("limit", limit), zap.Int("offset", offset))

	var messages []dtos.MessageDTO
	historyMessages, err := s.chatRepo.GetHistoryMessages(chatroomID, limit, offset)
	if err != nil {
		return nil, err
	}

	for _, msg := range historyMessages {
		dto := dtos.MessageDTO{
			Type:       "history",
			MessageID:  strconv.FormatUint(uint64(msg.MessageID), 10), // or convert uint to string if necessary
			ChatroomID: strconv.FormatUint(uint64(msg.ChatroomID), 10),
			SenderID:   strconv.FormatUint(uint64(msg.SenderID), 10),
			Content:    msg.Content,
			Timestamp:  msg.Timestamp.String(),
			Limit:      limit,
			Offset:     offset,
		}
		messages = append(messages, dto)
	}

	return messages, nil
}

// GetChatroomMembers retrieves a list of members in the specified chatroom.
func (s *chatService) GetChatroomMembers(chatroomID string) ([]string, error) {
	logger := s.logger.Named("GetChatroomMembers")
	logger.Info("Fetching chatroom members", zap.String("chatroomID", chatroomID))

	chatroomIDUint, err := strconv.ParseUint(chatroomID, 10, 0)
	if err != nil {
		logger.Error("Invalid chatroom ID", zap.Error(err))
		return nil, err
	}

	members, err := s.chatRepo.GetMembersOfChatroom(uint(chatroomIDUint))
	if err != nil {
		logger.Error("Failed to fetch chatroom members", zap.Error(err))
		return nil, err
	}

	for _, member := range members {
		logger.Info("Member found", zap.String("memberID", strconv.FormatUint(uint64(member.ID), 10)))
	}

	var memberIDs []string
	for _, member := range members {
		memberIDs = append(memberIDs, strconv.FormatUint(uint64(member.ID), 10))
	}

	return memberIDs, nil
}

// MarkMessageAsRead marks a message as read by the specified user.
func (s *chatService) MarkMessageAsRead(messageID, userID string) error {
	logger := s.logger.Named("MarkMessageAsRead")
	logger.Info("Marking message as read", zap.String("messageID", messageID), zap.String("userID", userID))

	messageIDUint, err := strconv.ParseUint(messageID, 10, 0)
	if err != nil {
		logger.Error("Invalid message ID", zap.Error(err))
		return err
	}

	userIDUint, err := strconv.ParseUint(userID, 10, 0)
	if err != nil {
		logger.Error("Invalid user ID", zap.Error(err))
		return err
	}

	err = s.chatRepo.MarkMessageAsRead(uint(messageIDUint), uint(userIDUint))
	if err != nil {
		logger.Error("Failed to mark message as read", zap.Error(err))
		return err
	}

	logger.Info("Message marked as read successfully")
	return nil
}

// JoinChatroom adds a user to the specified chatroom.
func (s *chatService) JoinChatroom(userID, chatroomID string) error {
	logger := s.logger.Named("JoinChatroom")
	logger.Info("User joining chatroom", zap.String("userID", userID), zap.String("chatroomID", chatroomID))

	userIDUint, err := strconv.ParseUint(userID, 10, 0)
	if err != nil {
		logger.Error("Invalid user ID", zap.Error(err))
		return err
	}

	chatroomIDUint, err := strconv.ParseUint(chatroomID, 10, 0)
	if err != nil {
		logger.Error("Invalid chatroom ID", zap.Error(err))
		return err
	}

	err = s.chatRepo.AddMemberToChatroom(uint(userIDUint), uint(chatroomIDUint))
	if err != nil {
		logger.Error("Failed to join chatroom", zap.Error(err))
		return err
	}

	logger.Info("User joined chatroom successfully")
	return nil
}

func (s *chatService) LeaveChatroom(userID string, chatroomID string) error {
	logger := s.logger.Named("LeaveChatroom")
	logger.Info("User leaving chatroom", zap.String("userID", userID), zap.String("chatroomID", chatroomID))

	userIDUint, err := strconv.ParseUint(userID, 10, 0)
	if err != nil {
		logger.Error("Invalid user ID", zap.Error(err))
		return err
	}

	chatroomIDUint, err := strconv.ParseUint(chatroomID, 10, 0)
	if err != nil {
		logger.Error("Invalid chatroom ID", zap.Error(err))
		return err
	}

	err = s.chatRepo.RemoveMemberFromChatroom(uint(chatroomIDUint), uint(userIDUint))
	if err != nil {
		logger.Error("Failed to leave chatroom", zap.Error(err))
		return err
	}

	logger.Info("User left chatroom successfully")
	return nil
}

func (s *chatService) CreateChatroom(name string, members []string, isPrivate bool) (string, error) {
	logger := s.logger.Named("CreateChatroom")
	logger.Info("Creating chatroom", zap.String("name", name))

	chatroomID, err := s.chatRepo.CreateChatroom(name, isPrivate)
	if err != nil {
		logger.Error("Failed to create chatroom", zap.Error(err))
		return "", err
	}

	for _, member := range members {
		memberIDUint, err := strconv.ParseUint(member, 10, 0)
		if err != nil {
			logger.Error("Invalid member ID", zap.Error(err))
			return "", err
		}
		err = s.chatRepo.AddMemberToChatroom(uint(chatroomID.ChatroomID), uint(memberIDUint))
		if err != nil {
			logger.Error("Failed to add member to chatroom", zap.Error(err))
			return "", err
		}
	}

	chatroomIDStr := strconv.FormatUint(uint64(chatroomID.ChatroomID), 10)

	logger.Info("Chatroom created successfully", zap.String("chatroomID", chatroomIDStr))
	return chatroomIDStr, nil
}

func (s *chatService) GetChatroomsForUser(userID string, limit int, offset int) (string, error) {
	logger := s.logger.Named("GetChatroomsForUser")
	logger.Info("Fetching chatrooms for user", zap.String("userID", userID), zap.Int("limit", limit), zap.Int("offset", offset))

	userIDUint, err := strconv.ParseUint(userID, 10, 0)
	if err != nil {
		logger.Error("Invalid user ID", zap.Error(err))
		return "", err
	}

	chatrooms, err := s.chatRepo.GetChatroomsByUser(uint(userIDUint), limit, offset)
	if err != nil {
		logger.Error("Failed to fetch chatrooms", zap.Error(err))
		return "", err
	}

	var chatroomIDs []string
	for _, chatroom := range chatrooms {
		chatroomIDs = append(chatroomIDs, strconv.FormatUint(uint64(chatroom.ChatroomID), 10))
	}

	chatroomIDsStr := strings.Join(chatroomIDs, ",")
	logger.Info("Chatrooms fetched successfully", zap.String("chatroomIDs", chatroomIDsStr))
	return chatroomIDsStr, nil
}
