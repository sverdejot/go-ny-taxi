package storage

import (
	"errors"

	"github.com/sverdejot/go-ny-taxi/internal/domain"
)

type inMemoryTripRepository struct {
	trips map[int]domain.Trip
}

func NewInMemoryTripRepository() *inMemoryTripRepository {
	return &inMemoryTripRepository{}
}

func (r *inMemoryTripRepository) Find(id int) (domain.Trip, bool) {
	v, ok := r.trips[id]
	return v, ok
}

func (r *inMemoryTripRepository) Add(trip domain.Trip) error {
	if _, found := r.trips[trip.Id]; found {
		return errors.New("existing trip")
	}
	r.trips[trip.Id] = trip
	return nil
}
