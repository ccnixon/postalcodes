package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	geo "github.com/kellydunn/golang-geo"
)

const (
	POSTALCODE   = "M4W 2L4"
	CSV          = "./CanadianPostalCodes.csv"
	LAT          = 43.677689
	LONG         = -79.390144
	MAX_DISTANCE = 3
)

func LoadCsv(p string) *csv.Reader {
	f, err := os.Open(p)
	if err != nil {
		log.Fatal(err)
	}
	r := csv.NewReader(bufio.NewReader(f))
	return r
}

// Prints n rows of csv. Useful for debugging.
func debug(i int, r *csv.Reader) {
	row := 0
	for row < i {
		record, err := r.Read()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(record)
		row++
	}
}

// func GetDistance()

func main() {
	// Create slice to store all relevant postal codes
	var codes [][]string
	// Load CSV into program
	r := LoadCsv(CSV)
	// debug(10, r)

	// Read all lines from CSV
	lines, err := r.ReadAll()
	if err != nil {
		log.Fatalf("error reading all lines: %v", err)
	}

	// Load starting Lat/Long into Point struct
	point := geo.NewPoint(LAT, LONG)

	for i, line := range lines {
		if i == 0 {
			// skip header line
			continue
		}
		lat, err := strconv.ParseFloat(line[2], 64)
		if err != nil {
			log.Fatal(err)
		}
		long, err := strconv.ParseFloat(line[3], 64)
		if err != nil {
			log.Fatal(err)
		}
		p2 := geo.NewPoint(lat, long)
		dist := point.GreatCircleDistance(p2)
		if dist <= MAX_DISTANCE {
			codes = append(codes, line)
		}
	}
	for _, code := range codes {
		fmt.Println(code)
	}
}
