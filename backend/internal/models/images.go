package models

import (
    "time"
    "github.com/google/uuid"
)

// Image represents an uploaded image file (JPEG, PNG, GIF)
type Image struct {
    ID         uuid.UUID `json:"id" db:"id"`
    FilePath   string    `json:"file_path" db:"file_path"`
    FileName   string    `json:"file_name" db:"file_name"`
    FileSize   int64     `json:"file_size" db:"file_size"`
    MimeType   string    `json:"mime_type" db:"mime_type"`
    EntityType string    `json:"entity_type" db:"entity_type"`
    EntityID   string    `json:"entity_id" db:"entity_id"` // UUID as string for flexibility
    CreatedAt  time.Time `json:"created_at" db:"created_at"`
    UpdatedAt time.Time `json:"updated_at" db:"updated_at"`   
}