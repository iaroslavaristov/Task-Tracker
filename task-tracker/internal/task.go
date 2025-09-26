package internal

import "time"

type Status string

const (
	StatusToDo       Status = "todo"
	StatusDone       Status = "done"
	StatusInProgress        = "in-progress"
	StatusNone              = "none"
)

type Task struct {
	ID          int64     `json:"id"`
	Description string    `json:"description"`
	Status      Status    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func NewTask(id int64, description string, status Status) *Task {
	return &Task{
		ID:          id,
		Description: description,
		Status:      status,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}
