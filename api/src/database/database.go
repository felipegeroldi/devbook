package database

import (
	"api/src/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // MySQL Driver
)

func Connect() (*sql.DB, error) {
	conn, err := sql.Open("mysql", config.ConnectionString)
	if err != nil {
		return nil, err
	}

	if err = conn.Ping(); err != nil {
		conn.Close()
		return nil, err
	}

	return conn, nil
}
