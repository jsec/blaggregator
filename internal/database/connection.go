package database

import (
	"database/sql"
	"log"
	"os"
)

func Connect() *sql.DB {
	url := os.Getenv("DB_URL")

	db, err := sql.Open("postgres", url)
	if err != nil {
		log.Fatal(err.Error())
	}

	return db
}
