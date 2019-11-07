package cobra

import (
	"Gophercises/Exercise17/secret"
	"fmt"

	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Sets a secret in your secret storage",
	Run:   checkSetValue,
}

func checkSetValue(cmd *cobra.Command, args []string) {
	_ = setValue(cmd, args)
	// if err != nil {
	// 	fmt.Printf("Error while setting value")
	// }
}

func setValue(cmd *cobra.Command, args []string) error {
	v := secret.File(encodingKey, secretsPath())
	key, value := args[0], args[1]
	err := v.Set(key, value)
	if err != nil {
		fmt.Println("Unable to set value!")
	}
	fmt.Println("Value set successfully!")
	return nil
}

func init() {
	RootCmd.AddCommand(setCmd)
}
