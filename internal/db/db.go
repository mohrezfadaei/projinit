package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB(filepath string) {
	var err error
	DB, err = sql.Open("sqlite3", filepath)
	if err != nil {
		panic(err)
	}
}

func Migrate() {
	createLicenseTable := `
	CREATE TABLE IF NOT EXISTS licenses (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		type TEXT,
		content TEXT
	);`

	createGitignoreTable := `
	CREATE TABLE IF NOT EXISTS gitignores (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		lang TEXT,
		content TEXT
	);`

	if _, err := DB.Exec(createLicenseTable); err != nil {
		panic(err)
	}

	if _, err := DB.Exec(createGitignoreTable); err != nil {
		panic(err)
	}
}
