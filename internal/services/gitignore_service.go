package services

import (
	"database/sql"
	"fmt"

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
