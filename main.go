package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	geo "github.com/kellydunn/golang-geo"
)

const (
	CSV = "./CanadianPostalCodes.csv"
)

type Location struct {
	Lat  float64 `json:"latt,string"`
	Long float64 `json:"longt,string"`
}

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

func main() {
	// load command line flags into pointers
	POSTALCODE := flag.String("p", "M4W2L4", "6 character Canadian Postal Code to use as center of search circumference")
	radius := flag.Float64("r", 3, "Distance in KM to set the search radius")
	flag.Parse()

	// build api request to geocoder.ca to get lat/long
	uri := "http://geocoder.ca/?locate=" + *POSTALCODE + "&json=1"

	// get lat/long from geocoder
	resp, err := http.Get(uri)
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	var location Location
	err = json.Unmarshal(body, &location)
	if err != nil {
		log.Fatal(err)
	}
	// ...
	// Create slice to store all relevant postal codes
	var codes [][]string
	// Load CSV into program
	r := LoadCsv(CSV)

	// Read all lines from CSV
	lines, err := r.ReadAll()
	if err != nil {
		log.Fatalf("error reading all lines: %v", err)
	}

	// Load starting Lat/Long into Point struct
	point := geo.NewPoint(location.Lat, location.Long)
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
		if dist <= *radius {
			codes = append(codes, line)
		}
	}
	// create a csv to store codes as output
	file, err := os.Create("results.csv")
	if err != nil {
		log.Fatalln(err)
	}
	w := csv.NewWriter(file)
	w.WriteAll(codes) // calls Flush internally
	if err := w.Error(); err != nil {
		log.Fatalln("error writing csv:", err)
	}
	// for _, code := range codes {
	// 	fmt.Println(code)
	// }
}
