package cmd

import (
	"Gophercises/Exercise7/task/db"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// addCmd adds a task to our list. Used as "task add" command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to your task list",
	Run: checkAddToList,
}
func checkAddToList(cmd *cobra.Command, args []string){ 
	_ = addToList(cmd,args)
 }
// Create function to a variable for mocking(unit testing for generating error for error cases)
var	createTaskFunc = db.CreateTask

func addToList(cmd *cobra.Command, args []string) error {
		task := strings.Join(args, " ")
		_, err := createTaskFunc(task)
		if err != nil {
			fmt.Println("Something went wrong:", err)
			return err
		}
		fmt.Printf("Added \"%s\" to your list.\n", task)
		return nil	
}
func init() {
	RootCmd.AddCommand(addCmd)
}
