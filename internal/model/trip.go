package model

import "time"

type Trip struct {
	Id         int       `json:"id"`
	VendorId   int       `json:"vendor_id"`
	Pickup     time.Time `json:"pickup_time"`
	Dropoff    time.Time `json:"dropoff_time"`
	Passengers int       `json:"passengers_count"`
	Duration   int       `json:"duration"`
}

type TripRepository interface {
	Find(id int) (Trip, bool)
	Add(trip Trip) error
}
