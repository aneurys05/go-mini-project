package db

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

func Connect() (*sql.DB, error) {
	conn := os.Getenv("DATABASE_URL")
	if conn == "" {
		conn = "postgres://postgres:postgres@localhost:5432/petpal?sslmode=disable"
	}

	return sql.Open("postgres", conn)
}
