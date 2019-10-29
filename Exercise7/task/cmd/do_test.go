package cmd

import (
	"Gophercises/Exercise7/task/db"
	"github.com/spf13/cobra"
	"testing"
	"errors"
)
// negative of alltask func, to check error and perform 100% unit testing
func allTaskN() ([]db.Task, error){
	task:= make([]db.Task,0)
	return task,errors.New("list error")
}
func deleteTaskfuncN(key int) (error){
	return errors.New("delete error")
}
func TestDoTasks(t *testing.T){
	var cmd *cobra.Command
	args := []string{"1","10","testing"}
	err := doTasks(cmd,args)
	if err != nil {
		t.Errorf("something went wrong while delete 1:%v",err)
	}
	// Mocking to perform unit testing for tasks,err := allTasksfunc()
	allTasksfunc = allTaskN
	args = []string{"1","10","testing"}
	err = doTasks(cmd,args)
	if err == nil {
		t.Errorf("something went wrong while delete 2:%v",err)
	}
	// Assigning variable to the actual function
	allTasksfunc = db.AllTasks
	
	//Mocking to perform unit testing for err := deleteTaskfunc(task.Key)
	deleteTaskfunc = deleteTaskfuncN
	args = []string{"1"}
	err = doTasks(cmd,args)
	if err == nil {
		t.Errorf("something went wrong while delete 3:%v",err)
	}
	// Assigning variable to the actual function
	deleteTaskfunc = db.DeleteTasks

}

func TestCheckDoTasks(t *testing.T) {
	var cmd *cobra.Command
	args := []string{}
	checkDoTasks(cmd,args)
}