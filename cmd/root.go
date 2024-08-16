package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "projinit",
	Short: "Project initialization tool",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	// Register subcommands here
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(importCmd)
}
