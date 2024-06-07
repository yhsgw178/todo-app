package models

import (
	"time"
)

type TaskStatus int

const (
	TODO TaskStatus = iota
	Processing
	Done
	Pending
)

// Task represents a task in the to-do list.
type Task struct {
	ID          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description,omitempty"`
	IsCompleted bool       `json:"is_completed"`
	DueDate     time.Time  `json:"due_date"`
	TaskStatus  TaskStatus `json:"task_status"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}
