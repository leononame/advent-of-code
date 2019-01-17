package day22

import (
	"container/heap"
	"fmt"
	"time"

	"gitlab.com/leononame/advent-of-code-2018/pkg/aoc"

	"github.com/sirupsen/logrus"

	"gitlab.com/leononame/advent-of-code-2018/pkg/geo"
)

var logger *logrus.Logger

func Run(c *aoc.Config) (result aoc.Result) {
	logger = c.Logger

	t0 := time.Now()
	m := parse(c.Input)
	result.ParseTime = time.Since(t0)

	t1 := time.Now()
	result.Solution1 = m.calcRisk()
	result.Duration1 = time.Since(t1)

	t2 := time.Now()
	result.Solution2 = part2(TraverseMap{m})
	result.Duration2 = time.Since(t2)
	return
}

func part2(m TraverseMap) int {
	first := Tile{Region: *m.get(geo.Point{}), tool: torch}
	pq := PriorityQueue{&Item{Tile: first}}
	costs := make(map[Tile]int)
	path := make(map[Tile]Tile)
	costs[first] = 0
	i := 0
	for len(pq) > 0 {
		current := heap.Pop(&pq).(*Item).Tile
		if current.Location == m.Target {
			return costs[current]
		}
		for _, n := range m.searchNeighbours(current) {
			cost := costs[current]
			// If neighbour is target but no torch is equipped, movement cost is really high
			// to prevent pathing
			if n.Location == m.Target && n.tool != torch {
				cost += 9999
			} else if n.tool != current.tool { // If tools are different, the cost is the change cost
				cost += change
			} else { // Otherwise, add movement cost
				cost += movement
			}

			if c, ok := costs[n]; !ok || cost < c {
				i++
				logger.Debugf("Step %3d: From %3d,%3d to %3d,%3d.", i,
					current.Location.GetX(),
					current.Location.GetY(),
					n.Location.GetX(),
					n.Location.GetY())
				logger.Debugf("Cost: %4d, Tool: %d\n", cost, n.tool)
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

func distance(from, to geo.Point) int {
	return abs(from.GetX()-to.GetX()) + abs(from.GetY()-to.GetY())
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func parse(input []string) *Map {
	var depth, tx, ty int
	fmt.Sscanf(input[0], "depth: %d", &depth)
	fmt.Sscanf(input[1], "target: %d,%d", &tx, &ty)
	m := Map{Depth: depth, Target: geo.Point{tx, ty}}
	mx, my := tx+overSize, ty+overSize
	for y := 0; y < my; y++ {
		m.Data = append(m.Data, []*Region{})
		for x := 0; x < mx; x++ {
			p := geo.Point{x, y}
			r := Region{Location: p}
			r.Geological = m.calcGeo(p)
			r.Erosion = (m.Depth + r.Geological) % 20183
			r.Terrain = r.Erosion % 3
			m.Data[y] = append(m.Data[y], &r)
		}
	}
	return &m
}
