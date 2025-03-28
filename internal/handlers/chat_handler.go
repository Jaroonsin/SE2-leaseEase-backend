package handlers

import (
	"LeaseEase/internal/dtos"
	"LeaseEase/internal/services"
	"LeaseEase/utils"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

type chatHandler struct {
	chatService services.ChatService
	clients     map[string]*websocket.Conn
}

func NewChatHandler(chatService services.ChatService) *chatHandler {
	return &chatHandler{
		chatService: chatService,
		clients:     make(map[string]*websocket.Conn),
	}
}

func (h *chatHandler) HandleWebSocketUpgrade(c *fiber.Ctx) error {
	senderID := c.Query("senderID")
	receiverID := c.Query("receiverID")

	if senderID == "" {
		log.Println("Sender ID missing")
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Sender ID missing")
	}

	if receiverID == "" {
		log.Println("Receiver ID missing")
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Receiver ID missing")
	}

	// Upgrade the HTTP connection to WebSocket + add senderID to header
	if websocket.IsWebSocketUpgrade(c) {
		c.Locals("senderID", senderID)
		c.Locals("receiverID", receiverID)
		return c.Next()
	}
	return utils.ErrorResponse(c, fiber.StatusUpgradeRequired, "Upgrade required")
}

// Handle WebSocket connections
func (h *chatHandler) HandleWebSocket(ws *websocket.Conn) {
	defer ws.Close()

	senderID := ws.Locals("senderID").(string)
	receiverID := ws.Locals("receiverID").(string)

	if senderID == "" || receiverID == "" {
		log.Println("Error: Missing senderID or receiverID")
		return
	}

	h.clients[senderID] = ws
	log.Printf("User connected: %s, Total connected users: %d", senderID, len(h.clients))

	//receiverID := strconv.FormatUint(uint64(req.ReceiverID), 10) // Convert receiver ID to string
	log.Println("ReceiverID from message:", receiverID)

	// Deliver history messages
	historyMessages, err := h.chatService.DeliverHistoryMessages(senderID, receiverID)
	if err != nil {
		log.Println("Error fetching chat history:", err)
	} else {
		for _, msg := range historyMessages {
			ws.WriteJSON(msg)
		}
		log.Printf("Finished delivering history messages to %s", senderID)
	}

	// Deliver offline messages
	offlineMessages, err := h.chatService.DeliverOfflineMessages(senderID, receiverID)
	if err == nil {
		for _, msg := range offlineMessages {
			ws.WriteJSON(msg)
		}
		log.Printf("Finished delivering offline messages to %s", senderID)
	} else {
		log.Println("Error fetching offline messages:", err)
	}

	for {
		var req dtos.SendMessageRequest
		if err := ws.ReadJSON(&req); err != nil {
			log.Println("Error reading message:", err)
			delete(h.clients, senderID)
			break
		}

		log.Println("Message received: ", req)

		// Check if receiver is online
		receiverIDStr := strconv.FormatUint(uint64(req.ReceiverID), 10)
		_, receiverOnline := h.clients[receiverIDStr]

		log.Print("Receiver online status: ", receiverOnline)

		// Store message
		err := h.chatService.ProcessMessage(req, receiverOnline)
		if err != nil {
			log.Println("Error processing message:", err)
			continue
		}

		// If receiver is online, send the message
		if receiverOnline {
			h.clients[string(receiverIDStr)].WriteJSON(req)
		}
	}
}
