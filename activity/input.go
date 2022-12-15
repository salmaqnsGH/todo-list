package activity

type GetActivityByIdInput struct {
	ID int `uri:"id" binding:"required"`
}
