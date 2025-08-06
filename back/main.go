package main

import (
    "github.com/gin-gonic/gin"
    "ia-api/handlers"
    "github.com/joho/godotenv"
    "log"
    "ia-api/bd_sqlite"
    "ia-api/models"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Erreur chargement .env")
    }
    db.Connect()

    // Crée la table messages si elle n'existe pas
    db.DB.AutoMigrate(&models.Message{})

    r := gin.Default()

    // Regroupe les routes ici
    r.POST("/chat", handlers.ChatHandler)

    r.Run(":8080")
}
