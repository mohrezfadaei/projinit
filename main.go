package main

import (
	"github.com/mohrezfadaei/projinit/cmd"
	"github.com/mohrezfadaei/projinit/internal/db"
)

func main() {
	db.InitDB("projinit.db")
	db.Migrate()
	cmd.Execute()
}
