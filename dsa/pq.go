package dsa

import "container/heap"

// Item represents a node in the priority queue
type Item struct {
	Value    string
	Priority int
}

// PriorityQueue implements a heap interface
type PriorityQueue struct {
	items []*Item
	less  func(i, j int) bool
}

func (pq PriorityQueue) Len() int           { return len(pq.items) }
func (pq PriorityQueue) Less(i, j int) bool { return pq.less(i, j) }
func (pq PriorityQueue) Swap(i, j int)      { pq.items[i], pq.items[j] = pq.items[j], pq.items[i] }

func (pq *PriorityQueue) Push(x any) {
	item := x.(*Item)
	pq.items = append(pq.items, item)
}

func (pq *PriorityQueue) Pop() any {
	old := pq.items
	n := len(old)
	item := old[n-1]
	pq.items = old[:n-1]
	return item
}

func NewMinHeap() *PriorityQueue {
	pq := &PriorityQueue{}
	pq.less = func(i, j int) bool { return pq.items[i].Priority < pq.items[j].Priority }
	heap.Init(pq)
	return pq
}

func NewMaxHeap() *PriorityQueue {
	pq := &PriorityQueue{}
	pq.less = func(i, j int) bool { return pq.items[i].Priority > pq.items[j].Priority }
	heap.Init(pq)
	return pq
}
