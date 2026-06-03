package models

import (
    "time"
    "github.com/google/uuid"
)

// Comment represents a comment on a post
type Comment struct {
    ID        uuid.UUID  `json:"id" db:"id"`
    UserID    uuid.UUID  `json:"user_id" db:"user_id"`
    PostID    uuid.UUID  `json:"post_id" db:"post_id"`
    ParentID  *uuid.UUID `json:"parent_id,omitempty" db:"parent_id"`
    Content   string     `json:"content" db:"content"`
    ImageID   *uuid.UUID `json:"image_id,omitempty" db:"image_id"`
    CreatedAt time.Time  `json:"created_at" db:"created_at"`
    UpdatedAt time.Time  `json:"updated_at" db:"updated_at"`
    DeletedAt *time.Time `json:"-" db:"deleted_at"`
}