package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

var Config Configuration

type Configuration struct {
	Licenses   map[string]string `yaml:"licenses"`
	Gitignores map[string]string `yaml:"gitignores"`
}

func LoadConfig(path string) {
	file, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("Error reading config file: %v\n", err)
		os.Exit(1)
	}

	err = yaml.Unmarshal(file, &Config)
	if err != nil {
		fmt.Printf("Error parsing config file: %v\n", err)
		os.Exit(1)
	}
}
