package cmd

import "github.com/spf13/cobra"

var (
	projectName string
	gitInit     bool
	noLicense   bool
	licenseType string
	noGitIgnore bool
	noReadme    bool
	gitUserName string
	gitEmail    string
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
	initCmd.Flags().StringVar(&gitEmail, "user.email", "", "Git commiter email")
	initCmd.MarkFlagRequired("name")

	rootCmd.AddCommand(initCmd)
}

func initPrject(cmd *cobra.Command, args []string) {
}
