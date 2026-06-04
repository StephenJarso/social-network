package models

import (
    "time"
    "github.com/google/uuid"
)

// PrivateMessage represents a 1-on-1 private message between users
type PrivateMessage struct {
    ID          uuid.UUID `json:"id" db:"id"`
    SenderID    uuid.UUID `json:"sender_id" db:"sender_id"`
    RecipientID uuid.UUID `json:"recipient_id" db:"recipient_id"`
    Content     string    `json:"content" db:"content"`
    IsRead      bool      `json:"is_read" db:"is_read"`
    CreatedAt   time.Time `json:"created_at" db:"created_at"`
}

// GroupMessage represents a message sent in a group chat room
type GroupMessage struct {
    ID        uuid.UUID `json:"id" db:"id"`
    GroupID   uuid.UUID `json:"group_id" db:"group_id"`
    SenderID  uuid.UUID `json:"sender_id" db:"sender_id"`
    Content   string    `json:"content" db:"content"`
    CreatedAt time.Time `json:"created_at" db:"created_at"`
}