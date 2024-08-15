package project

import (
	"fmt"
	"os"
	"path/filepath"
)

func CreateReadmeFile(projectPath string) {
	filePath := filepath.Join(projectPath, "README.md")
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Printf("Error creating README.md file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("# %s\n", filepath.Base(projectPath)))
	if err != nil {
		fmt.Printf("Error writing to README.md file: %v\n", err)
		os.Exit(1)
	}
}
