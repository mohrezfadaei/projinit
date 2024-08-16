package services

import (
	"fmt"
	"os"
	"os/exec"
)

type ProjectService struct {
	LicenseService   *LicenseService
	GitignoreService *GitignoreService
}

func NewProjectService(ls *LicenseService, gs *GitignoreService) *ProjectService {
	return &ProjectService{
		LicenseService:   ls,
		GitignoreService: gs,
	}
}

func (ps *ProjectService) InitializeProject(
	path, projectName, licenseType string,
	year int, userName, userEmail string,
	lang string,
	noLicense, noGitignore, noReadme, gitInit bool) error {

	if !noReadme {
		readmePath := fmt.Sprintf("%s/README.md", path)
		readmeContent := fmt.Sprintf("# %s\n", projectName)
		if err := os.WriteFile(readmePath, []byte(readmeContent), 0644); err != nil {
			return fmt.Errorf("error creating README.md: %w", err)
		}
	}

	if !noLicense {
		licenseContent, err := ps.LicenseService.GenerateLicenseContent(
			licenseType, userName, userEmail, year)
		if err != nil {
			return fmt.Errorf("error generating license content: %w", err)
		}
		licensePath := fmt.Sprintf("%s/LICENSE", path)
		if err := os.WriteFile(licensePath, []byte(licenseContent), 0644); err != nil {
			return fmt.Errorf("error creating LICENSE: %w", err)
		}
	}

	if !noGitignore {
		gitignoreContent, err := ps.GitignoreService.GenerateGitignoreContent(lang)
		if err != nil {
			return fmt.Errorf("error generating gitignore content: %w", err)
		}
		gitignorePath := fmt.Sprintf("%s/.gitignore", path)
		if err := os.WriteFile(gitignorePath, []byte(gitignoreContent), 0644); err != nil {
			return fmt.Errorf("error creating .gitignore: %w", err)
		}
	}

	if gitInit {
		if err := ps.initializeGitRepo(path, userName, userEmail); err != nil {
			return fmt.Errorf("error initializing Git repository: %w", err)
		}
	}

	return nil
}

func (ps *ProjectService) initializeGitRepo(path, userName, userEmail string) error {
	cmd := exec.Command("git", "init")
	cmd.Dir = path
	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("error initializing git repository: %s: %s", err, string(output))
	}

	cmd = exec.Command("git", "config", "--local", "user.name", userName)
	cmd.Dir = path
	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("error setting git user.name: %s: %s", err, string(output))
	}

	cmd = exec.Command("git", "config", "--local", "user.email", userEmail)
	cmd.Dir = path
	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("error setting git user.email: %s: %s", err, string(output))
	}

	return nil
}
