package cli

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/sverdejot/go-ny-taxi/internal/model"
	"github.com/sverdejot/go-ny-taxi/internal/storage/postgres"
)

func parseTrip(fields []string) model.Trip {
	// eror handling, huh?
	pickup, _ := time.Parse(time.DateTime, fields[2])
	dropoff, _ := time.Parse(time.DateTime, fields[3])
	passengers, _ := strconv.Atoi(fields[4])
	duration, _ := strconv.Atoi(fields[10])
	id, _ := strconv.Atoi(fields[0][2:])
	vendor_id, _ := strconv.Atoi(fields[1])

	return model.Trip{
		Id:         id,
		VendorId:   vendor_id,
		Pickup:     pickup,
		Dropoff:    dropoff,
		Passengers: passengers,
		Duration:   duration,
	}
}

func imprt(cs string, fp string) {
	trips := parseFile(fp)

	db := storage.Initialize(cs)
	defer db.Close()

	repo := storage.NewPostgresTripRepository(db)

	for _, t := range trips {
		if err := repo.Add(t); err != nil {
			log.Fatal(err)
		}
	}
}

func parseFile(fp string) []model.Trip {
	start := time.Now()
	f, err := os.Open(fp)
	defer f.Close()

	if err != nil {
		log.Fatalf(err.Error())
	}

	csvFile := csv.NewReader(f)
	// discard header
	_, _ = csvFile.Read()

	var i int
	trips := []model.Trip{}
	for {
		row, err := csvFile.Read()

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf(err.Error())
		}

		trips = append(trips, parseTrip(row))
		i++
	}

	fmt.Println("parsing complete. took: ", time.Since(start), " total: ", i)
	return trips
}
