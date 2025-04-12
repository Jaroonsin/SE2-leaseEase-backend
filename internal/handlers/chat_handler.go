package handlers

import (
	"LeaseEase/internal/dtos"
	"LeaseEase/internal/services"
	"LeaseEase/utils"
	"log"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

type chatHandler struct {
	chatService services.ChatService
	clients     map[string]*websocket.Conn
	mu          sync.RWMutex // protects clients map
}

func NewChatHandler(chatService services.ChatService) *chatHandler {
	return &chatHandler{
		chatService: chatService,
		clients:     make(map[string]*websocket.Conn),
	}
}

// Handle WebSocket connections
func (h *chatHandler) HandleWebSocket(ws *websocket.Conn) {
	defer ws.Close()

	for {
		var msg dtos.MessageDTO

		if err := ws.ReadJSON(&msg); err != nil {
			log.Println("Error reading message:", err)
			return
		}

		switch msg.Type {
		case "message":
			if err := h.HandleNewMessage(&msg); err != nil {
				log.Println("Error handling message:", err)
			}
		case "read":
			if err := h.HandleMessageRead(msg.ChatroomID, msg.MessageID, msg.SenderID); err != nil {
				log.Println("Error marking message as read:", err)
			}
		case "join":
			h.mu.Lock()
			h.clients[msg.SenderID] = ws
			h.mu.Unlock()

			if err := h.HandleUserJoinChatroom(msg.SenderID, msg.ChatroomID); err != nil {
				log.Println("Error handling join:", err)
			}
		case "leave":
			if err := h.HandleUserLeaveChatroom(msg.SenderID, msg.ChatroomID); err != nil {
				log.Println("Error handling leave:", err)
			}

			h.mu.Lock()
			delete(h.clients, msg.SenderID)
			h.mu.Unlock()
		case "history":
			err := h.HandleRetrieveChatHistory(ws, msg.ChatroomID, msg.Limit, msg.Offset)
			if err != nil {
				log.Println("Error retrieving chat history:", err)
			}
		case "start":
			h.mu.Lock()
			h.clients[msg.SenderID] = ws
			h.mu.Unlock()
			if err := h.HandleStartPage(ws, msg.SenderID, msg.Limit, msg.Offset); err != nil {
				log.Println("Error handling start page:", err)
			}
		default:
			log.Println("Unknown message type:", msg.Type)
		}
	}
}

func (h *chatHandler) HandleNewMessage(message *dtos.MessageDTO) error {
	if err := h.chatService.CreateMessage(message.ChatroomID, message.SenderID, message.Content); err != nil {
		log.Println("Error saving message to database:", err)
		return err
	}

	log.Printf("New message from %s in chatroom %s: %s\n", message.SenderID, message.ChatroomID, message.Content)
	return h.BroadcastMessageToChatroom(message.ChatroomID, message)
}

func (h *chatHandler) BroadcastMessageToChatroom(chatroomID string, message *dtos.MessageDTO) error {
	log.Printf("Broadcasting message to chatroom %s: %+v\n", chatroomID, message)

	// Example: get members from DB (pseudo-code)
	members, err := h.chatService.GetChatroomMembers(chatroomID)
	if err != nil {
		return err
	}
	log.Printf("Members in chatroom %s: %v\n", chatroomID, members)

	h.mu.RLock()
	defer h.mu.RUnlock()

	for _, userID := range members {
		if userID == message.SenderID {
			log.Printf("Skipping message broadcast to sender %s\n", userID)
			continue
		}

		if ws, ok := h.clients[userID]; ok {
			if err := ws.WriteJSON(message); err != nil {
				log.Printf("Error sending message to user %s: %v\n", userID, err)
			}
			log.Printf("Message sent to user %s\n", userID)
		}
	}

	return nil
}

func (h *chatHandler) HandleUserDisconnect(userID string) error {
	h.mu.Lock()
	defer h.mu.Unlock()

	delete(h.clients, userID)
	log.Printf("User %s disconnected\n", userID)
	return nil
}

func (h *chatHandler) HandleMessageRead(chatroomID, messageID, userID string) error {
	log.Printf("Marking message %s as read by user %s\n", messageID, userID)
	return h.chatService.MarkMessageAsRead(chatroomID, messageID, userID)
}

func (h *chatHandler) HandleUserJoinChatroom(userID, chatroomID string) error {
	log.Printf("User %s joined chatroom %s\n", userID, chatroomID)
	return h.chatService.JoinChatroom(userID, chatroomID)
}

func (h *chatHandler) HandleUserLeaveChatroom(userID, chatroomID string) error {
	log.Printf("User %s left chatroom %s\n", userID, chatroomID)
	return h.chatService.LeaveChatroom(userID, chatroomID)
}

func (h *chatHandler) HandleRetrieveChatHistory(ws *websocket.Conn, chatroomID string, limit int, offset int) error {
	messages, err := h.chatService.GetChatHistory(chatroomID, limit, offset)
	if err != nil {
		return err
	}

	for _, msg := range messages {
		if err := ws.WriteJSON(msg); err != nil {
			log.Println("Failed to send message history to user:", err)
		}
	}

	return nil
}

func (h *chatHandler) CreateChatroom(c *fiber.Ctx) error {
	log.Printf("Creating chatroom\n")

	var createChatroomDTO dtos.CreateChatroomDTO
	if err := c.BodyParser(&createChatroomDTO); err != nil {
		log.Println("Error parsing request body:", err)
		return err
	}
	chatroomID, err := h.chatService.CreateChatroom(createChatroomDTO.Name, createChatroomDTO.Members, createChatroomDTO.IsPrivate)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}
	return utils.SuccessResponse(c, fiber.StatusCreated, "Chatroom created successfully", chatroomID)
}

func (h *chatHandler) HandleStartPage(ws *websocket.Conn, userID string, limit int, offset int) error {
	log.Printf("Handling start page for user %s\n", userID)

	// Example: Fetch chatrooms and messages for the user (pseudo-code)
	chatrooms, err := h.chatService.GetChatroomsForUser(userID, limit, offset)
	if err != nil {
		return err
	}

	h.mu.RLock()
	defer h.mu.RUnlock()

	for _, chatroom := range chatrooms {
		if ws, ok := h.clients[userID]; ok {
			if err := ws.WriteJSON(chatroom); err != nil {
				log.Printf("Error sending chatroom data to user %s: %v\n", userID, err)
			}
		}
	}

	return nil
}
