package repository

import (
	"backend/models"

	"gorm.io/gorm"
)

type TaskRepository struct {
	Db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{Db: db}
}

func (repository *TaskRepository) CreateTask(task *models.Task) error {
	return repository.Db.Create(task).Error
}

func (repository *TaskRepository) FindAll() ([]models.Task, error) {
	var tasks []models.Task

	err := repository.Db.Find(&tasks).Error
	return tasks, err
}

func (repository *TaskRepository) FindOneTask(id int) (models.Task, error) {
	var task models.Task
	err := repository.Db.First(&task, id).Error
	return task, err
}

func (repository *TaskRepository) DeleteOneTask(id int) (models.Task, error) {
	var task models.Task
	err := repository.Db.Delete(&task, id).Error
	return task, err
}

func (repository *TaskRepository) EditTask(id int, input models.Task) (models.Task, error) {
	var task models.Task

	if err := repository.Db.First(&task, id).Error; err != nil {
		return task, err
	}

	task.Title = input.Title
	task.Description = input.Description
	task.Status = input.Status

	if err := repository.Db.Save(&task).Error; err != nil {
		return task, err
	}

	return task, nil
}

func (repository *TaskRepository) EditStatusTask(id int, status bool) (models.Task, error) {
	var task models.Task

	if err := repository.Db.Find(&task, id).Error; err != nil {
		return task, err
	}

	task.Status = status

	if err := repository.Db.Save(&task).Error; err != nil {
		return task, err
	}

	return task, nil
}
