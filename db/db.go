package db

import (
	"database/sql"
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