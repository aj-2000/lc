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
	for _, v := range []string{"A", "B", "C", "D"} {
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
	}

	for _, edge := range edges {
		graph.AddDirectedEdge(edge.from, edge.to, edge.weight)
	}

	bridges := graph.FindBridges()

	log.Println("Bridges:")
	for _, bridge := range bridges {
		log.Println(bridge)
	}
}
