package services

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/mohrezfadaei/projinit/internal/db"
)

type LicenseService struct{}

func NewLicenseService() *LicenseService {
	return &LicenseService{}
}

// GenerateLicenseContent generates the license content by replacing the template variables
func (ls *LicenseService) GenerateLicenseContent(
	licenseType, userName, userEmail string, year int) (string, error) {

	var content string
	err := db.DB.QueryRow("SELECT contnet FROM licenses WHERE type = ?", licenseType).Scan(&content)
	if err != nil {
		return "", fmt.Errorf("error retrieving license content from database: %w", err)
	}

	data := struct {
		Year      int
		UserName  string
		UserEmail string
	}{
		Year:      year,
		UserName:  userName,
		UserEmail: userEmail,
	}

	tmpl, err := template.New("license").Parse(content)
	if err != nil {
		return "", fmt.Errorf("error parsing license template: %w", err)
	}

	var renderedContent bytes.Buffer
	if err := tmpl.Execute(&renderedContent, data); err != nil {
		return "", fmt.Errorf("error executing license template: %w", err)
	}

	return renderedContent.String(), nil
}
