package activity

type Service interface {
	GetActivities() ([]Activity, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetActivities() ([]Activity, error) {
	activities, err := s.repository.FindAll()
	if err != nil {
		return activities, err
	}

	return activities, nil
}
