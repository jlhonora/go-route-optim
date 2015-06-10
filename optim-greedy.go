package main

import (
	"log"
	"math"
)

func OptimizeRouteGreedy(routeProblem *RouteProblem) {
	route := routeProblem.Route
	lastWaypoint := route.Start
	pointsLength := len(route.Points)
	var closestIndex int

	// Get a starting point
	_, closestIndex = GetClosest(route.Points, lastWaypoint)

	pointIndexes := make([]int, pointsLength)

	// Initialize a map for remaining points to be
	// included
	remainingPoints := make(map[int]bool)
	for i := 0; i < pointsLength; i++ {
		remainingPoints[i] = true
	}

	// Exclude the first point and add it to the list
	delete(remainingPoints, closestIndex)
	pointIndexes[0] = closestIndex

	// Holds the count of added points
	count := 1

	// Last added point's index
	lastIndex := closestIndex

	for count < pointsLength {
		bestValue := math.MaxFloat64
		bestIndex := -1
		for key, _ := range remainingPoints {
			var currentCost = routeProblem.Costs[lastIndex][key]
			if currentCost < bestValue {
				bestIndex = key
				bestValue = currentCost
			}
		}

		// If a valid candidate is found then
		// add it
		if bestIndex > -1 {
			delete(remainingPoints, bestIndex)
			pointIndexes[count] = bestIndex
		} else {
			log.Fatal("Bug! index: ", count)
		}
		count++
	}

	// Reorder points
	for i := 0; i < pointsLength; i++ {
		route.Points[pointIndexes[i]].Slot = i
	}
}
