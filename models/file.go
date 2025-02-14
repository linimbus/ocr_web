package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type File struct {
	ID          uuid.UUID `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Content     string    `json:"content" db:"content"`
	Size        int64     `json:"size" db:"size"`
	ContentType string    `json:"content_type" db:"content_type"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}
