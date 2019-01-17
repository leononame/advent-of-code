package day06

import (
	"strconv"
	"strings"
	"time"

	"github.com/sirupsen/logrus"

	"gitlab.com/leononame/advent-of-code-2018/pkg/geo"

	"gitlab.com/leononame/advent-of-code-2018/pkg/aoc"
)

func Run(c *aoc.Config) (result aoc.Result) {
	logger = c.Logger

	t0 := time.Now()
	var ps []geo.Point
	for _, s := range c.Input {
		ps = append(ps, parseCoordinate(s))
	}
	result.ParseTime = time.Since(t0)

	t1 := time.Now()
	result.Solution1 = part1(ps)
	result.Duration1 = time.Since(t1)

	t2 := time.Now()
	result.Solution2 = part2(ps)
	result.Duration2 = time.Since(t2)
	return

	return
}

var logger *logrus.Logger

func part1(cs []geo.Point) int {
	data := mapCoordinates(cs)
	mx, my := findMax(cs)
	grid := buildGrid(data, mx, my)
	sizes := calcSizes(grid)

	removeBorders(sizes, grid)
	c, size := getMax(sizes)
	logger.Debug("Coordinate", data[c], "has the highest area with a total size of", size)
	return size
}

func part2(cs []geo.Point) int {
	mx, my := findMax(cs)
	count := 0
	for x := 0; x < mx; x++ {
		for y := 0; y < my; y++ {
			if isValid(x, y, cs) {
				count++
			}
		}
	}
	logger.Debug("The area of the safe region (part 2) is", count)
	return count
}

func isValid(x, y int, cs []geo.Point) bool {
	sum := 0
	ref := geo.Point{x, y}
	for _, c := range cs {
		sum += ref.Manhattan(c)
	}
	return sum < 10000
}

func removeBorders(sizes map[int]int, grid [][]int) {
	sx := len(grid)
	sy := len(grid[0])
	for x := range grid {
		delete(sizes, grid[x][0])
		delete(sizes, grid[x][sy-1])
	}
	for y := range grid[0] {
		delete(sizes, grid[0][y])
		delete(sizes, grid[sx-1][y])

	}
}

func getMax(sizes map[int]int) (c, size int) {
	max := 0
	maxID := 0
	for k, v := range sizes {
		if v > max {
			max = v
			maxID = k
		}
	}
	return maxID, max
}

func calcSizes(grid [][]int) map[int]int {
	sizes := make(map[int]int)

	for i, _ := range grid {
		for j, _ := range grid[i] {
			idx := grid[i][j]
			sizes[idx]++
		}
	}
	return sizes
}

func buildGrid(data map[int]geo.Point, sx, sy int) [][]int {
	grid := make([][]int, sx)
	for i, _ := range grid {
		grid[i] = make([]int, sy)
	}

	for i, _ := range grid {
		for j, _ := range grid[i] {
			grid[i][j] = getClosest(data, i, j)
		}
	}
	return grid
}

func mapCoordinates(cs []geo.Point) map[int]geo.Point {
	r := make(map[int]geo.Point)
	for i, _ := range cs {
		r[i+1] = cs[i]
	}
	return r
}

func getClosest(data map[int]geo.Point, x int, y int) int {
	ref := geo.Point{x, y}
	distance := 500
	current := 0
	for i, c := range data {
		d := c.Manhattan(ref)
		if d < distance {
			distance = d
			current = i
		} else if d == distance {
			current = 0
		}
	}
	return current

}

func findMax(cs []geo.Point) (int, int) {
	x, y := 0, 0
	for _, c := range cs {
		if c.X > x {
			x = c.X
		}
		if c.Y > y {
			y = c.Y
		}
	}
	return x, y
}

func parseCoordinate(s string) geo.Point {
	data := strings.Split(s, ", ")
	return geo.Point{
		X: atoi(data[0]),
		Y: atoi(data[1]),
	}
}

func atoi(s string) int {
	r, _ := strconv.Atoi(s)
	return r
}
