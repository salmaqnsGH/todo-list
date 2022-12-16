package todo

type TodoIdInput struct {
	ID int `uri:"id" binding:"required"`
}
