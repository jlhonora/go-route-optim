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
	_, closestIndex := GetClosest(route.Points, lastWaypoint)

	var proposedRoute Route

	// The closest point will always be the first one
	route.Points[closestIndex], route.Points[0] = route.Points[0], route.Points[closestIndex]
	pointsLength := len(route.Points)

	proposedRoute.Points = route.Points

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

	// Assign slots by order
	for i := 0; i < len(route.Points); i++ {
		route.Points[i].Slot = i
	}
}

func shuffle(pointsLength int, route *Route, shuffledRoute *Route) {
	shuffledRoute.Points = route.Points
	// The first point never moves
	index_a := rand.Intn(pointsLength-1) + 1
	index_b := rand.Intn(pointsLength-1) + 1

	shuffledRoute.Points[index_a], shuffledRoute.Points[index_b] = shuffledRoute.Points[index_b], shuffledRoute.Points[index_a]
	shuffledRoute.Points[index_a].Slot, shuffledRoute.Points[index_b].Slot = shuffledRoute.Points[index_b].Slot, shuffledRoute.Points[index_a].Slot
}

func shouldAcceptRoute(temperature float64, currentDistance float64, proposedDistance float64) bool {
	if proposedDistance < currentDistance {
		return true
	}
	return math.Exp((currentDistance-proposedDistance)/temperature) > rand.Float64()
}
