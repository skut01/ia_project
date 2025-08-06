package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"ia-api/models"
	"ia-api/services"
	"ia-api/bd_sqlite"
)

func ChatHandler(c *gin.Context) {
	var req models.MessageRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format JSON invalide"})
		return
	}

	// 1. Récupérer l’historique depuis la DB
	var history []models.Message
	db.DB.Where("conversation_id = ?", req.ConversationID).Order("created_at asc").Find(&history)

	// 2. Convertir en []ChatMessage pour l’appel API
	messages := make([]models.ChatMessage, len(history))
	for i, m := range history {
		messages[i] = models.ChatMessage{
			Role:    m.Role,
			Content: m.Content,
		}
	}

	// 3. Ajouter le nouveau message reçu
	messages = append(messages, req.Message)

	// 4. Appeler l’API LLM
	response, err := services.CallOpenRouter(messages, req.Model)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 5. Sauvegarder le message utilisateur
	db.DB.Create(&models.Message{
		Content:        req.Message.Content,
		Role:           req.Message.Role,
		ConversationID: &req.ConversationID,
	})

	// 6. Sauvegarder la réponse assistant
	db.DB.Create(&models.Message{
		Content:        response,
		Role:           "assistant",
		ConversationID: &req.ConversationID,
	})

	// 7. Retourner la réponse au client
	c.JSON(http.StatusOK, gin.H{"response": response})
}

/*Exemple de requête JSON a envoyer :{
  "conversation_id": 1,
  "model": "openrouter/horizon-beta",
  "message": {
    "role": "user",
    "content": "quelle était ma derniere question ?"
  }
}
  */
