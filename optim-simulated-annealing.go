package main

import (
	"fmt"
	"math"
	"math/rand"
)

const STARTING_TEMP = 100.0
const COOLING_RATE = 0.0005

func OptimizeRouteSimulatedAnnealing(routeProblem *RouteProblem) {
	route := routeProblem.Route
	lastWaypoint := route.Start
	pointsLength := len(route.Points)

	// Get a starting point
	_, closestIndex := GetClosest(route.Points, lastWaypoint)

	// The closest point will always be the first one
	route.Points.Swap(0, closestIndex)

	var proposedRoute Route

	proposedRoute.Points = make([]Point, pointsLength)
	clonePoints(pointsLength, *route, &proposedRoute)
	proposedRoute.Start = route.Start

	currentDistance := route.TotalDistance()
	currentTemp := STARTING_TEMP
	iter := 0
	for currentTemp > 1 {
		shuffle(pointsLength, &proposedRoute)
		shuffle(pointsLength, &proposedRoute)
		proposedDistance := proposedRoute.TotalDistance()

		// Only for debugging/plotting purposes
		fmt.Println(iter, currentTemp, currentDistance, proposedDistance)

		if shouldAcceptRoute(currentTemp, currentDistance, proposedDistance) {
			currentDistance = proposedDistance
			clonePoints(pointsLength, proposedRoute, route)
		} else {
			// If the route wasn't accepted then go back to the original one
			clonePoints(pointsLength, *route, &proposedRoute)
		}
		currentTemp *= (1 - COOLING_RATE)

		iter += 1
	}
}

// TODO: Only works for routes with 3 or more points
func shuffle(pointsLength int, shuffledRoute *Route) {
	// The first point never moves
	index_a := rand.Intn(pointsLength-1) + 1
	index_b := index_a
	if index_a == 1 {
		index_b = 2
	} else if index_a == (pointsLength - 1) {
		index_b = index_a - 1
	} else {
		delta := -1
		if rand.Intn(2) == 1 {
			delta = 1
		}
		index_b = index_a + delta
	}

	shuffledRoute.Points.Swap(index_a, index_b)
	shuffledRoute.Points[index_a].Slot = index_a
	shuffledRoute.Points[index_b].Slot = index_b
}

func shouldAcceptRoute(temperature float64, currentDistance float64, proposedDistance float64) bool {
	if proposedDistance < currentDistance {
		return true
	}
	return math.Exp(25*(currentDistance-proposedDistance)/temperature) > rand.Float64()
}
