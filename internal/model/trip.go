package model

import "time"

type Trip struct {
	Id         int       `json:"id"`
	VendorId   int       `json:"vendor_id"`
	Pickup     time.Time `json:"pickup"`
	Dropoff    time.Time `json:"dropoff"`
	Passengers int       `json:"passengers"`
	Duration   int       `json:"duration"`
}

type TripRepository interface {
	Find(id int) (Trip, bool)
	Add(trip Trip) error
}
