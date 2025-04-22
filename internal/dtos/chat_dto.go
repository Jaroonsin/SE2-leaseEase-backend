package dtos

// MessageDTO represents the structure of the message sent over WebSocket
type MessageDTO struct {
	Type       string `json:"type"`
	ChatroomID string `json:"chatroom_id"`
	SenderID   string `json:"sender_id"`
	Content    string `json:"content"`
	MessageID  string `json:"message_id,omitempty"` // Used when marking a message as read
	Timestamp  string `json:"timestamp,omitempty"`
	Limit      int    `json:"limit,omitempty"`
	Offset     int    `json:"offset,omitempty"`
}

// ReadStatusDTO represents the structure for marking a message as read
type ReadStatusDTO struct {
	MessageID string `json:"message_id"`
	UserID    string `json:"user_id"`
}

// JoinChatroomDTO represents the structure for a user joining a chatroom
type JoinChatroomDTO struct {
	ChatroomID string `json:"chatroom_id"`
	UserID     string `json:"user_id"`
}

// LeaveChatroomDTO represents the structure for a user leaving a chatroom
type LeaveChatroomDTO struct {
	ChatroomID string `json:"chatroom_id"`
	UserID     string `json:"user_id"`
}

// CreateChatroomDTO represents the structure for creating a new chatroom
type CreateChatroomDTO struct {
	Name      string   `json:"name"`
	Members   []string `json:"members"` // List of user IDs to add to the chatroom
	IsPrivate bool     `json:"is_private"`
}

// ChatroomDTO represents the structure of a chatroom
type ChatroomDTO struct {
	ChatroomID           string `json:"chatroom_id"`
	Name                 string `json:"name"`
	IsPrivate            bool   `json:"is_private"`
	LastReadMessageID    string `json:"last_read_message_id"`
	LastMessageID        string `json:"last_message_id"`
	LastMessageContent   string `json:"last_message_content"`
	LastMessageTimestamp string `json:"last_message_timestamp"`
}
