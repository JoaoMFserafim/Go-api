package service

import (
	"backend/models"
	"backend/repository"
)

type TaskService struct {
	Repository *repository.TaskRepository
}

func NewTaskService(repository *repository.TaskRepository) *TaskService {
	return &TaskService{Repository: repository}
}

func (service *TaskService) CreateTask(task *models.Task) error {
	return service.Repository.CreateTask(task)
}

func (service *TaskService) FindAllTasks() ([]models.Task, error) {
	return service.Repository.FindAll()
}

func (service *TaskService) FindOneTask(id int) (models.Task, error) {
	return service.Repository.FindOneTask(id)
}

func (service *TaskService) DeleteOneTask(id int) (models.Task, error) {
	return service.Repository.DeleteOneTask(id)
}

func (service *TaskService) EditOne(id int, input models.Task) (models.Task, error) {
	return service.Repository.EditTask(id, input)
}

func (service *TaskService) EditStatusTask(id int) (models.Task, error) {
	var task models.Task
	service.Repository.Db.Find(&task, id)
	service.Repository.Db.Save(task)
	if task.Status {
		return service.Repository.EditStatusTask(id, false)
	} else {
		return service.Repository.EditStatusTask(id, true)
	}
}
