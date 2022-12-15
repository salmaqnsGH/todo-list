package activity

type Service interface {
	GetActivities() ([]Activity, error)
	GetActivityByID(input GetActivityByIdInput) (Activity, error)
	CreateActivity(input CreateActivityInput) (Activity, error)
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

func (s *service) GetActivityByID(input GetActivityByIdInput) (Activity, error) {
	activity, err := s.repository.FindByID(input.ID)
	if err != nil {
		return activity, err
	}

	return activity, nil
}

func (s *service) CreateActivity(input CreateActivityInput) (Activity, error) {
	activity := Activity{}
	activity.Email = input.Email
	activity.Title = input.Title

	newActivity, err := s.repository.Create(activity)
	if err != nil {
		return newActivity, err
	}

	return newActivity, nil
}
