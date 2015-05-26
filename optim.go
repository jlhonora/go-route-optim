package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Waypoint struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Point struct {
	Id       int      `json:"id"`
	Waypoint Waypoint `json:"waypoint"`
	Slot     int      `json:"slot"`
}

// This enables us to declare methods for
// points as a slice
type Points []Point

type Route struct {
	Points       Points `json:"points"`
	Start        Point  `json:"start"`
	Costs        [][]float64
	PointIndexes []int
}

func OptimizeHandler(w http.ResponseWriter, request *http.Request) {
	var route = DecodeRoute(request)
	fmt.Println("Route: ", route)

	OptimizeRoute(&route)
	EncodeRoute(route, w)
}

func DecodeRoute(request *http.Request) Route {
	decoder := json.NewDecoder(request.Body)
	var route Route
	err := decoder.Decode(&route)
	if err != nil {
		fmt.Println("Error decoding route")
		log.Fatal(err)
	}

	return route
}

func EncodeRoute(route Route, w http.ResponseWriter) {
	encoder := json.NewEncoder(w)
	err := encoder.Encode(route)

	if err != nil {
		log.Fatal(err)
	}
}

func OptimizeRoute(route *Route) {
	fmt.Printf("Optimizing route with %d elements\n", len(route.Points))
	fmt.Printf("Starting distance: %f\n", route.TotalDistance())
	route.Init()
	OptimizeRouteMST(route)
	fmt.Println("Route: ", route)
	fmt.Printf("Final distance: %f\n", route.TotalDistance())
	fmt.Println("Done")
}
