package main

import (
    "github.com/gin-gonic/gin"
    "ia-api/handlers"
    "github.com/joho/godotenv"
    "log"
    "ia-api/bd_sqlite"
    "ia-api/models"
    "ia-api/globals"
)



func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Erreur chargement .env")
    }
    db.Connect()

    // Crée la table messages si elle n'existe pas
    db.DB.AutoMigrate(&models.Conversation{}, &models.Message{})
    createDefaultConv()
    

    r := gin.Default()

    // Regroupe les routes ici
    r.POST("/chat", handlers.ChatHandler)

    r.Run(":8080")
}

func createDefaultConv(){
// Créer la conversation par défaut s’il n’y en a pas
    var existingConv models.Conversation
    result := db.DB.First(&existingConv)

    if result.Error != nil {
        if result.Error.Error() == "record not found" {
            newConv := models.Conversation{
                Title: "Conversation unique",
            }
            db.DB.Create(&newConv)
            globals.DefaultConversation = &newConv
            log.Println("Conversation créée :", newConv.ID)
        } else {
            log.Fatal("Erreur DB :", result.Error)
        }
    } else {
        globals.DefaultConversation = &existingConv
        log.Println("📦 Conversation existante :", existingConv.ID)
    }
}
