package main

import (
	"github.com/mohrezfadaei/projinit/cmd"
	"github.com/mohrezfadaei/projinit/config"
	"github.com/mohrezfadaei/projinit/internal/db"
)

func main() {
	db.InitDB("projinit.db")
	db.Migrate()
	config.LoadConfig("config/config.yaml")
	cmd.Execute()
}
