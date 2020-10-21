package models

import (
	"time"
)

// Campaign Struct
type Campaign struct {
	ID        uint      `json:"id"`
	Name      string    `json:"campaign_name"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
