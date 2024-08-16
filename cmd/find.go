package cmd

import (
	"fmt"

	"github.com/mohrezfadaei/projinit/internal/services"

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
			findLicense()
		case "gitignore":
			findGitignore()
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

func findLicense() {
	if findID == 0 && findName == "" {
		fmt.Println("Error: Please provide either --name or --id to find a license.")
		return
	}

	licenseService := services.NewLicenseService()
	var license *services.License
	var err error

	if findID != 0 {
		license, err = licenseService.FindLicenseByID(findID)
	} else {
		license, err = licenseService.FindLicenseByName(findName)
	}

	if err != nil {
		fmt.Println("Error finding license:", err)
		return
	}

	if license != nil {
		fmt.Printf("Found License:\nID: %d\nType: %s\nContent:\n%s\n", license.ID, license.Type, license.Content)
	} else {
		fmt.Println("License not found.")
	}
}

func findGitignore() {
	if findID == 0 && findName == "" {
		fmt.Println("Error: Please provide either --name or --id to find a gitignore.")
		return
	}

	gitignoreService := services.NewGitignoreService()
	var gitignore *services.Gitignore
	var err error

	if findID != 0 {
		gitignore, err = gitignoreService.FindGitignoreByID(findID)
	} else {
		gitignore, err = gitignoreService.FindGitignoreByName(findName)
	}

	if err != nil {
		fmt.Println("Error finding gitignore:", err)
		return
	}

	if gitignore != nil {
		fmt.Printf("Found Gitignore:\nID: %d\nLanguage: %s\nContent:\n%s\n", gitignore.ID, gitignore.Language, gitignore.Content)
	} else {
		fmt.Println("Gitignore not found.")
	}
}
