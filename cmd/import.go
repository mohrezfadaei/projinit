package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	importType string
	lang       string
	path       string
)

var importCmd = &cobra.Command{
	Use:   "import [license|gitignore]",
	Short: "Import LICENSE or .gitignore from a file or URL",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		resourceType := args[0]
		switch resourceType {
		case "license":
		case "gitignore":
		default:
			fmt.Println("Invalid type. Use 'license' or 'gitignore'.")
		}
	},
}

func init() {
	importCmd.Flags().StringVarP(&importType, "type", "t", "", "Type of license (required for license)")
	importCmd.Flags().StringVarP(&lang, "language", "l", "", "Language or tool for license (e.g., go, python, react, etc.)")
	importCmd.Flags().StringVarP(&path, "path", "p", "", "Path or URL to import from")
}

func importLicense() {}

func importGitignore() {}
