package day22

// This is shamelessly stolen from:
// https://godoc.org/container/heap#example-package--PriorityQueue

import (
	"fmt"
)

const (
	neither = 1 << iota
	torch
	gear
)

func tools(terrain int) int {
	switch terrain {
	case rocky:
		return gear | torch
	case wet:
		return gear | neither
	case narrow:
		return neither | torch
	default:
		panic(fmt.Errorf("unknown region terrain: %d", terrain))
	}
}

const movement = 1
const change = 7

// An Item is something we manage in a priority queue.
type Item struct {
	Tile
	priority int // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}
