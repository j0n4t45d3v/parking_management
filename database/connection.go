package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var connection *sql.DB

const (
	host     = "127.0.0.1"
	port     = 5432
	user     = "postgres"
	password = "senhasecreta"
	dbname   = "postgres"
)

func GetConnection() (*sql.DB, error) {

	if connection == nil {
		psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
		connection, _ = sql.Open("postgres", psqlInfo)
	}

	if err := connection.Ping(); err != nil {
		return nil, err
	}

	return connection, nil
}
