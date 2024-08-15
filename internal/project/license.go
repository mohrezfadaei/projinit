package project

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"text/template"

	"github.com/mohrezfadaei/projinit/internal/config"
)

func CreateLicenseFile(licenseType, projectPath string, year int, username, email string) {
	licenseURL, ok := config.Config.Licenses[licenseType]
	if !ok {
		fmt.Printf("Unsupported license type: %s\n", licenseType)
		os.Exit(1)
	}

	resp, err := http.Get(licenseURL)
	if err != nil {
		fmt.Printf("Error downloading LICENSE file: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading LICENSE file: %v\n", err)
		os.Exit(1)
	}

	filePath := filepath.Join(projectPath, "LICENSE")
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Printf("Error creating LICENSE file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	tmpl, err := template.New("license").Parse(string(body))
	if err != nil {
		fmt.Printf("Error parsing license template: %v\n", err)
		os.Exit(1)
	}

	data := struct {
		Year     int
		Username string
		Email    string
	}{
		Year:     year,
		Username: username,
		Email:    email,
	}

	err = tmpl.Execute(file, data)
	if err != nil {
		fmt.Printf("Error writing LICENSE file: %v\n", err)
		os.Exit(1)
	}
}
