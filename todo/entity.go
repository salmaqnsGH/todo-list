package todo

import "time"

type Todo struct {
	ID              int        `json:"id,omitempty"`
	ActivityGroupId int        `json:"activity_group_id,omitempty"`
	Title           string     `json:"title,omitempty"`
	IsActive        string     `json:"is_active,omitempty"`
	Priority        string     `json:"priority,omitempty"`
	CreatedAt       *time.Time `json:"created_at,omitempty"`
	UpdatedAt       *time.Time `json:"updated_at,omitempty"`
	DeletedAt       *time.Time `json:"deleted_at,omitempty"`
}
