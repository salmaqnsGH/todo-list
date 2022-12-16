package activity

import (
	"time"
)

type Activity struct {
	ID        int        `json:"id,omitempty"`
	Email     string     `json:"email,omitempty"`
	Title     string     `json:"title,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}
