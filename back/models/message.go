package models

import "gorm.io/gorm"

type ChatMessage struct {
	Role    string `json:"role"`    // "user", "assistant", etc.
	Content string `json:"content"`
}

type MessageRequest struct {
	Messages       []ChatMessage `json:"messages"`
	Model          string        `json:"model"`
	ConversationID uint          `json:"conversation_id"` // pour regrouper
}

type Message struct {
    gorm.Model
    Content        string
    Role           string
    ConversationID *uint
    Conversation   Conversation `gorm:"foreignKey:ConversationID"`
    ParentID       *uint
}