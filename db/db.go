package db

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

// NewMyStorage initializes a new MySQL storage with the provided configuration.
//
// Parameter: cfg - the MySQL configuration.
// Returns: *sql.DB, error
func NewMyStorage(cfg mysql.Config) (*sql.DB, error) {
	db, err := sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}
