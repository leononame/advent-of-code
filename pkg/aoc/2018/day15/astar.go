package day15

import (
	"container/heap"

	"gitlab.com/leononame/advent-of-code-2018/pkg/geo"
)

func (t *target) path(from geo.Point, max int) {
	first := Item{pos: from}
	pq := PriorityQueue{&first}
	costs := make(map[geo.Point]int)
	neighbours := t.enemy.cave.neighbours
	path := make(map[geo.Point]geo.Point)

	for len(pq) > 0 {
		current := heap.Pop(&pq).(*Item).pos
		if current == t.pos {
			t.distance = costs[current]
			return
		}
		for _, n := range neighbours(current) {
			cost := costs[current] + 1
			if cost > max {
				continue
			}
			if c, ok := costs[n]; !ok || cost < c {
				costs[n] = cost
				next := Item{pos: n, priority: cost + t.pos.Manhattan(n)}
				heap.Push(&pq, &next)
				path[n] = current
			}
		}
	}
	t.distance = max + 1
}

func reversePath(p geo.Point, path map[geo.Point]geo.Point) []geo.Point {
	ps := []geo.Point{p}
	for p, ok := path[p]; ok; p, ok = path[p] {
		ps = append(ps, p)
	}
	return ps
}
