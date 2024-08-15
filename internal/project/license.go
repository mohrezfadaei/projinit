package project

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/mohrezfadaei/projinit/config"
	"github.com/mohrezfadaei/projinit/internal/utils"
)

func CreateLicenseFile(licenseType, projectPath string, year int, username, email string) {
	licenseFilePath := filepath.Join("public", "licenses", fmt.Sprintf("%s.licenses", licenseType))
	licenseURL, ok := config.Config.Licenses[licenseType]
	if !ok {
		fmt.Printf("Unsupported license type: %s\n", licenseType)
		os.Exit(1)
	}

	body, err := utils.FetchResource(licenseFilePath, licenseURL)
	if err != nil {
		fmt.Printf("Error fetching LICENSE file: %v\n", err)
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
