package cobra

import (
	"Gophercises/Exercise7/task/db"
	"path/filepath"
	"testing"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

// TestGetValue it performs unit testing for getValue function.
func TestGetValue(t *testing.T) {
	var cmd *cobra.Command
	args := []string{"twitter"}
	err := getValue(cmd, args)
	if err != nil {
		t.Errorf("Invalid Value:%v", err)
	}
	args = []string{"twitter_api_key"}
	err = getValue(cmd, args)
	if err != nil {
		t.Errorf("Something went wrong:%v", err)
	}
}

func TestCheckGetValue(t *testing.T) {
	var cmd *cobra.Command
	args := []string{"1"}
	checkGetValue(cmd, args)
}
func init() {
	var home, _ = homedir.Dir()
	dbPath := filepath.Join(home, ".secrets")
	_ = db.Init(dbPath)
}
