package cmd

import (
	"time"

	"github.com/spf13/cobra"
)

var (
	noLicense   bool
	licenseType string
	year        int
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
}
