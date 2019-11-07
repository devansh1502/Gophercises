package cobra

import (
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "secret",
	Short: "Secret is an API key and other secrets manager",
}
var encodingKey string

func init() {
	RootCmd.PersistentFlags().StringVarP(&encodingKey, "Key", "k", "", "The key to use when encoding and decoding secrets.")
}

var dir = homedir.Dir

func secretsPath() string {
	home, err := dir()
	if err != nil {
		panic(err)
	}
	return filepath.Join(home, ".secrets")
}
