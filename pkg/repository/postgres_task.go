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

func (r *PostgresTaskRepository) GetTaskByID(id int) (*models.Task, error) {
	query := `SELECT id, title, description, is_completed, due_date, task_status, created_at, updated_at FROM tasks WHERE id = $1;`

	row := r.DB.QueryRow(query, id)
	return scanTask(row)
}

func scanTask(rows *sql.Row) (*models.Task, error) {
	task := &models.Task{}
	err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.IsCompleted, &task.DueDate, &task.TaskStatus, &task.CreatedAt, &task.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return task, nil
}
