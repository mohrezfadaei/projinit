package config

import (
	"fmt"
	"os"

	_ "embed"

	"gopkg.in/yaml.v2"
)

//go:embed config.yaml
var embeddedConfig []byte
var Config Configuration

type Configuration struct {
	Licenses   map[string]string `yaml:"licenses"`
	Gitignores map[string]string `yaml:"gitignores"`
}

func LoadConfig(path string) {
	err := yaml.Unmarshal(embeddedConfig, &Config)
	if err != nil {
		fmt.Printf("Error parsing config file: %v\n", err)
		os.Exit(1)
	}
}
