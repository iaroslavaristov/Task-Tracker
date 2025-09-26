package cmd

import (
	"task-tracker/internal"

	"github.com/spf13/cobra"
)

func ListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List tasks",
		Long: `List all tasks. You can filter tasks by status

    Example:
    task-tracker list todo
    task-tracker list in-progress
    task-tracker list done
    `,
		RunE: func(cmd *cobra.Command, args []string) error {
			return ListTasks(args)
		},
	}
	return cmd
}

func ListTasks(args []string) error {
	if len(args) > 0 {
		status := internal.TaskStatusFromString(args[0])
		return internal.ListTasks(status)
	}

	return internal.ListTasks("all")
}
