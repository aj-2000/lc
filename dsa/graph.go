package dsa

import (
	"fmt"
)

type Vertex struct {
	key   string
	edges map[string]int
}

func (v *Vertex) GetKey() string {
	return v.key
}

func (v *Vertex) ForEachEdge(f func(string, int)) {
	for key, weight := range v.edges {
		f(key, weight)
	}
}

type Graph struct {
	vertices map[string]*Vertex
}

func NewGraph() *Graph {
	return &Graph{vertices: make(map[string]*Vertex)}
}

func (g *Graph) GetAllVertex() []*Vertex {
	vertices := make([]*Vertex, 0)
	for _, v := range g.vertices {
		vertices = append(vertices, v)
	}
	return vertices
}

func (g *Graph) AddVertex(key string) {
	if _, ok := g.vertices[key]; ok {
		return
	}
	g.vertices[key] = &Vertex{key, make(map[string]int)}
}

func (g *Graph) AddDirectedEdge(from, to string, weight int) {
	g.vertices[from].edges[to] = weight
}

func (g *Graph) AddUnDirectedEdge(from, to string, weight int) {
	g.AddDirectedEdge(from, to, weight)
	g.AddDirectedEdge(to, from, weight)
}

func (g *Graph) RemoveDirectedEdge(from, to string) {
	delete(g.vertices[from].edges, to)
}

func (g *Graph) RemoveUnDirectedEdge(from, to string) {
	g.RemoveDirectedEdge(from, to)
	g.RemoveDirectedEdge(to, from)
}

func (g *Graph) ForEachVertex(f func(*Vertex)) {
	for _, v := range g.vertices {
		f(v)
	}
}

func (g *Graph) GetVertex(key string) *Vertex {
	return g.vertices[key]
}

func (g *Graph) FindShortestPath(start, end string) ([]string, int) {
	return dijkstra(g, start, end)
}

func (g *Graph) FindBridges() [][]string {
	return tarzans(g)
}

func (g *Graph) Print() {
	graph := ""
	for _, v := range g.vertices {
		graph += fmt.Sprintf("%s -> ", v.GetKey())
		v.ForEachEdge(func(neighbor string, weight int) {
			graph += fmt.Sprintf("%s[%d], ", neighbor, weight)
		})
		graph += "\n"
	}
	fmt.Println(graph)
}
