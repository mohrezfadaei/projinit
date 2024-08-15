package project

import (
	"fmt"
	"os"
	"os/exec"
)

func RunGitInit(projectPath, userName, userEmail string) {
	cmd := exec.Command("git", "init", projectPath)
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error running git init: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Initialized empty Git repository in", projectPath)

	if userName != "" {
		cmd = exec.Command("git", "-C", projectPath, "config", "user.name", userName)
		err = cmd.Run()
		if err != nil {
			fmt.Printf("Error setting git user.name: %v\n", err)
			os.Exit(1)
		}
	}

	if userEmail != "" {
		cmd = exec.Command("git", "-C", projectPath, "config", "user.email", userEmail)
		err = cmd.Run()
		if err != nil {
			fmt.Printf("Error setting git user.email: %v\n", err)
			os.Exit(1)
		}
	}
}
