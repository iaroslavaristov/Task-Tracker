package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

func RootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "task-tracker",
		Short: "CLI Task Tracker",
		RunE:  func(cmd *cobra.Command, args []string) error {},
	}

	return cmd
}
