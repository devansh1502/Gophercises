package cobra

import (
	"errors"
	"testing"

	homedir "github.com/mitchellh/go-homedir"
)

func dirN() (string, error) {
	return "", errors.New("error")

}
func TestSecretsPath(t *testing.T) {
	dir = dirN

	defer func() {
		if r := recover(); r != nil {
			t.Log("No issues")
			dir = homedir.Dir
		}
	}()

	_ = secretsPath()
	t.Log("Unexpected Behavior")

}
