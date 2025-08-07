package handlers

import (
	"net/http"
	"strconv"

	db "ia-api/bd_sqlite"
	"ia-api/models"

	"github.com/gin-gonic/gin"
)

func GetConversation(c *gin.Context) {
	convIDParam := c.Param("id")

	// Convertir string → uint
	convID, err := strconv.ParseUint(convIDParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de conversation invalide"})
		return
	}

	var messages []models.Message
	result := db.DB.
		Where("conversation_id = ?", uint(convID)).
		Order("created_at asc").
		Find(&messages)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, messages)
}
