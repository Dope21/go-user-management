package repository

import (
	"database/sql"
	"user-management/constants/configs"

	_ "github.com/lib/pq"
)

var db *sql.DB

func ConnectDB() (*sql.DB, error) {
	cfg := configs.LoadConfig()
	var err error
	db, err = sql.Open("postgres", cfg.DBURL)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}