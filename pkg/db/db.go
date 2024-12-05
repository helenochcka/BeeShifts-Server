package db

import (
	"BeeShifts-Server/internal/repositories"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDatabase(host string, port int, user string, password string, dbname string) error {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return repositories.ConnErr
	}

	DB = db
	return nil
}
