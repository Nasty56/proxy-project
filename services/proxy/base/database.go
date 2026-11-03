package base

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func ConnectToDb() (*sql.DB, error) {
	databaseHost := GetEnv("DATABASE_HOST", "localhost")
	databaseUser := GetEnv("DATABASE_USER", "student")
	connection := fmt.Sprintf("postgresql://%s:qwerty@%s:5432/proxy_db?sslmode=disable", databaseUser, databaseHost)
	var db *sql.DB
	var err error
	db, err = sql.Open("postgres", connection)
	if err != nil {
		return nil, err
	}
	return db, nil
}