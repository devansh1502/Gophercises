package cobra

import (
	"Gophercises/Exercise17/secret"
	"fmt"

	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Gets a secret in your secret storage",
	Run:   checkGetValue,
}

func checkGetValue(cmd *cobra.Command, args []string) {
	_ = getValue(cmd, args)
	// if err != nil {
	// 	fmt.Printf("Error while retrieving value")
	// }

}
func getValue(cmd *cobra.Command, args []string) error {
	v := secret.File(encodingKey, secretsPath())
	key := args[0]
	value, err := v.Get(key)
	if err != nil {
		fmt.Println("no value set")
		return err
	}
	fmt.Printf("%s = %s\n", key, value)
	return nil
}
func init() {
	RootCmd.AddCommand(getCmd)
}
