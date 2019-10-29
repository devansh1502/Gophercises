package cmd

import (
	// "Gophercises/Exercise7/task/db"
	"fmt"
	// "os"

	"github.com/spf13/cobra"
)

// listCmd Displays the full list of tasks. Used as "task list" command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all of your tasks.",
	Run: checkDisplayTasks,
}
func checkDisplayTasks(cmd *cobra.Command, args []string){
	_ = displayTasks(cmd,args)
	}
// DisplayTasks returns all the task.(Command to display: task list)
func displayTasks(cmd *cobra.Command, args []string) error {
	tasks, err := allTasksfunc()
	if err != nil {
		fmt.Println("Something went wrong:", err)
		// os.Exit(1)
	}
	if len(tasks) == 0 {
		fmt.Println("You have no tasks to complete!")
		return err
	}
	fmt.Println("You have the following tasks:")
	for i, task := range tasks {
		fmt.Printf("%v. %v\n", i+1, task.Value)
	}
	return err
}

func init() {
	RootCmd.AddCommand(listCmd)
}
