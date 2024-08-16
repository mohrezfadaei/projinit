package cmd

import (
	"fmt"

	"github.com/mohrezfadaei/projinit/internal/services"
	"github.com/spf13/cobra"
)

var (
	importType string
	importPath string
	lang       string
)

var importCmd = &cobra.Command{
	Use:   "import [license|gitignore]",
	Short: "Import LICENSE or .gitignore from a file or URL",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		resourceType := args[0]
		switch resourceType {
		case "license":
			importLicense()
		case "gitignore":
			importGitignore()
		default:
			fmt.Println("Error: Invalid type. Use 'license' or 'gitignore'.")
		}
	},
}

func init() {
	importCmd.Flags().StringVarP(&importType, "type", "t", "", "Type of license (required for license)")
	importCmd.Flags().StringVarP(&importPath, "path", "p", "", "Path or URL to import from")
	importCmd.Flags().StringVarP(&lang, "language", "l", "", "Language or tool for license (e.g., go, python, react, etc.)")
}

func importLicense() {
	if importType == "" || importPath == "" {
		fmt.Println("Error: Both -type and -path are required for license import.")
		return
	}

	licenseService := services.NewLicenseService()
	err := licenseService.ImportLicense(importType, importPath)
	if err != nil {
		fmt.Printf("Error importing license: %v\n", err)
		return
	}

	fmt.Println("License imported successfully.")
}

func importGitignore() {
	if lang == "" || importPath == "" {
		fmt.Println("Error: Both -language and -path are required for gitignore import.")
		return
	}

	gitignoreService := services.NewGitignoreService()
	err := gitignoreService.ImportGitignore(lang, importPath)
	if err != nil {
		fmt.Printf("Error importing gitignore: %v\n", err)
		return
	}

	fmt.Println("Gitignore imported successfully.")
}
