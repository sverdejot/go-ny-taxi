package storage

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func Init(cs string) *sql.DB {
	db, err := sql.Open("postgres", cs)

	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	return db
}
