package cmd

import (
	"errors"
	"github.com/spf13/cobra"
	"strconv"
	"task-tracker/internal"
)

func UpdateDescriptionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "update",
		Short: "Update task",
		Long: `Update a task by providing the task ID and the new status
    Example:
    task-tracker update 1 'new description'
    `,
		RunE: func(cmd *cobra.Command, args []string) error {
			return UpdateTaskDescription(args)
		},
	}
}

func UpdateTaskDescription(args []string) error {
	if len(args) != 2 {
		return errors.New("Please provide a task ID and the new description")
	}

	taskID, err := strconv.ParseInt(args[0], 10, 64)
	if err != nil {
		return err
	}

	newDescription := args[1]

	return internal.UpdateTaskDescription(taskID, newDescription)
}

func UpdateTaskStatusDone() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mark-done",
		Short: "Mark a task as done",
		RunE: func(cmd *cobra.Command, args []string) error {
			return UpdateTaskStatus(args, internal.StatusDone)
		},
	}
	return cmd
}

func UpdateTaskStatusInProgress() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mark-in-progress",
		Short: "Mark a task as done",
		RunE: func(cmd *cobra.Command, args []string) error {
			return UpdateTaskStatus(args, internal.StatusInProgress)
		},
	}
	return cmd
}

func UpdateTaskStatusToDo() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mark-todo",
		Short: "Mark a task as done",
		RunE: func(cmd *cobra.Command, args []string) error {
			return UpdateTaskStatus(args, internal.StatusToDo)
		},
	}
	return cmd
}

func UpdateTaskStatus(args []string, status internal.Status) error {
	if len(args) == 0 {
		return errors.New("please provide a task ID")
	}
	taskID, err := strconv.ParseInt(args[0], 10, 64)
	if err != nil {
		return err
	}

	return internal.UpdateTaskStatus(taskID, status)
}
