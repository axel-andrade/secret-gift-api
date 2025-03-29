package domain

import (
	"time"
)

// Base defines common fields and methods for all entities.
type Base struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// NewBase initializes a new Base with timestamps.
func NewBase(id string) Base {
	now := time.Now()
	return Base{
		ID:        id,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

func (b *Base) UpdateTimestamp() {
	b.UpdatedAt = time.Now()
}
