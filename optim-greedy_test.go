package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestOptimGreedy(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	route := getMockRoute()
	shufflePoints(route.Points)

	fmt.Println("Route: ", route)

	var routeProblem RouteProblem
	routeProblem.Init(&route)
	OptimizeRouteGreedy(&routeProblem)
	fmt.Println("Optimized Route: ", route)

	routeOrdered := true
	for i := 0; i < len(route.Points); i++ {
		if route.Points[i].Slot != route.Points[i].Id {
			routeOrdered = false
			break
		}
	}

	if !routeOrdered {
		t.Error("Optimization failed, check the order")
	}
}

func getMockRoute() Route {
	var route Route
	route.Start = Waypoint{Latitude: -33.4640, Longitude: -70.7341}

	var waypoints []Waypoint
	waypoints = append(waypoints, Waypoint{Latitude: -33.4208, Longitude: -70.9641})
	waypoints = append(waypoints, Waypoint{Latitude: -33.4041, Longitude: -71.0578})
	waypoints = append(waypoints, Waypoint{Latitude: -33.4023, Longitude: -71.1206})
	waypoints = append(waypoints, Waypoint{Latitude: -33.4056, Longitude: -71.1371})

	for i := 0; i < len(waypoints); i++ {
		route.Points = append(route.Points, Point{Slot: rand.Intn(len(waypoints)), Waypoint: waypoints[i], Id: i})
	}

	return route
}

func shufflePoints(points Points) {
	n := len(points)
	for i := n - 1; i > 0; i-- {
		j := rand.Intn(i)
		points.Swap(i, j)
	}
}
