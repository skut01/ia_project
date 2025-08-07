package main

import (
	db "ia-api/bd_sqlite"
	"ia-api/globals"
	"ia-api/handlers"
	"ia-api/models"
	"log"
    "github.com/gin-contrib/cors"
    "time"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
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

    // CORS
    r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // frontend Vite
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Regroupe les routes ici
	r.POST("/chat", handlers.ChatHandler)
	r.GET("/conv/:id", handlers.GetConversation)

	r.Run(":8080")
}

func createDefaultConv() {
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
