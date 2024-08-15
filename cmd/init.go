package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/mohrezfadaei/projinit/internal/project"
	"github.com/spf13/cobra"
)

var (
	projectName  string
	gitInit      bool
	noLicense    bool
	licenseType  string
	noGitIgnore  bool
	noReadme     bool
	gitUserName  string
	gitUserEmail string
	year         int
)

func init() {
	var initCmd = &cobra.Command{
		Use:   "init [go|python|react|ansible|terraform]",
		Short: "Initialize a new project",
		Args:  cobra.ExactArgs(1),
		Run:   initPrject,
	}

	initCmd.Flags().StringVarP(&projectName, "name", "n", "", "Name of the project (required)")
	initCmd.Flags().BoolVar(&gitInit, "git-init", false, "Initialize a new Git repository")
	initCmd.Flags().BoolVar(&noLicense, "no-license", false, "Don't create LICENSE file")
	initCmd.Flags().StringVar(&licenseType, "license-type", "", "Type of license (apache, mit, etc.)")
	initCmd.Flags().BoolVar(&noGitIgnore, "no-gitignore", false, "Don't create .gitignore file")
	initCmd.Flags().BoolVar(&noReadme, "no-readme", false, "Don't create README file")
	initCmd.Flags().StringVar(&gitUserName, "user.name", "", "Git commiter name")
	initCmd.Flags().StringVar(&gitUserEmail, "user.email", "", "Git commiter email")
	initCmd.Flags().IntVar(&year, "year", time.Now().Year(), "Year for license (default: current year)")
	initCmd.MarkFlagRequired("name")

	rootCmd.AddCommand(initCmd)
}

func initPrject(cmd *cobra.Command, args []string) {
	lang := args[0]

	err := os.Mkdir(projectName, 0755)
	if err != nil {
		fmt.Printf("Error creating project directory: %v\n", err)
		os.Exit(1)
	}

	projectPath := filepath.Join(".", projectName)

	if !noLicense && licenseType != "" {
		project.CreateLicenseFile(licenseType, projectPath, year, gitUserName, gitUserEmail)
	}

	if !noGitIgnore {
		project.CreateGitignoreFile(lang, projectPath)
	}

	if !noReadme {
		project.CreateReadmeFile(projectPath)
	}

	if gitInit {
		project.RunGitInit(projectPath, gitUserName, gitUserEmail)
	}
}
