package project

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/mohrezfadaei/projinit/internal/config"
)

func CreateGitignoreFile(lang, projectPath string) {
	gitignoreURL, ok := config.Config.Gitignores[lang]
	if !ok {
		fmt.Printf("Unsupported language or tool: %s\n", lang)
		os.Exit(1)
	}

	resp, err := http.Get(gitignoreURL)
	if err != nil {
		fmt.Printf("Error downloading .gitignore file: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading .gitignore file: %v\n", err)
		os.Exit(1)
	}

	filePath := filepath.Join(projectPath, ".gitignore")
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Printf("Error creating .gitignore file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	_, err = file.Write(body)
	if err != nil {
		fmt.Printf("Error writing to .gitignore file: %v\n", err)
		os.Exit(1)
	}
}
