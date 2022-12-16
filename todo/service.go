package todo

type Service interface {
	GetTodosByActivityID(activityID string) ([]Todo, error)
	GetTodos() ([]Todo, error)
	GetTodoByID(input TodoIdInput) (Todo, error)
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
