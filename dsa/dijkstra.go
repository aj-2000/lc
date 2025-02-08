package dsa

import (
	"container/heap"
)

func Dijkstra(g *Graph, start, end string) ([]string, int) {
	// Min-heap priority queue
	pq := NewMinHeap()
	heap.Init(pq)

	// Distance and previous node tracking
	distances := make(map[string]int)
	prev := make(map[string]string)
	visited := make(map[string]bool)

	// Initialize distances
	g.ForEachVertex(func(v *Vertex) {
		distances[v.GetKey()] = 1<<31 - 1 // Max int (infinity)
	})

	distances[start] = 0

	heap.Push(pq, &Item{Value: start, Priority: 0})

	for pq.Len() > 0 {
		item := heap.Pop(pq).(*Item)
		node := item.Value

		if node == end {
			break
		}

		visited[node] = true
		g.GetVertex(node).ForEachEdge(func(neighbor string, weight int) {
			if !visited[neighbor] {
				newDist := distances[node] + weight
				if newDist < distances[neighbor] {
					distances[neighbor] = newDist
					prev[neighbor] = node
					heap.Push(pq, &Item{Value: neighbor, Priority: newDist})
				}
			}
		})
	}

	// Reconstruct shortest path
	path := []string{}
	for at := end; at != ""; at = prev[at] {
		path = append([]string{at}, path...)
	}

	// If path doesn't contain the start node, no valid path exists
	if len(path) == 0 || path[0] != start {
		return nil, -1
	}

	return path, distances[end]
}
