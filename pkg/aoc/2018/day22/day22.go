package day22

import (
	"container/heap"
	"fmt"

	"gitlab.com/leononame/advent-of-code-2018/pkg/geo"
	"gitlab.com/leononame/advent-of-code-2018/pkg/geo/points"
)

func Run(input []string, c points.Constructor) {
	m := parse(input, c)
	fmt.Println(m)
	fmt.Println("Part 1:", m.calcRisk())
	fmt.Println("Part 2:", part2(TraverseMap{m}, c))
}

func part2(m TraverseMap, pointer points.Constructor) int {
	first := Tile{Region: *m.get(pointer(0, 0)), tool: torch}
	pq := PriorityQueue{&Item{Tile: first}}
	costs := make(map[Tile]int)
	path := make(map[Tile]Tile)
	costs[first] = 0
	i := 0
	for len(pq) > 0 {
		current := heap.Pop(&pq).(*Item).Tile
		if points.Equal(current.Location, m.Target) {
			return costs[current]
		}
		for _, n := range m.searchNeighbours(current) {
			cost := costs[current]
			// If neighbour is target but no torch is equipped, movement cost is really high
			// to prevent pathing
			if points.Equal(n.Location, m.Target) && n.tool != torch {
				cost += 9999
			} else if n.tool != current.tool { // If tools are different, the cost is the change cost
				cost += change
			} else { // Otherwise, add movement cost
				cost += movement
			}

			if c, ok := costs[n]; !ok || cost < c {
				i++
				// fmt.Printf("Step %3d: From %3d,%3d to %3d,%3d.", i,
				// 	current.Location.GetX(),
				// 	current.Location.GetY(),
				// 	n.Location.GetX(),
				// 	n.Location.GetY())
				// fmt.Printf("Cost: %4d, Tool: %d\n", cost, n.tool)
				costs[n] = cost
				next := Item{Tile: n, priority: cost + distance(n.Location, m.Target)}
				heap.Push(&pq, &next)
				path[n] = current
			}
		}
	}
	return 0
}

func reversePath(t Tile, path map[Tile]Tile) []Tile {
	ts := []Tile{t}
	for t, ok := path[t]; ok; t, ok = path[t] {
		ts = append(ts, t)
	}
	return ts
}

func distance(from, to geo.Pointer) int {
	return abs(from.GetX()-to.GetX()) + abs(from.GetY()-to.GetY())
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func parse(input []string, c points.Constructor) *Map {
	var depth, tx, ty int
	fmt.Sscanf(input[0], "depth: %d", &depth)
	fmt.Sscanf(input[1], "target: %d,%d", &tx, &ty)
	m := Map{Depth: depth, Target: points.NewClassic(tx, ty)}
	mx, my := tx+overSize, ty+overSize
	for y := 0; y < my; y++ {
		m.Data = append(m.Data, []*Region{})
		for x := 0; x < mx; x++ {
			p := c(x, y)
			r := Region{Location: p}
			r.Geological = m.calcGeo(p)
			r.Erosion = (m.Depth + r.Geological) % 20183
			r.Terrain = r.Erosion % 3
			m.Data[y] = append(m.Data[y], &r)
		}
	}
	return &m
}
