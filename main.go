package main

import (
	"github.com/mohrezfadaei/projinit/cmd"
	"github.com/mohrezfadaei/projinit/internal/config"
)

func main() {
	config.LoadConfig("config/config.yaml")
	cmd.Execute()
}
