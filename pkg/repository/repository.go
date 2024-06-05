package repository

import (
	"github.com/yhsgw178/todoapp/pkg/models"
)

type TaskRepository interface {
	CreateTask(task *models.Task) error
	// GetTaskByID(id int) (*models.Task, error)
	// GetAllTasks() ([]*models.Task, error)
	// UpdateTask(task *models.Task) error
	// DeleteTask(id int) error
}
