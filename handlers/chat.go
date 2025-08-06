package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "ia-api/models"
    "ia-api/services"
)

func ChatHandler(c *gin.Context) {
    var msg models.MessageRequest

    if err := c.BindJSON(&msg); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Format JSON invalide"})
        return
    }

    response, err := services.CallOpenRouter(msg.Messages, msg.Model)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"response": response})
}
