package main

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"sort"

	_ "modernc.org/sqlite"
)

func main() {
	db, err := sql.Open("sqlite", "hutang.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS schema_migrations (
			version TEXT PRIMARY KEY
		)
	`)
	if err != nil {
		panic(err)
	}

	files, err := filepath.Glob("migrations/*.up.sql")
	if err != nil {
		panic(err)
	}

	sort.Strings(files)

	for _, file := range files {
		version := filepath.Base(file)

		var exists string
		err := db.QueryRow(
			"SELECT version FROM schema_migrations WHERE version = ?",
			version,
		).Scan(&exists)

		if err == nil {
			continue
		}

		sqlBytes, err := os.ReadFile(file)
		if err != nil {
			panic(err)
		}

		fmt.Println("Running migration:", version)

		_, err = db.Exec(string(sqlBytes))
		if err != nil {
			panic(err)
		}

		_, err = db.Exec(
			"INSERT INTO schema_migrations(version) VALUES (?)",
			version,
		)
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("Migration finished successfully")
}
