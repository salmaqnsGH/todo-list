package todo

import "time"

type Service interface {
	GetTodosByActivityID(activityID string) ([]Todo, error)
	GetTodos() ([]Todo, error)
	GetTodoByID(input TodoIdInput) (Todo, error)
	CreateTodo(input CreateTodoInput) (Todo, error)
	DeleteTodo(input TodoIdInput) error
	UpdateTodo(inputID TodoIdInput, inputData CreateTodoInput) (Todo, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetTodosByActivityID(activityID string) ([]Todo, error) {
	todos, err := s.repository.FindAllByActivityId(activityID)
	if err != nil {
		return todos, err
	}

	return todos, nil
}

func (s *service) GetTodos() ([]Todo, error) {
	todos, err := s.repository.FindAll()
	if err != nil {
		return todos, err
	}

	return todos, nil
}

func (s *service) GetTodoByID(input TodoIdInput) (Todo, error) {
	todo, err := s.repository.FindByID(input.ID)
	if err != nil {
		return todo, err
	}

	return todo, nil
}

func (s *service) CreateTodo(input CreateTodoInput) (Todo, error) {
	todo := Todo{}
	todo.ActivityGroupId = input.ActivityGroupId
	todo.Title = input.Title

	newTodo, err := s.repository.Create(todo)
	if err != nil {
		return newTodo, err
	}

	return newTodo, nil
}

func (s *service) DeleteTodo(input TodoIdInput) error {
	err := s.repository.Delete(input.ID)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) UpdateTodo(inputID TodoIdInput, inputData CreateTodoInput) (Todo, error) {
	todo, err := s.repository.FindByID(inputID.ID)
	if err != nil {
		return todo, err
	}

	current_time := time.Now()
	todo.Title = inputData.Title
	todo.UpdatedAt = &current_time

	updatedTodo, err := s.repository.Update(todo)
	if err != nil {
		return updatedTodo, err
	}

	return updatedTodo, nil
}
