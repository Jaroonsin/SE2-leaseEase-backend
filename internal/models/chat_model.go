package models

import (
	"time"
)

type Chatroom struct {
	ChatroomID         uint      `gorm:"primaryKey"`
	Name               string    `gorm:"type:varchar(255);default:null" json:"name,omitempty"`
	IsPrivate          bool      `gorm:"default:false" json:"is_private"`
	LastMessageID      *uint     `gorm:"default:0"`
	LastMessageContent string    `gorm:"type:text;default:null" json:"last_message_content,omitempty"`
	CreatedAt          time.Time `json:"created_at"`

	// Relationships
	// LastMessage *Message `gorm:"foreignKey:LastMessageID;references:MessageID;constraint:OnDelete:SET NULL"`
	Messages []Message        `gorm:"foreignKey:ChatroomID;reference:ChatroomID" json:"messages,omitempty"`
	Members  []ChatroomMember `gorm:"foreignKey:ChatroomID;"`
}

type ChatroomMember struct {
	ChatroomID        uint      `gorm:"primaryKey" json:"chatroom_id"`
	UserID            uint      `gorm:"primaryKey" json:"user_id"`
	JoinedAt          time.Time `json:"joined_at"`
	LastSeenMessageID uint      `gorm:"default:0"`

	// Relationships
	// Chatroom Chatroom `gorm:"foreignKey:ChatroomID;references:ChatroomID;"`
	User User `gorm:"foreignKey:UserID;references:ID"`
}

type Message struct {
	MessageID  uint      `gorm:"primaryKey;autoIncrement" json:"message_id"`
	ChatroomID uint      `gorm:"not null"`
	SenderID   uint      `json:"sender_id"`
	Content    string    `gorm:"type:text" json:"content"`
	Timestamp  time.Time `gorm:"autoCreateTime" json:"sent_at"`

	// Relationships

	// Chatroom Chatroom `gorm:"foreignKey:ChatroomID;references:ChatroomID" json:"chatroom,omitempty"`
	Sender User `gorm:"foreignKey:SenderID;references:ID" json:"sender,omitempty"`
	// MessageReads []MessageRead `gorm:"foreignKey:MessageID" json:"reads,omitempty"`
}

type MessageRead struct {
	MessageID uint      `gorm:"primaryKey;references:MessageID" json:"message_id"`
	UserID    uint      `gorm:"primaryKey" json:"user_id"`
	ReadAt    time.Time `json:"read_at"`

	// Relationships
	// Message Message `gorm:"foreignKey:MessageID" json:"message,omitempty"`
	// User User `gorm:"foreignKey:UserID" json:"user,omitempty"`
}
