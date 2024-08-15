package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func FetchResource(localPath, url string) ([]byte, error) {
	if _, err := os.Stat(localPath); err == nil {
		return os.ReadFile(localPath)
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error downloading file: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading downloaded file: %v", err)
	}

	err = os.MkdirAll(filepath.Dir(localPath), 0755)
	if err != nil {
		return nil, fmt.Errorf("error creating directories for file: %v", err)
	}

	err = os.WriteFile(localPath, body, 0644)
	if err != nil {
		return nil, fmt.Errorf("error writing file locally: %v", err)
	}

	return body, nil
}
