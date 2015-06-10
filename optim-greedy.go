package main

import (
	"log"
	"math"
)

func OptimizeRouteGreedy(routeProblem *RouteProblem) {
	var route = routeProblem.Route
	var lastWaypoint = route.Start
	var pointsLength = len(route.Points)
	var closestIndex int

	// Get a starting point
	_, closestIndex = GetClosest(route.Points, lastWaypoint)

	var pointIndexes = make([]int, pointsLength)

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
	var count = 1

	// Last added point's index
	var lastIndex = closestIndex

	for count < pointsLength {
		var bestValue = math.MaxFloat64
		var bestIndex = -1
		for key, _ := range remainingPoints {
			var currentCost float64
			currentCost = routeProblem.Costs[lastIndex][key]
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
	var newPoints = make([]Point, pointsLength)
	for i := 0; i < pointsLength; i++ {
		newPoints[i] = route.Points[pointIndexes[i]]
		newPoints[i].Slot = i
	}
	route.Points = newPoints
}
