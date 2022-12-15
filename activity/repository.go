package activity

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Activity, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Activity, error) {
	var activities []Activity
	err := r.db.Find(&activities).Error

	if err != nil {
		return activities, err
	}

	return activities, nil
}
