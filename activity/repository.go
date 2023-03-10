package activity

import (
	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]Activity, error)
	FindByID(ID int) (Activity, error)
	Create(activity Activity) (Activity, error)
	Delete(ID int) error
	Update(activity Activity) (Activity, error)
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

func (r *repository) FindByID(ID int) (Activity, error) {
	var activity Activity
	if err := r.db.Where("id = ?", ID).First(&activity).Error; err != nil {
		return activity, err
	}

	return activity, nil
}

func (r *repository) Create(activity Activity) (Activity, error) {
	err := r.db.Create(&activity).Error

	if err != nil {
		return activity, err
	}

	return activity, nil
}

func (r *repository) Delete(ID int) error {
	var activity Activity
	if err := r.db.Where("id = ?", ID).First(&activity).Delete(&activity).Error; err != nil {
		return err
	}

	return nil
}

func (r *repository) Update(activity Activity) (Activity, error) {
	err := r.db.Save(&activity).Error

	if err != nil {
		return activity, err
	}

	return activity, nil
}
