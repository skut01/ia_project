package db

import (
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "log"
)

var DB *gorm.DB

func Connect() {
    var err error
    DB, err = gorm.Open(sqlite.Open("conversation.db"), &gorm.Config{})
    if err != nil {
        log.Fatal("Erreur de connexion à SQLite :", err)
    }
}
