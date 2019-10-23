package cmd

import "github.com/spf13/cobra"

// RootCmd It is the variable assigned to our CLI. Used as "task" command
var RootCmd = &cobra.Command{
	Use:   "task",
	Short: "Task is a CLI task manager",
}
