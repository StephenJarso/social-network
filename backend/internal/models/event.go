package models

import (
    "time"
    "github.com/google/uuid"
)

// Event represents an event created within a group
type Event struct {
    ID          uuid.UUID `json:"id" db:"id"`
    GroupID     uuid.UUID `json:"group_id" db:"group_id"`
    Title       string    `json:"title" db:"title"`
    Description string    `json:"description" db:"description"`
    DateTime    time.Time `json:"date_time" db:"date_time"`
    CreatedBy   uuid.UUID `json:"created_by" db:"created_by"`
    CreatedAt   time.Time `json:"created_at" db:"created_at"`
    UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

// EventResponse represents a user's RSVP to an event
type EventResponse struct {
    ID        uuid.UUID `json:"id" db:"id"`
    EventID   uuid.UUID `json:"event_id" db:"event_id"`
    UserID    uuid.UUID `json:"user_id" db:"user_id"`
    Response  string    `json:"response" db:"response"`
    CreatedAt time.Time `json:"created_at" db:"created_at"`
    UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}