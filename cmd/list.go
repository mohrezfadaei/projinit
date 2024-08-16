package cmd

import (
	"fmt"

	"github.com/mohrezfadaei/projinit/internal/services"
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
			listLicenses()
		case "gitignore":
			listGitignores()
		case "all":
			listLicenses()
			listGitignores()
		default:
			fmt.Println("Error: Invalid type. Use 'license', 'gitignore', or 'all'.")
		}
	},
}

func listLicenses() {
	licenseService := services.NewLicenseService()
	licenses, err := licenseService.ListLicenses()
	if err != nil {
		fmt.Println("Error listing licenses:", err)
		return
	}

	if len(licenses) == 0 {
		fmt.Println("No licenses found.")
		return
	}

	fmt.Println("Available Licenses:")
	for _, license := range licenses {
		fmt.Printf("- %s (ID: %d)\n", license.Type, license.ID)
	}
}

func listGitignores() {
	gitignoreService := services.NewGitignoreService()
	gitignores, err := gitignoreService.ListGitignores()
	if err != nil {
		fmt.Println("Error listing gitignores:", err)
		return
	}

	if len(gitignores) == 0 {
		fmt.Println("No gitignore templates found.")
		return
	}

	fmt.Println("Available .gitignore Templates:")
	for _, gitignore := range gitignores {
		fmt.Printf("- %s (ID: %d)\n", gitignore.Language, gitignore.ID)
	}
}
