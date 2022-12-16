package todo

import (
	"strconv"

	"gorm.io/gorm"
)

type Repository interface {
	FindAllByActivityId(activityID string) ([]Todo, error)
	FindAll() ([]Todo, error)
	FindByID(ID int) (Todo, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAllByActivityId(activityID string) ([]Todo, error) {
	var todos []Todo
	activity_id, _ := strconv.Atoi(activityID)

	err := r.db.Where("activity_group_id = ?", activity_id).Find(&todos).Error

	if err != nil {
		return todos, err
	}

	return todos, nil
}

func (r *repository) FindAll() ([]Todo, error) {
	var todos []Todo

	err := r.db.Find(&todos).Error

	if err != nil {
		return todos, err
	}

	return todos, nil
}

func (r *repository) FindByID(ID int) (Todo, error) {
	var todo Todo
	if err := r.db.Where("id = ?", ID).First(&todo).Error; err != nil {
		return todo, err
	}

	return todo, nil
}
