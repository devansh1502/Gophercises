package cobra

import (
	"testing"

	"github.com/spf13/cobra"
)

// TestSetValue it performs unit testing for SetValue function.
func TestSetValue(t *testing.T) {
	var cmd *cobra.Command
	// args := []string{"twitter"}
	// err := setValue(cmd, args)
	// if err != nil {
	// 	t.Errorf("Unexpected Behavior!%v", err)
	// }
	args := []string{"twitter_api_key", "some-value"}
	err := setValue(cmd, args)
	if err != nil {
		t.Errorf("Something went wrong:%v", err)
	}

}

func TestCheckSetValue(t *testing.T) {
	var cmd *cobra.Command
	args := []string{"twitter_api_key", "some-value-1"}
	checkSetValue(cmd, args)
}
