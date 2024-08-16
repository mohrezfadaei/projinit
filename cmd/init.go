package cmd

import (
	"fmt"
	"time"

	"github.com/mohrezfadaei/projinit/internal/services"
	"github.com/spf13/cobra"
)

var (
	noLicense   bool
	licenseType string
	year        int
	userName    string
	userEmail   string
	noGitignore bool
	gitInit     bool
	projectName string
	noReadme    bool
	//lang        string
)

var initCmd = &cobra.Command{
	Use:   "init [OPTIONS] PATH",
	Short: "Initializes a new project with LICENSE, README.md, and .gitignore",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		path := args[0]

		if gitInit {
			for i := 1; i < len(args); i++ {
				switch args[i] {
				case "user.name":
					if i+1 < len(args) {
						userName = args[i+1]
						i++
					} else {
						fmt.Printf("Error: Missing user.name argument\n")
					}
					return
				case "user.email":
					if i+1 < len(args) {
						userEmail = args[i+1]
						i++
					} else {
						fmt.Printf("Error: Missing user.email argument\n")
						return
					}
				}

			}
		}

		licenseService := services.NewLicenseService()
		gitignoreService := services.NewGitignoreService()
		projectService := services.NewProjectService(licenseService, gitignoreService)

		err := projectService.InitializeProject(
			path,
			projectName,
			licenseType,
			year,
			userName,
			userEmail,
			lang,
			noLicense,
			noGitignore,
			noReadme,
			gitInit,
		)

		if err != nil {
			fmt.Println("Error initializing project:", err)
			return
		}

		fmt.Println("Project initialized successfully!")
	},
}

func init() {
	initCmd.Flags().BoolVar(&noLicense, "no-license", false, "Don't create LICENSE")
	initCmd.Flags().StringVar(&licenseType, "license-type", "", "Type of license")
	initCmd.Flags().IntVar(&year, "year", time.Now().Year(), "Year for the LICENSE file")
	initCmd.Flags().BoolVar(&noGitignore, "no-gitignore", false, "Don't create .gitignore")
	initCmd.Flags().BoolVar(&gitInit, "git-init", false, "Initialize a git repository")
	initCmd.Flags().StringVar(&projectName, "name", "ProjectName", "Project name")
	initCmd.Flags().BoolVar(&noReadme, "no-readme", false, "Don't create README.md")
	initCmd.Flags().StringVar(&lang, "lang", "Go", "Programming language for .gitignore")

	rootCmd.AddCommand(initCmd)
}
