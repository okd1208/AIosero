package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() error {
	var err error
	DB, err = sql.Open("postgres", "host=postgres user=Othello password=password dbname=Othello sslmode=disable")
	return err
}
