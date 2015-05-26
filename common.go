package main

import "math"

const (
	EARTH_RADIUS = 6371
)

// TODO: optimize for small distances
func (p *Waypoint) Distance(p2 Waypoint) float64 {
	dLat := (p2.Latitude - p.Latitude) * (math.Pi / 180.0)
	dLon := (p2.Longitude - p.Longitude) * (math.Pi / 180.0)

	lat1 := p.Latitude * (math.Pi / 180.0)
	lat2 := p2.Latitude * (math.Pi / 180.0)

	a1 := (dLat * dLat)
	a2 := (dLon * dLon) * math.Cos(lat1) * math.Cos(lat2)

	a := (a1 + a2) / 4

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return EARTH_RADIUS * c
}

func (p *Point) Distance(p2 Point) float64 {
	return p.Waypoint.Distance(p2.Waypoint)
}

// TODO: Use interfaces for the points themselves
func GetClosest(points []Point, point Point) (Point, int) {
	var bestDistance = math.MaxFloat64
	var bestPoint Point
	var bestIndex = 0
	for idx, cand := range points {
		var d = cand.Distance(point)
		if d < bestDistance {
			bestDistance = d
			bestPoint = cand
			bestIndex = idx
		}
	}

	return bestPoint, bestIndex
}

func (r *Route) TotalDistance() float64 {
	var lastPoint = r.Start
	var totalDistance = 0.0
	for _, p := range r.Points {
		totalDistance += lastPoint.Distance(p)
		lastPoint = p
	}
	return totalDistance
}

// Build distance matrix and initialize
// point indexes
func (r *Route) Init() {
	pointsLength := len(r.Points)
	r.PointIndexes = make([]int, pointsLength)
	r.Costs = make([][]float64, pointsLength)

	// Build distance matrix
	// Only build the upper diagonal
	for i := 0; i < pointsLength; i++ {
		r.Costs[i] = make([]float64, pointsLength)
		r.PointIndexes[i] = r.Points[i].Slot
		for j := i; j < pointsLength; j++ {
			if i == j {
				r.Costs[i][j] = 0.0
			} else {
				r.Costs[i][j] = r.Points[i].Distance(r.Points[j])
			}
		}
	}

	// Reflect the lower diagonal
	for i := 0; i < pointsLength; i++ {
		for j := i; j < pointsLength; j++ {
			r.Costs[j][i] = r.Costs[i][j]
		}
	}
}

func (slice Points) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}
