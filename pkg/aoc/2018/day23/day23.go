package day23

import (
	"fmt"

	"gitlab.com/leononame/advent-of-code-2018/pkg/geo"
)

type nanoBot struct {
	r int
	geo.Point3
}

type nanoBots []nanoBot

func (n nanoBot) inRange(p geo.Point3) bool {
	return n.Manhattan(p) <= n.r
}

func (n nanoBots) points() (points []geo.Point3) {
	points = make([]geo.Point3, len(n))
	for i, b := range n {
		points[i] = b.Point3
		// points = append(points, b.Point3)
	}
	return
}

func Run(input []string) {
	bots := parse(input)
	part1(bots)
	part2(bots)
}

func part1(bots []nanoBot) {
	i := maxRange(bots)
	fmt.Println("Part 1:", strength(bots, i))
}

func part2(bots nanoBots) {
	origin := geo.Point3{}
	var maxCount, bestDistance int
	bestPoint := origin
	// expected: 129293598
	r := geo.FromPointCloud(bots.points())
	// Search tree
	// See: https://www.reddit.com/r/adventofcode/comments/a8s17l/2018_day_23_solutions/ecf450e/https://www.reddit.com/r/adventofcode/comments/a8s17l/2018_day_23_solutions/ecf450e/
	//
	// First iteration is easily done by hand
	size := r.LongestSide() / 2
	r.Min = geo.Point3{
		bestPoint.GetX() - size,
		bestPoint.GetY() - size,
		bestPoint.GetZ() - size}
	r.Max = geo.Point3{
		bestPoint.GetX() + size + 1,
		bestPoint.GetY() + size + 1,
		bestPoint.GetZ() + size + 1}
	for ; size > 0; size /= 2 {
		maxCount = 0
		for x := r.Min.GetX(); x < r.Max.GetX(); x += size {
			for y := r.Min.GetY(); y < r.Max.GetY(); y += size {
				for z := r.Min.GetZ(); z < r.Max.GetZ(); z += size {
					p := geo.Point3{x, y, z}
					// Count the bots in range of the current cube
					// Our reference point (variable p) is the center of the cube
					// We count all bots that touch that cube, hence we increase
					// the search radius by the current cube size
					count := 0
					for _, b := range bots {
						if b.Manhattan(p) < b.r+size {
							count++
						}
					}
					if count > maxCount {
						bestPoint = p
						maxCount = count
						bestDistance = origin.Manhattan(p)
					} else if count == maxCount {
						d := origin.Manhattan(p)
						if d < bestDistance {
							bestDistance = d
							bestPoint = p
						}
					}
				}
			}
		}
		r.Min = geo.Point3{
			bestPoint.GetX() - size,
			bestPoint.GetY() - size,
			bestPoint.GetZ() - size}
		r.Max = geo.Point3{
			bestPoint.GetX() + size + 1,
			bestPoint.GetY() + size + 1,
			bestPoint.GetZ() + size + 1}
		fmt.Printf("Size: %d, Count: %d, Distance: %d, Location: %d,%d,%d\n",
			size, maxCount, bestDistance,
			bestPoint.GetX(),
			bestPoint.GetY(),
			bestPoint.GetZ())
	}
	fmt.Println("Part2:", bestDistance)
}

func strength(bots []nanoBot, i int) int {
	cur := bots[i]
	r := cur.r
	strength := 0
	for _, b := range bots {
		if cur.Manhattan(b.Point3) > r {
			continue
		}
		strength++
	}
	return strength
}

func maxRange(bs []nanoBot) int {
	max := 0
	maxIdx := 0
	for i, b := range bs {
		if b.r > max {
			max = b.r
			maxIdx = i
		}
	}
	return maxIdx
}

func parse(input []string) []nanoBot {
	var bots []nanoBot
	for _, l := range input {
		var x, y, z, r int
		fmt.Sscanf(l, "pos=<%d,%d,%d>, r=%d", &x, &y, &z, &r)
		bots = append(bots, nanoBot{r, geo.Point3{x, y, z}})
	}
	return bots
}
