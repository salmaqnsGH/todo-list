package activity

import "time"

type Service interface {
	GetActivities() ([]Activity, error)
	GetActivityByID(input ActivityIdInput) (Activity, error)
	CreateActivity(input CreateActivityInput) (Activity, error)
	DeleteActivity(input ActivityIdInput) error
	UpdateActivity(inputID ActivityIdInput, inputData CreateActivityInput) (Activity, error)
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

func (s *service) GetActivityByID(input ActivityIdInput) (Activity, error) {
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

func (s *service) DeleteActivity(input ActivityIdInput) error {
	err := s.repository.Delete(input.ID)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) UpdateActivity(inputID ActivityIdInput, inputData CreateActivityInput) (Activity, error) {
	activity, err := s.repository.FindByID(inputID.ID)
	if err != nil {
		return activity, err
	}

	current_time := time.Now()
	activity.Title = inputData.Title
	activity.Email = inputData.Email
	activity.UpdatedAt = &current_time

	updatedActivity, err := s.repository.Update(activity)
	if err != nil {
		return updatedActivity, err
	}

	return updatedActivity, nil
}
