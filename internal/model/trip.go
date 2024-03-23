package model

import "time"

type Trip struct {
	Id, VendorId    string
	Pickup, Dropoff time.Time
	Passengers      int
	Duration        int
}

