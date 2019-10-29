package cmd

import (
	"Gophercises/Exercise7/task/db"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// doCmd Marks a command as Completed or deletes a task from the list. Used as "task do" command
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Mark a task as complete",
	Run: checkDoTasks,
}
func checkDoTasks(cmd *cobra.Command,args []string){ 
	_ = doTasks(cmd,args)
}
// Below variables are assinged to function(taking function as variable) to perform mocking!
var allTasksfunc = db.AllTasks
var deleteTaskfunc = db.DeleteTasks
// DoTasks This returns the list with the deleted task.
// (Command: task do 1, where "1" is the task number in the list)
func doTasks(cmd *cobra.Command, args []string) error {
	var ids []int
	for _, arg := range args {
		 id, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Println("Failed to parse argument", arg)
		} else {
			ids = append(ids, id)
		}
	}
	tasks,err := allTasksfunc()
	if err != nil {
		fmt.Println("Something went wrong:", err)
		return err
	}
	for _, id := range ids {
		if id <= 0 || id > len(tasks) {
			fmt.Println("Invalid task number:", id)
			continue
		}
		task := tasks[id-1]
		err := deleteTaskfunc(task.Key)
		if err != nil {
			fmt.Printf("Failed to mark \"%d\" as completed.Error:%s\n", id, err)
			return err
		} else{
			fmt.Printf("Marked \"%d\" as completed\n", id)
		}
	}
	return nil
}

func init() {
	RootCmd.AddCommand(doCmd)
}
