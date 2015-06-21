package main

import "math"

const (
	EARTH_RADIUS = 6371
)

// A LinkedList to reorder routes
type Node struct {
	Next  *Node
	Point *Point
}

// TODO: optimize for small distances
func (p *Waypoint) Distance(p2 Waypoint) float64 {
	dLat := (p2.Latitude - p.Latitude) * (math.Pi / 180.0)
	dLon := (p2.Longitude - p.Longitude) * (math.Pi / 180.0)

	lat1 := p.Latitude * (math.Pi / 180.0)
	lat2 := p2.Latitude * (math.Pi / 180.0)

	a1 := (dLat * dLat)
	a2 := (dLon * dLon) * math.Cos(lat1) * math.Cos(lat2)

	a := (a1 + a2) / 4.0

	c := 2.0 * math.Atan2(math.Sqrt(a), math.Sqrt(1.0-a))

	return EARTH_RADIUS * c
}

func (p *Point) Distance(p2 Point) float64 {
	return p.Waypoint.Distance(p2.Waypoint)
}

// TODO: Use interfaces for the points themselves
func GetClosest(points []Point, waypoint Waypoint) (Point, int) {
	var bestDistance = math.MaxFloat64
	var bestPoint Point
	var bestIndex = 0
	for i := 0; i < len(points); i++ {
		cand := points[i]
		var d = cand.Waypoint.Distance(waypoint)
		if d < bestDistance {
			bestDistance = d
			bestPoint = cand
			bestIndex = i
		}
	}

	return bestPoint, bestIndex
}

func (r *Route) TotalDistance() float64 {
	var lastPoint = r.Start
	var totalDistance = 0.0
	r.reorderBySlot()
	for _, p := range r.Points {
		totalDistance += lastPoint.Distance(p.Waypoint)
		lastPoint = p.Waypoint
	}
	return totalDistance
}

// Build distance matrix and initialize
// point indexes
func (rp *RouteProblem) Init(r *Route) {
	r.reorderBySlot()
	pointsLength := len(r.Points)
	rp.Costs = make([][]float64, pointsLength)
	rp.Route = r

	// Build distance matrix
	// Only build the upper diagonal
	for i := 0; i < pointsLength; i++ {
		rp.Costs[i] = make([]float64, pointsLength)
		for j := i; j < pointsLength; j++ {
			if i == j {
				rp.Costs[i][j] = 0.0
			} else {
				rp.Costs[i][j] = r.Points[i].Waypoint.Distance(r.Points[j].Waypoint)
			}
		}
	}

	// Reflect the lower diagonal
	for i := 0; i < pointsLength; i++ {
		for j := i; j < pointsLength; j++ {
			rp.Costs[j][i] = rp.Costs[i][j]
		}
	}
}

func (route *Route) reorderBySlot() {
	pointsLength := len(route.Points)

	// First check if we should reorder
	shouldReorder := false
	for i := 0; i < pointsLength; i++ {
		if route.Points[i].Slot != i {
			shouldReorder = true
			break
		}
	}

	// If the route is ordered then just return
	if !shouldReorder {
		return
	}

	// Iterate over the points. There could be multiple
	// points with the same slot, so we form a linked
	// list of points for each slot
	pointsNodes := make([]*Node, pointsLength)
	for i := 0; i < pointsLength; i++ {
		var point = &(route.Points[i])
		var slot = point.Slot

		// Assign this point to a new node
		var node Node
		node.Point = point

		// If there's already an element for
		// this slot then move it up the list
		if pointsNodes[slot] != nil {
			node.Next = pointsNodes[slot]
		}

		// This new node is the first element of
		// the list
		pointsNodes[slot] = &node
	}

	// Flatten the list
	// Make a new slice for the points
	newPoints := make([]Point, pointsLength)
	index := 0
	for i := 0; i < pointsLength; i++ {
		// If this node is nil just skip it
		if pointsNodes[i] == nil {
			continue
		}
		node := pointsNodes[i]
		for {
			if node == nil {
				break
			}
			// Assign the point to the list
			newPoints[index] = *(node.Point)

			// Correct the slot
			newPoints[index].Slot = index

			// Move to the next element
			node = node.Next

			index++
		}
	}
	route.Points = newPoints
}

func (slice Points) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}
