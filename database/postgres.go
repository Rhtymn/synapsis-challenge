package database

import (
	"database/sql"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func ConnectPostgresDB(connStr string) (*sql.DB, error) {
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}