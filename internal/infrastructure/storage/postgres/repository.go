package postgres

import (
	"database/sql"
	"log"

	"github.com/sverdejot/go-ny-taxi/internal/domain"
)

var findQuery string = `SELECT * FROM trips WHERE id = $1;`
var insertQuery string = `INSERT INTO trips VALUES ($1, $2, $3, $4, $5, $6);`

type PostgresTripRepository struct {
	db *sql.DB
}

func NewPostgresTripRepository(db *sql.DB) *PostgresTripRepository {
	return &PostgresTripRepository{db}
}

func (r *PostgresTripRepository) Find(id int) (trip domain.Trip, found bool) {
	row := r.db.QueryRow(findQuery, id)

	switch err := row.Scan(&trip.Id, &trip.VendorId, &trip.Pickup, &trip.Dropoff, &trip.Passengers, &trip.Duration); {
	case err == nil:
		return trip, true
	case err == sql.ErrNoRows:
		return
	default:
		log.Fatal(err)
		return
	}
}

func (r *PostgresTripRepository) Add(trip domain.Trip) error {
	_, err := r.db.Exec(insertQuery, trip.Id, trip.VendorId, trip.Pickup, trip.Dropoff, trip.Passengers, trip.Duration)

	return err
}
