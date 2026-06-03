package models

import (
    "time"
    "github.com/google/uuid"
)

// Post represents a post created by a user
type Post struct {
    ID           uuid.UUID  `json:"id" db:"id"`
    UserID       uuid.UUID  `json:"user_id" db:"user_id"`
    Content      string     `json:"content" db:"content"`
    ImageID      *uuid.UUID `json:"image_id,omitempty" db:"image_id"`
    PrivacyLevel string     `json:"privacy_level" db:"privacy_level"`
    CreatedAt    time.Time  `json:"created_at" db:"created_at"`
    UpdatedAt    time.Time  `json:"updated_at" db:"updated_at"`
    DeletedAt    *time.Time `json:"-" db:"deleted_at"`
}

// PostRecipient specifies which specific users can see a "private" post
type PostRecipient struct {
    ID        uuid.UUID `json:"id" db:"id"`
    PostID    uuid.UUID `json:"post_id" db:"post_id"`
    UserID    uuid.UUID `json:"user_id" db:"user_id"`
    CreatedAt time.Time `json:"created_at" db:"created_at"`
}