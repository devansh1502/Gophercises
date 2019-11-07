package secret

import (
	"path/filepath"
	"testing"

	"github.com/mitchellh/go-homedir"
)

var encodingKey string
var v = File(encodingKey, secretsPath())

func TestGet(t *testing.T) {
	// Go through the error statements again!
	key := "key"
	err := v.Set(key, "New Key")
	if err == nil {
		t.Error("Error Expected!")
	}

	key = "key"
	_, err = v.Get(key)
	if err == nil {
		t.Error("Error Expected!")
	}

	key = ""
	_, err = v.Get(key)
	if err == nil {
		t.Error("Error Expected!")
	}
}

func secretsPath() string {
	home, err := homedir.Dir()
	if err != nil {
		panic(err)
	}
	return filepath.Join(home, ".secrets")
}
