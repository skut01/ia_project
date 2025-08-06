package models

import (
    "gorm.io/gorm"
)

type Conversation struct {
    gorm.Model
    Title    string
    Messages []Message
}
