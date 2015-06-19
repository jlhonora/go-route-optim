package main

import (
	"math"
	"math/rand"
)

const STARTING_TEMP = 500.0
const COOLING_RATE = 0.003

func OptimizeRouteSimulatedAnnealing(routeProblem *RouteProblem) {
	route := routeProblem.Route
	lastWaypoint := route.Start

	// Get a starting point
	closestPoint, closestIndex := GetClosest(route.Points, lastWaypoint)

	var proposedRoute Route

	// Exclude the closest point from the
	// calculation
	route.Points = append(route.Points[:closestIndex],
		route.Points[closestIndex+1:]...)
	pointsLength := len(route.Points)

	currentTemp := STARTING_TEMP
	for currentTemp > 1 {
		shuffle(pointsLength, route, &proposedRoute)

		currentDistance := route.TotalDistance()
		proposedDistance := proposedRoute.TotalDistance()

		if shouldAcceptRoute(currentTemp, currentDistance, proposedDistance) {
			route.Points = proposedRoute.Points
		}
		currentTemp *= (1 - COOLING_RATE)
	}

	// Add the first point back
	route.Points = append([]Point{closestPoint}, route.Points...)
	for i := 0; i < len(route.Points); i++ {
		route.Points[i].Slot = i
	}
}

func shuffle(pointsLength int, route *Route, shuffledRoute *Route) {
	shuffledRoute.Points = route.Points
	index_a := rand.Intn(pointsLength)
	index_b := rand.Intn(pointsLength)

	shuffledRoute.Points[index_a], shuffledRoute.Points[index_b] =
		shuffledRoute.Points[index_b], shuffledRoute.Points[index_a]
}

func shouldAcceptRoute(temperature float64, currentDistance float64, proposedDistance float64) bool {
	if proposedDistance < currentDistance {
		return true
	}
	return math.Exp((currentDistance-proposedDistance)/temperature) > rand.Float64()
}
