package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"user-management/configs"

	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {
	cfg := configs.LoadConfig()
	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func RunMigrations(db *sql.DB) {
	migrationsDir := "db/migrations"

	files, err := os.ReadDir(migrationsDir)
	if err != nil {
		log.Fatalf("Failed to read migrations directory: %v", err)
	}

	var migrationFiles []string
	for _, file := range files {
		if filepath.Ext(file.Name()) == ".sql" {
			migrationFiles = append(migrationFiles, filepath.Join(migrationsDir, file.Name()))
		}
	}
	sort.Strings(migrationFiles)

	for _, file := range migrationFiles {
		fmt.Printf("Running migration: %s\n", file)

		query, err := os.ReadFile(file)
		if err != nil {
			log.Fatalf("Failed to read migration file %s: %v", file, err)
		}

		_, err = db.Exec(string(query))
		if err != nil {
			log.Fatalf("Failed to execute migration %s: %v", file, err)
		}
	}

	fmt.Println("All migrations completed successfully.")
}