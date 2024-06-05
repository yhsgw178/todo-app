package repository

import (
	"database/sql"

	"github.com/yhsgw178/todoapp/pkg/models"
)

type PostgresTaskRepository struct {
	DB *sql.DB
}

func NewPostgresTaskRepository(db *sql.DB) *PostgresTaskRepository {
	return &PostgresTaskRepository{DB: db}
}

func (r *PostgresTaskRepository) CreateTask(task *models.Task) error {
	query := `
	INSERT INTO tasks (title, description, is_completed, due_date, task_status, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5, NOW(), NOW())
	RETURNING id, created_at, updated_at;
`
	err := r.DB.QueryRow(query, task.Title, task.Description, task.IsCompleted, task.DueDate, task.TaskStatus).
		Scan(&task.ID, &task.CreatedAt, &task.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}
