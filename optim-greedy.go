package main

func OptimizeRouteGreedy(route *Route) {
	var optimRoute = Route{}
	optimRoute.Start = route.Start

	var lastPoint = optimRoute.Start

	var count int
	var closestIndex int

	for len(route.Points) > 0 {
		lastPoint, closestIndex = GetClosest(route.Points, lastPoint)
		lastPoint.Slot = count
		count = count + 1
		optimRoute.Points = append(optimRoute.Points, lastPoint)
		// TODO: Delete element
		route.Points = append(route.Points[:closestIndex], route.Points[closestIndex+1:]...)
	}

	route.Points = optimRoute.Points
}
