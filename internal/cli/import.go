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
)

func parseTrip(fields []string) model.Trip {
	// eror handling, huh?
	pickup, _ := time.Parse(time.DateTime, fields[2])
	dropoff, _ := time.Parse(time.DateTime, fields[3])
	passengers, _ := strconv.Atoi(fields[4])
	duration, _ := strconv.Atoi(fields[10])

	return model.Trip{
		Id: fields[0],
		VendorId: fields[1],
		Pickup: pickup,
		Dropoff: dropoff,
		Passengers: passengers,
		Duration: duration,
	}
}

var discard model.Trip

func parseFile(fp string) {
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
	for {
		row, err := csvFile.Read()

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf(err.Error())
		}
		
		discard = parseTrip(row)
		i++
	}
	fmt.Println("parsing complete. took: ", time.Since(start), " total: ", i)
}
