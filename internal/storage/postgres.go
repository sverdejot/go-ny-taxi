package storage

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

func Initialize() {
	var err error
	db, err = sql.Open("postgres", "postgres://user:pass@localhost/nytaxi?sslmode=disable")

	db.Query("SELECT 1;")

	if err != nil {
		log.Fatal(err)
	}
}
