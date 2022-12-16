package todo

import (
	"fmt"
	"time"
)

type TodoFormatter struct {
	ID              int        `json:"id"`
	ActivityGroupId string     `json:"activity_group_id"`
	Title           string     `json:"title"`
	IsActive        string     `json:"is_active"`
	Priority        string     `json:"priority"`
	CreatedAt       *time.Time `json:"created_at"`
	UpdatedAt       *time.Time `json:"updated_at"`
	DeletedAt       *time.Time `json:"deleted_at"`
}

func FormatTodos(todos []Todo) []TodoFormatter {
	todosFormatter := []TodoFormatter{}

	for _, todo := range todos {
		todoFormatter := FormatTodo(todo)
		todosFormatter = append(todosFormatter, todoFormatter)
	}

	return todosFormatter
}

func FormatTodo(todo Todo) TodoFormatter {
	todoFormatter := TodoFormatter{}
	todoFormatter.ID = todo.ID
	todoFormatter.ActivityGroupId = fmt.Sprint(todo.ActivityGroupId)
	todoFormatter.Title = todo.Title
	todoFormatter.IsActive = todo.IsActive
	todoFormatter.Priority = todo.Priority
	todoFormatter.CreatedAt = todo.CreatedAt
	todoFormatter.UpdatedAt = todo.UpdatedAt

	if todo.DeletedAt == nil {
		todoFormatter.DeletedAt = nil
	}

	return todoFormatter
}

type CreateTodoResponse struct {
	CreatedAt       *time.Time `json:"created_at"`
	UpdatedAt       *time.Time `json:"updated_at"`
	ID              int        `json:"id"`
	ActivityGroupId string     `json:"activity_group_id"`
	Title           string     `json:"title"`
	IsActive        string     `json:"is_active"`
	Priority        string     `json:"priority"`
}

func FormatCreateTodo(todo Todo) CreateTodoResponse {
	var createTodoResponse CreateTodoResponse
	createTodoResponse.CreatedAt = todo.CreatedAt
	createTodoResponse.UpdatedAt = todo.UpdatedAt
	createTodoResponse.ID = todo.ID
	createTodoResponse.ActivityGroupId = fmt.Sprint(todo.ActivityGroupId)
	createTodoResponse.Title = todo.Title
	createTodoResponse.IsActive = todo.IsActive
	createTodoResponse.Priority = todo.Priority

	return createTodoResponse
}
