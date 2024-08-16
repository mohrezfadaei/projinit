package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list [license|gitignore|all]",
	Short: "Lists available LICENSE and .gitignore templates",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		resourceType := args[0]

		switch resourceType {
		case "license":
			// not implemented yet
		case "gitignore":
			// not implemented yet
		case "all":
			// not implemented yet
		default:
			fmt.Println("Error: Invalid type. Use 'license', 'gitignore', or 'all'.")
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
