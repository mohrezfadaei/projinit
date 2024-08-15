package project

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/mohrezfadaei/projinit/internal/config"
	"github.com/mohrezfadaei/projinit/internal/utils"
)

func CreateGitignoreFile(lang, projectPath string) {
	gitignoreFilePath := filepath.Join("public", "gitignores", fmt.Sprintf("%s.gitignore", lang))
	gitignoreURL, ok := config.Config.Gitignores[lang]
	if !ok {
		fmt.Printf("Unsupported language or tool: %s\n", lang)
		os.Exit(1)
	}

	gitignoreContent, err := utils.FetchResource(gitignoreFilePath, gitignoreURL)
	if err != nil {
		fmt.Printf("Error fetching .gitignore file: %v\n", err)
		os.Exit(1)
	}

	filePath := filepath.Join(projectPath, ".gitignore")
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Printf("Error creating .gitignore file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	_, err = file.Write(gitignoreContent)
	if err != nil {
		fmt.Printf("Error writing to .gitignore file: %v\n", err)
		os.Exit(1)
	}
}
