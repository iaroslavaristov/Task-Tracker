package cmd

import (
	"errors"
	"task-tracker/internal"

	"github.com/spf13/cobra"
)

func AddCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "Add a new task to list",
		RunE: func(cmd *cobra.Command, args []string) error {
			return AddTaskCmd(args)
		},
	}
	return cmd
}

func AddTaskCmd(args []string) error {
	if len(args) == 0 {
		return errors.New("Task description is required")
	}

	return internal.AddTask(args[0])
}
