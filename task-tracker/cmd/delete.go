package cmd

import (
	"errors"
	"strconv"
	"task-tracker/internal"

	"github.com/spf13/cobra"
)

func DeleteCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete task by ID",
		RunE: func(cmd *cobra.Command, args []string) error {
			return DeleteTask(args)
		},
	}
	return cmd
}

func DeleteTask(args []string) error {
	if len(args) == 0 {
		return errors.New("task id is required")
	}

	taskID, err := strconv.ParseInt(args[0], 10, 64)
	if err != nil {
		return err
	}

	return internal.DeleteTask(taskID)
}
