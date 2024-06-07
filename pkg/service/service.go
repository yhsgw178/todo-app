package service

import (
	"github.com/yhsgw178/todoapp/pkg/models"
	"github.com/yhsgw178/todoapp/pkg/repository"
)

type TaskService interface {
	CreateTask(task *models.Task) error
	GetTaskByID(id int) (*models.Task, error)
	// GetAllTasks() ([]*models.Task, error)
	// UpdateTask(task *models.Task) error
	// DeleteTask(id int) error
}

type TaskServiceImpl struct {
	repo repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) TaskService {
	return &TaskServiceImpl{
		repo: repo,
	}
}

func (t *TaskServiceImpl) CreateTask(task *models.Task) error {
	return t.repo.CreateTask(task)
}

func (t *TaskServiceImpl) GetTaskByID(id int) (*models.Task, error) {
	return t.repo.GetTaskByID(id)
}
