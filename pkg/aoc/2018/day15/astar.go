package day15

import (
	"container/heap"

	"gitlab.com/leononame/advent-of-code-2018/pkg/geo"
	"gitlab.com/leononame/advent-of-code-2018/pkg/geo/points"
)

func (t *target) path(from geo.Pointer, max int) {
	first := Item{pos: from}
	pq := PriorityQueue{&first}
	costs := make(map[geo.Pointer]int)
	neighbours := t.enemy.cave.neighbours
	path := make(map[geo.Pointer]geo.Pointer)

	for len(pq) > 0 {
		current := heap.Pop(&pq).(*Item).pos
		if points.Equal(current, t.pos) {
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
				next := Item{pos: n, priority: cost + points.Manhattan(t.pos, n)}
				heap.Push(&pq, &next)
				path[n] = current
			}
		}
	}
	t.distance = max + 1
}

func reversePath(p geo.Pointer, path map[geo.Pointer]geo.Pointer) []geo.Pointer {
	ps := []geo.Pointer{p}
	for p, ok := path[p]; ok; p, ok = path[p] {
		ps = append(ps, p)
	}
	return ps
}
