package cmd

import (
	"github.com/spf13/cobra"
	"testing"
	"path/filepath"
	"github.com/mitchellh/go-homedir"
	"Gophercises/Exercise7/task/db"
	"errors"
)

func createTaskN(task string) (int, error){
	return -1, errors.New("error")
}

func TestAddToList(t *testing.T) {
	var cmd *cobra.Command
	temp := createTaskFunc
	createTaskFunc = createTaskN
	args := []string {"task","executed"}
	err := addToList(cmd,args)
	if err == nil{
		t.Errorf("Something went wrong while adding 1:%v",err)
	}
	createTaskFunc = temp
	err = addToList(cmd,args)
	if err != nil{
		t.Log("task executed!")
	}
}

func init() {
	var home, _ = homedir.Dir()
	dbPath := filepath.Join(home, "tasks.db")
	_ = db.Init(dbPath)
}

func TestCheckAddToList(t *testing.T){
	var cmd *cobra.Command
	args := []string{}
	checkAddToList(cmd,args)
}