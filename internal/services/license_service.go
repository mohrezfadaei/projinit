package services

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
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

func (ls *LicenseService) ImportLicense(licenseType, path string) error {
	content, err := ls.fetchContent(path)
	if err != nil {
		return err
	}

	_, err = db.DB.Exec("INSERT INTO licenses (type, content) VALUES (?, ?)", licenseType, content)
	if err != nil {
		return fmt.Errorf("error inserting license into database: %w", err)
	}
	return nil
}

func (ls *LicenseService) fetchContent(path string) (string, error) {
	if path[:4] == "http" {
		resp, err := http.Get(path)
		if err != nil {
			return "", fmt.Errorf("error fetching from URL: %w", err)
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return "", fmt.Errorf("error reading response body: %w", err)
		}
		return string(body), nil
	}

	content, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("error reading file: %w", err)
	}
	return string(content), nil
}

func (ls *LicenseService) ListLicenses() ([]db.License, error) {
	rows, err := db.DB.Query("SELECT id, type FROM licenses")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var licenses []db.License
	for rows.Next() {
		var license db.License
		if err := rows.Scan(&license.ID, &license.Type); err != nil {
			return nil, err
		}
		licenses = append(licenses, license)
	}

	return licenses, nil
}
