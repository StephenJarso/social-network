package models

import (
    "time"
    "github.com/google/uuid"
)

// Notification represents a user notification
type Notification struct {
    ID          uuid.UUID  `json:"id" db:"id"`
    RecipientID uuid.UUID  `json:"recipient_id" db:"recipient_id"`
    InitiatorID *uuid.UUID `json:"initiator_id,omitempty" db:"initiator_id"`
    Type        string     `json:"type" db:"type"`
    ReferenceID string     `json:"reference_id" db:"reference_id"` // Can be UUID as string or other identifier
    Message     string     `json:"message" db:"message"`
    IsRead      bool       `json:"is_read" db:"is_read"`
    CreatedAt   time.Time  `json:"created_at" db:"created_at"`
}