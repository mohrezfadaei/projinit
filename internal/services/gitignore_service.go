package services

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/mohrezfadaei/projinit/internal/db"
)

type GitignoreService struct{}

func NewGitignoreService() *GitignoreService {
	return &GitignoreService{}
}

func (gs *GitignoreService) GenerateGitignoreContent(lang string) (string, error) {
	var content string
	err := db.DB.QueryRow("SELECT content FROM gitignores WHERE language =?", lang).Scan(&content)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("no .gitignore template found for language: %s", lang)
		}
		return "", fmt.Errorf("error retrieving gitignore content from database: %w", err)
	}

	return content, nil
}

func (gs *GitignoreService) ImportGitignore(lang, path string) error {
	content, err := gs.fetchContent(path)
	if err != nil {
		return err
	}

	_, err = db.DB.Exec("INSERT INTO gitignores (lang, content) VALUES (?, ?)", lang, content)
	if err != nil {
		return fmt.Errorf("error inserting gitignore into database: %w", err)
	}
	return nil
}

func (gs *GitignoreService) fetchContent(path string) (string, error) {
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
