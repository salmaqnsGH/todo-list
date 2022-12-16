package todo

type Service interface {
	GetTodosByActivityID(activityID string) ([]Todo, error)
	GetTodos() ([]Todo, error)
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
