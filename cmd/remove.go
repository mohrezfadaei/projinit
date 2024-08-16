package cmd

import (
	"fmt"

	"github.com/mohrezfadaei/projinit/internal/services"
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
			removeLicense()
		case "gitignore":
			removeGitignore()
		default:
			fmt.Println("Error: Invalid type. Use 'license' or 'gitignore'.")
		}
	},
}

func init() {
	removeCmd.Flags().StringVar(&removeName, "name", "", "Remove by name")
	removeCmd.Flags().IntVar(&removeID, "id", 0, "Remove by ID")
}

func removeLicense() {
	if removeID == 0 && removeName == "" {
		fmt.Println("Error: Please provide either --name or --id to remove a license.")
		return
	}

	licenseService := services.NewLicenseService()
	var err error

	if removeID != 0 {
		err = licenseService.RemoveLicenseByID(removeID)
	} else {
		err = licenseService.RemoveLicenseByName(removeName)
	}

	if err != nil {
		fmt.Println("Error removing license:", err)
		return
	}

	fmt.Println("License removed successfully!")
}

func removeGitignore() {
	if removeID == 0 && removeName == "" {
		fmt.Println("Error: Please provide either --name or --id to remove a gitignore.")
		return
	}

	gitignoreService := services.NewGitignoreService()
	var err error

	if removeID != 0 {
		err = gitignoreService.RemoveGitignoreByID(removeID)
	} else {
		err = gitignoreService.RemoveGitignoreByName(removeName)
	}

	if err != nil {
		fmt.Println("Error removing gitignore:", err)
		return
	}

	fmt.Println("Gitignore removed successfully!")
}
