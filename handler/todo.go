package handler

import (
	"fmt"
	"net/http"
	"todo-list/helper"
	"todo-list/todo"

	"github.com/gin-gonic/gin"
)

type todoHandler struct {
	service todo.Service
}

func NewTodoHandler(service todo.Service) *todoHandler {
	return &todoHandler{service}
}

func (h *todoHandler) GetTodos(c *gin.Context) {
	activityID := c.Query("activity_group_id")
	if activityID == "" {
		todos, err := h.service.GetTodos()
		if err != nil {
			response := helper.APIResponse("Failed get todos", http.StatusBadRequest, "error", nil)
			c.JSON(http.StatusBadRequest, response)
			return
		}

		response := helper.APIResponse("Success", http.StatusOK, "Success", todo.FormatTodos(todos))
		c.JSON(http.StatusOK, response)

	} else {

		err := c.ShouldBindUri(&activityID)
		if err != nil {
			response := helper.APIResponse("Failed to get todos", http.StatusBadRequest, "error", nil)
			c.JSON(http.StatusBadRequest, response)
			return
		}

		todos, err := h.service.GetTodosByActivityID(activityID)
		if err != nil {
			response := helper.APIResponse("Failed get todos", http.StatusBadRequest, "error", nil)
			c.JSON(http.StatusBadRequest, response)
			return
		}

		response := helper.APIResponse("Success", http.StatusOK, "Success", todo.FormatTodos(todos))
		c.JSON(http.StatusOK, response)
	}
}

func (h *todoHandler) GetTodoById(c *gin.Context) {
	var input todo.TodoIdInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to get todo", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	todoDetail, err := h.service.GetTodoByID(input)
	if err != nil {
		errMessage := fmt.Sprintf("Todo with ID %v Not Found", input)

		response := helper.FormatNotFoundError(errMessage, todoDetail)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success", http.StatusOK, "Success", todo.FormatTodo(todoDetail))
	c.JSON(http.StatusOK, response)
}

func (h *todoHandler) CreateTodo(c *gin.Context) {
	var input todo.CreateTodoInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var todo todo.Todo
		response := helper.FormatBadRequest("title cannot be null 1", todo)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// TODO: refactor error
	newTodo, err := h.service.CreateTodo(input)
	if err != nil {
		var todo todo.Todo
		response := helper.FormatBadRequest(fmt.Sprint(err), todo)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIResponse("Success", http.StatusOK, "Success", todo.FormatCreateTodo(newTodo))
	c.JSON(http.StatusOK, response)
}
