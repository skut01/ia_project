package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"ia-api/models"
	"ia-api/services"
	"ia-api/bd_sqlite"
    "log"
	"ia-api/globals"
)

func ChatHandler(c *gin.Context) {
	var msg models.MessageRequest
	conv_id := &globals.DefaultConversation.ID
	if err := c.BindJSON(&msg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format JSON invalide"})
		return
	}

	// 1. Enregistrer chaque message reçu
	for _, m := range msg.Messages {
		message := models.Message{
			Content:        m.Content,
			Role:           m.Role,
			ConversationID: conv_id,
			ParentID:       nil, // si tu veux gérer les forks plus tard
		}
		db.DB.Create(&message)
        log.Printf("Message sauvegardé : %s (%s)", m.Content, m.Role)
	}

    

	// 2. Appeler l'IA avec tous les messages
	response, err := services.CallOpenRouter(msg.Messages, msg.Model)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 3. Enregistrer la réponse IA
	assistantMessage := models.Message{
		Content:        response,
		Role:           "assistant",
		ConversationID: conv_id,
	}
    
	db.DB.Create(&assistantMessage)
    log.Printf("Réponse IA sauvegardée : %s", response)
	// 4. Retourner la réponse
	c.JSON(http.StatusOK, gin.H{"response": response})
}