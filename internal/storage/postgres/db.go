package storage

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func Initialize() *sql.DB {
	db, err := sql.Open("postgres", "postgres://user:pass@localhost/nytaxi?sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	return db
}
