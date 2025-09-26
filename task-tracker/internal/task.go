package internal

import (
	"fmt"
	"log"
	"time"
)

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

func ListTasks(status Status) error {
	tasks, err := ReadTasksFromFile()
	if err != nil {
		log.Println("Cannot read tasks from file:", err)
		return err
	}

	if len(tasks) == 0 {
		log.Println("No tasks found")
		return nil
	}

	var filteredTasks []Task

	switch status {
	case "all":
		filteredTasks = tasks
	default:
		for _, task := range tasks {
			if task.Status == status {
				filteredTasks = append(filteredTasks, task)
			}
		}
	}

	for _, task := range filteredTasks {
		updatedTimeFormat := task.UpdatedAt.Format("03-05-2025 17:08:09")
		fmt.Printf("%d %s %s (%s)", task.ID, task.Status, task.Description, updatedTimeFormat)
	}
	fmt.Println()

	return nil
}

func TaskStatusFromString(status string) Status {
	switch status {
	case "todo":
		return StatusToDo
	case "done":
		return StatusDone
	case "in-progress":
		return StatusInProgress
	default:
		return StatusNone
	}
}
