package models

import (
    "time"
    "github.com/google/uuid"
)

// Session represents a user's login session (cookie-based)
type Session struct {
    ID         uuid.UUID  `json:"id" db:"id"`           // session token (UUID)
    UserID     uuid.UUID  `json:"user_id" db:"user_id"` // foreign key to users table
    ExpiresAt  time.Time  `json:"expires_at" db:"expires_at"`
    CreatedAt  time.Time  `json:"created_at" db:"created_at"`
}