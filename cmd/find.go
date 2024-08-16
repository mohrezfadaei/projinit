package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	findName string
	findID   int
)

var findCmd = &cobra.Command{
	Use:   "find [license|gitignore]",
	Short: "Find a specific LICENSE or .gitignore template by name or ID",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		resourceType := args[0]

		switch resourceType {
		case "license":
			// not implemented yet
		case "gitignore":
			// not implemented yet
		default:
			fmt.Println("Error: Invalid type. Use 'license' or 'gitignore'.")
		}
	},
}

func init() {
	findCmd.Flags().StringVar(&findName, "name", "", "Search by name")
	findCmd.Flags().IntVar(&findID, "id", 0, "Search by ID")

	// Register the find command
	rootCmd.AddCommand(findCmd)
}
