package domain

import (
	"time"

	"github.com/google/uuid"
)

type Wish struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"name"`
	Description string    `json:"description"`
	Status      Status    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}
