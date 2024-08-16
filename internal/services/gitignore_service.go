package services

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/mohrezfadaei/projinit/internal/db"
)

type Gitignore struct {
	ID       int
	Language string
	Content  string
}

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

func (gs *GitignoreService) ListGitignores() ([]db.Gitignore, error) {
	rows, err := db.DB.Query("SELECT id, lang FROM gitignores")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var gitignores []db.Gitignore
	for rows.Next() {
		var gitignore db.Gitignore
		if err := rows.Scan(&gitignore.ID, &gitignore.Language); err != nil {
			return nil, err
		}
		gitignores = append(gitignores, gitignore)
	}

	return gitignores, nil
}

func (gs *GitignoreService) FindGitignoreByID(id int) (*Gitignore, error) {
	var gitignore db.Gitignore
	err := db.DB.QueryRow("SELECT id, lang, content FROM gitignores WHERE id = ?", id).Scan(&gitignore.ID, &gitignore.Language, &gitignore.Content)
	if err != nil {
		return nil, err
	}

	return &Gitignore{
		ID:       gitignore.ID,
		Language: gitignore.Language,
		Content:  gitignore.Content,
	}, nil
}

func (gs *GitignoreService) FindGitignoreByName(name string) (*Gitignore, error) {
	var gitignore db.Gitignore
	err := db.DB.QueryRow("SELECT id, lang, content FROM gitignores WHERE lang = ?", name).Scan(&gitignore.ID, &gitignore.Language, &gitignore.Content)
	if err != nil {
		return nil, err
	}

	// Convert db.Gitignore to services.Gitignore
	return &Gitignore{
		ID:       gitignore.ID,
		Language: gitignore.Language,
		Content:  gitignore.Content,
	}, nil
}

func (gs *GitignoreService) RemoveGitignoreByID(id int) error {
	_, err := db.DB.Exec("DELETE FROM gitignores WHERE id = ?", id)
	return err
}

func (gs *GitignoreService) RemoveGitignoreByName(name string) error {
	_, err := db.DB.Exec("DELETE FROM gitignores WHERE lang = ?", name)
	return err
}
