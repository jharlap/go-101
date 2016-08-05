package main

import "fmt"

func main() {
	type GeoPlace struct {
		Name     string
		Lat, Lon float64
	}

	p := GeoPlace{
		Name: "CN Tower",
		Lat:  43.6426,
		Lon:  -79.3871, // note trailing comma
	}

	p = GeoPlace{"Somewhere in NYC", 40.7484, -73.9857}
	p.Name = "Empire State Building"

	fmt.Printf("%s is near (%.3f, %.3f)\n", p.Name, p.Lat, p.Lon)
	// Empire State Building is near (40.748, -73.986)
}
