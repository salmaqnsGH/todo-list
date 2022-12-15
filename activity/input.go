package activity

import "time"

type GetActivityByIdInput struct {
	ID int `uri:"id" binding:"required"`
}

type CreateActivityInput struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
