package models

type ChatMessage struct {
    Role    string `json:"role"`    // "user", "assistant", "system"
    Content string `json:"content"` // le texte du message
}

type MessageRequest struct {
    Messages []ChatMessage `json:"messages"`     // 👈 l'historique complet
    Model    string        `json:"model"`        // facultatif
}
