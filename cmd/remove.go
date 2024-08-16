package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	removeName string
	removeID   int
)

var removeCmd = &cobra.Command{
	Use:   "remove [license|gitignore]",
	Short: "Remove a LICENSE or .gitignore template from the database by name or ID",
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
	removeCmd.Flags().StringVar(&removeName, "name", "", "Remove by name")
	removeCmd.Flags().IntVar(&removeID, "id", 0, "Remove by ID")
}
