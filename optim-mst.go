package main

import (
	"fmt"
	"math"
)

type Tree struct {
	Nodes []*Tree
	Value Point
}

// TODO: Implement a more general graph
type Graph struct {
	Elements Points
	Costs    [][]float64
}

func OptimizeRouteMST(routeProblem *RouteProblem) {
	// The root of the tree is the closest to the start
	var route = routeProblem.Route
	_, startIndex := GetClosest(route.Points, route.Start)

	route.Points.Swap(startIndex, len(route.Points)-1)

	graph := BuildGraph(routeProblem)

	// Build tree
	tree := BuildMST(graph)

	// Preorder walk
	points := Flatten(&tree)

	// Update slots
	for i := 0; i < len(points); i++ {
		points[i].Slot = i
	}

	route.Points = points
}

// Build graph for points
func BuildGraph(routeProblem *RouteProblem) Graph {
	// Initialize the graph
	var Graph Graph
	Graph.Elements = routeProblem.Route.Points

	// TODO: Check if this is really necessary
	// or if the matrix is already allocated
	Graph.Costs = routeProblem.Costs

	return Graph
}

// Build minimum spanning tree with start as root
func BuildMST(graph Graph) Tree {
	added := make([]bool, len(graph.Elements))

	// Mark the first element as added
	added[0] = true

	// A hash representation will be used, where each
	// entry has a slice of points connected to the
	// ley point.
	var graphHash map[Point][]*Point
	graphHash = make(map[Point][]*Point)

	// Iterate over every point
	for i := 0; i < len(graph.Elements); i++ {
		// 1. Get the closest non-added element from the last point
		costs := graph.Costs[i]
		minCost := math.MaxFloat64
		minIndex := -1

		// TODO: use smarter data structures for this.
		// It could be as simple as a map with the point as
		// key and the index as value
		for j := 0; j < len(costs); j++ {
			// Only consider added nodes
			if j == i || !added[j] {
				continue
			}
			if costs[j] < minCost {
				minIndex = j
				minCost = costs[j]
			}
		}

		// 2. Add it to the tree
		if minIndex >= 0 {
			// Mark the current node as added
			added[i] = true

			// Get the connecting node
			minPoint := graph.Elements[minIndex]

			// Initalize slice of points for this head node
			if graphHash[minPoint] == nil {
				graphHash[minPoint] = make([]*Point, 0)
			}

			// Update the slice with the added point
			graphHash[minPoint] = append(graphHash[minPoint], &graph.Elements[i])
		} else {
			fmt.Printf("Node %d not added\n", i)
		}
		// 3. Repeat
	}

	// Transform the hash-represented tree to a formal
	// tree.
	// Actually we don't need the tree itself, since
	// the traversal with the hash is trivial.
	return HashToTree(graphHash, graph.Elements[0])
}

// Given the following tree:
//     3
//    / \
//   4   8
//     / | \
//    9 10  11
//
// it could be represented with a map (hash table)
// like this:
//
// 3: 4 - 8
// 8: 9 - 10 - 11
//
// This method converts it to a Tree struct
func HashToTree(hash map[Point][]*Point, root Point) Tree {
	// Build a tree and specify the value
	var tree Tree
	tree.Value = root

	// Iterate over every element of this root
	for i := 0; i < len(hash[root]); i++ {
		newRootPointer := hash[root][i]

		// TODO: The slice shouldn't contain the key-root
		// (second condition of the if statement should
		// always be false)
		if newRootPointer == nil || *newRootPointer == root {
			continue
		}

		// Form a tree with each children
		var newTree = HashToTree(hash, *newRootPointer)

		// Append the formed tree to this element's nodes
		tree.Nodes = append(tree.Nodes, &newTree)
	}
	return tree
}

// Build minimum spanning tree with start as root
func BuildMSTDummy(graph Graph) Tree {
	var tree Tree

	// The first element is the root
	tree.Value = graph.Elements[0]

	lastTree := &tree

	// This is just a toy implementation
	for _, p := range graph.Elements[1:] {
		newTree := new(Tree)
		newTree.Value = p

		lastTree.Nodes = append(lastTree.Nodes, newTree)

		lastTree = newTree
	}
	return tree
}

// Preorder tree walk for a tree
func Flatten(tree *Tree) Points {
	var points Points = Points{tree.Value}
	if len(tree.Nodes) > 0 {
		for i := 0; i < len(tree.Nodes); i++ {
			points = append(points, Flatten(tree.Nodes[i])...)
		}
	}
	return points
}
