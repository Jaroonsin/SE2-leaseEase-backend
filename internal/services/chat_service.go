package services

import (
	"LeaseEase/internal/dtos"
	"LeaseEase/internal/models"
	"LeaseEase/internal/repositories"
	"time"

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

// Process incoming message
func (s *chatService) ProcessMessage(msg dtos.SendMessageRequest, isReceiverOnline bool) error {
	logger := s.logger.Named("ProcessMessage")
	logger.Info("Processing message", zap.String("senderID", string(msg.SenderID)), zap.String("receiverID", string(msg.ReceiverID)))

	storedMsg := models.Message{
		SenderID:   msg.SenderID,
		ReceiverID: msg.ReceiverID,
		Content:    msg.Content,
		Timestamp:  time.Now(),
		Delivered:  isReceiverOnline,
	}

	err := s.chatRepo.StoreMessage(storedMsg)
	if err != nil {
		logger.Error("Failed to get undelivered messages", zap.Error(err))
		return err
	}

	return nil
}

// Fetch and deliver undelivered messages
func (s *chatService) DeliverOfflineMessages(userID string, receiverID string) ([]dtos.MessageResponse, error) {
	logger := s.logger.Named("DeliverOfflineMessages")
	logger.Info("Fetching undelivered messages", zap.String("userID", userID))
	messages, err := s.chatRepo.GetUndeliveredMessages(userID, receiverID)
	if err != nil {
		return nil, err
	}

	var messageResponses []dtos.MessageResponse
	for _, msg := range messages {
		messageResponses = append(messageResponses, dtos.MessageResponse{
			ID:         msg.ID,
			SenderID:   msg.SenderID,
			ReceiverID: msg.ReceiverID,
			Content:    msg.Content,
			Timestamp:  msg.Timestamp,
			Delivered:  msg.Delivered,
		})
	}

	err = s.chatRepo.MarkMessagesAsDelivered(userID, receiverID)
	if err != nil {
		logger.Error("Failed to mark messages as delivered", zap.Error(err))
		return nil, err
	}

	logger.Info("Undelivered messages fetched and marked as delivered", zap.Int("count", len(messageResponses)))
	return messageResponses, nil
}

// Fetch and deliver history messages
func (s *chatService) DeliverHistoryMessages(userID string, receiverID string) ([]dtos.MessageResponse, error) {
	logger := s.logger.Named("DeliverHistoryMessages")
	logger.Info("Fetching history messages", zap.String("userID", userID), zap.String("receiverID", receiverID))

	// Fetch message history between sender and receiver
	messages, err := s.chatRepo.GetHistoryMessages(userID, receiverID)
	if err != nil {
		return nil, err
	}

	var messageResponses []dtos.MessageResponse
	for _, msg := range messages {
		messageResponses = append(messageResponses, dtos.MessageResponse{
			ID:         msg.ID,
			SenderID:   msg.SenderID,
			ReceiverID: msg.ReceiverID,
			Content:    msg.Content,
			Timestamp:  msg.Timestamp,
			Delivered:  msg.Delivered,
		})
	}

	logger.Info("History messages fetched", zap.Int("count", len(messageResponses)))
	return messageResponses, nil
}
