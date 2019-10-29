package db

import (
	"path/filepath"
	"testing"

	"github.com/mitchellh/go-homedir"
)

func TestInit(t *testing.T) {
	dbPath := ""
	err := Init(dbPath)
	if err == nil {
		t.Errorf("Error is:%v", err)
	}
	var home, _ = homedir.Dir()
	dbPath = filepath.Join(home, "tasks.db")
	err = Init(dbPath)
	if err != nil {
		t.Errorf(":%v", err)
	}
}

func TestCreateTask(t *testing.T) {
	// var home, _ = homedir.Dir()
	// dbPath := filepath.Join(home, "tasks.db")
	// _ = Init(dbPath)
	task := "testing"
	_, err := CreateTask(task)
	if err != nil {
		t.Errorf(":%v", err)
	}
}

func TestAllTasks(t *testing.T) {
	_, err := AllTasks()
	if err != nil {
		t.Errorf(":%v", err)
	}
}

func TestDeleteTasks(t *testing.T) {
	// var home, _ = homedir.Dir()
	// dbPath := filepath.Join(home, "tasks.db")
	// _ = Init(dbPath)
	tasks, err := AllTasks()
	var key int
	for _, task := range tasks {
		if task.Value == "testing" {
			key = task.Key
			break
		}else {
			t.Log("Value might be different or task doesn't exist")
		}
	}
	if err != nil {
		t.Errorf("Something went Wrong:%v", err)
	}
	err1 := DeleteTasks(key)
	t.Log("task found and deleted")
	if err1 != nil || key == 0{
		t.Errorf("Error is:%v", err)
	} else {
		t.Log("task Deleted")
		AllTasks()
	}
}
