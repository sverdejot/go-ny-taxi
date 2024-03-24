package storage

import (
	"errors"

	"github.com/sverdejot/go-ny-taxi/internal/model"
)

type InMemoryTripRepository struct {
	trips map[int]model.Trip
}

func NewInMemoryTripRepository() *InMemoryTripRepository {
	return &InMemoryTripRepository{}
}

func (r *InMemoryTripRepository) Find(id int) (model.Trip, bool) {
	v, ok := r.trips[id]
	return v, ok
}

func (r *InMemoryTripRepository) Add(trip model.Trip) error {
	if _, found := r.trips[trip.Id]; found {
		return errors.New("existing trip")
	}
	r.trips[trip.Id] = trip
	return nil
}
