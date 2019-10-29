package cmd

import (
	"Gophercises/Exercise7/task/db"
	"github.com/spf13/cobra"
	"testing"
)

func TestDisplayTasks(t *testing.T){
	var cmd *cobra.Command
	// Checking case wherein the task list is empty
	args := []string{"1"}
	doErr := doTasks(cmd,args)
	err := displayTasks(cmd,args)
	if doErr != nil {
		t.Errorf("something is wrong while listing 1:%v",err)
	}
	if err != nil {
		t.Errorf("something is wrong while listing 2:%v",err)
	}
	// Checking wherein the full list is returned.
	args = []string{"task added"}
	addErr := addToList(cmd,args)
	displayErr := displayTasks(cmd,args)
	if addErr != nil {
		t.Errorf("Error while adding task:%v",addErr)
	}else if displayErr != nil {
		t.Errorf("Error while task list:%v",displayErr)
	}
	// Unit Testing with mocking
	// Checking case wherein any error while calling db.AllTasks()
	allTasksfunc = allTaskN
	args = []string{}
	err = displayTasks(cmd,args)
	if err == nil {
		t.Errorf("something went wrong while listing 1:%v",err)
	}
	allTasksfunc = db.AllTasks
}

func TestCheckDisplayTasks(t *testing.T) {
	var cmd *cobra.Command
	args := []string{}
	checkDisplayTasks(cmd,args)
}