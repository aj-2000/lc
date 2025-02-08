package main

import (
	"container/heap"
	"strings"
)

type Item struct {
	Value    string
	Priority int
	Path     string
}

type PriorityQueue []Item

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].Priority < pq[j].Priority }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x any) {
	*pq = append(*pq, x.(Item))
}

func (pq *PriorityQueue) Pop() any {
	n := len(*pq)
	item := (*pq)[n-1]
	*pq = (*pq)[:n-1]
	return item
}

type Vertex struct {
	key   string
	edges map[string]int
}

type Graph struct {
	vertices map[string]*Vertex
}

func (g *Graph) AddVertex(key string) {
	g.vertices[key] = &Vertex{key, make(map[string]int)}
}

func (g *Graph) AddEdge(from, to string, distance int) {
	g.vertices[from].edges[to] = distance
	g.vertices[to].edges[from] = distance
}

func NewGraph() *Graph {
	return &Graph{
		vertices: make(map[string]*Vertex),
	}
}

func (g *Graph) Dijkstra(start, end string) (string, int) {
	pq := &PriorityQueue{}
	heap.Init(pq)

	heap.Push(pq, Item{start, 0, start})

	for pq.Len() > 0 {
		item := heap.Pop(pq).(Item)
		if item.Value == end {
			return item.Path, item.Priority
		}

		path := item.Path
		distance := item.Priority

		neighbours := g.vertices[item.Value]
		for vertex, edge := range neighbours.edges {
			if !strings.Contains(path, vertex) {
				heap.Push(pq, Item{vertex, distance + edge, item.Path + vertex})
			}
		}
	}

	return "", -1
}

func main() {
	// Create a graph
	graph := NewGraph()

	// Add vertices
	graph.AddVertex("A")
	graph.AddVertex("B")
	graph.AddVertex("C")
	graph.AddVertex("D")
	graph.AddVertex("E")
	graph.AddVertex("F")
	graph.AddVertex("G")

	// Add edges
	graph.AddEdge("A", "B", 1)
	graph.AddEdge("A", "C", 2)
	graph.AddEdge("B", "D", 3)
	graph.AddEdge("C", "D", 1)
	graph.AddEdge("D", "E", 1)
	graph.AddEdge("E", "F", 1)
	graph.AddEdge("E", "G", 2)
	graph.AddEdge("F", "G", 1)

	// Find shortest path from A to G
	path, distance := graph.Dijkstra("A", "G")

	// Print the path and distance
	println("Path:", path)
	println("Distance:", distance)
}
