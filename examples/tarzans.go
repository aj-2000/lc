package examples

import (
	"log"

	"github.com/aj-2000/lc/dsa"
)

func Tarzans() {
	log.Println("Running Tarzans example")

	// Create a graph
	graph := dsa.NewGraph()

	// Add vertices
	for _, v := range []string{"A", "B", "C", "D", "E"} {
		graph.AddVertex(v)
	}

	// Add edges
	edges := []struct {
		from, to string
		weight   int
	}{
		{"A", "B", 1},
		{"B", "C", 2},
		{"C", "A", 3},
		{"A", "D", 1},
		{"B", "E", 2},
	}

	for _, edge := range edges {
		graph.AddUnDirectedEdge(edge.from, edge.to, edge.weight)
	}

	// Print the graph
	log.Println("Graph created")
	graph.Print()

	bridges := graph.FindBridges()
	log.Println("Bridges:", bridges)
}
