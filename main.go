package main

import (
    "github.com/gin-gonic/gin"
    "ia-api/handlers"
    "github.com/joho/godotenv"
    "log"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Erreur chargement .env")
    }

    r := gin.Default()

    // Regroupe les routes ici
    r.POST("/chat", handlers.ChatHandler)

    r.Run(":8080")
}
