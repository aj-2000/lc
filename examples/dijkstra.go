package examples

import (
	"log"

	"github.com/aj-2000/lc/dsa"
)

func Dijkstra() {
	log.Println("Running Dijkstra example")
	// Create a graph
	graph := dsa.NewGraph()

	// Add vertices
	for _, v := range []string{"A", "B", "C", "D", "E", "F", "G"} {
		graph.AddVertex(v)
	}

	// Add edges
	edges := []struct {
		from, to string
		weight   int
	}{
		{"A", "B", 1},
		{"A", "C", 2},
		{"B", "D", 3},
		{"C", "D", 1},
		{"D", "E", 1},
		{"E", "F", 1},
		{"E", "G", 2},
		{"F", "G", 1},
	}

	for _, edge := range edges {
		graph.AddUnDirectedEdge(edge.from, edge.to, edge.weight)
	}

	log.Println("Graph created")
	graph.Print()

	// Find shortest path from A to G
	log.Println("Finding shortest path from A to G")
	path, distance := graph.FindShortestPath("A", "G")

	// Print the path and distance
	log.Printf("Path: %v", path)
	log.Printf("Distance: %d", distance)
}
