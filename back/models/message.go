package models

import "gorm.io/gorm"

type ChatMessage struct {
	Role    string `json:"role"`    // "user", "assistant", etc.
	Content string `json:"content"`
}

type MessageRequest struct {
	Message       ChatMessage `json:"message"`
	Model          string        `json:"model"`
	ConversationID uint          `json:"conversation_id"` // pour regrouper
}

type Message struct {
    gorm.Model
    Content        string
    Role           string
    ConversationID *uint
    ParentID       *uint
    Parent         *Message        `gorm:"foreignKey:ParentID"`
    Children       []Message       `gorm:"foreignKey:ParentID"`
}