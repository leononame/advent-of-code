package main

import (
	"fmt"
	"strconv"
	"strings"

	"gitlab.com/leononame/advent-of-code-2018/pkg/util"
)

type coordinate struct {
	x int
	y int
}

func main() {
	fmt.Println("Challenge:\t2018-06")
	input := util.GetInput("input/day06")
	var cs []coordinate
	for _, s := range input {
		cs = append(cs, parseCoordinate(s))
	}
	part1(cs)
}

func part1(cs []coordinate) {
	data := mapCoordinates(cs)
	mx, my := findMax(cs)
	grid := buildGrid(data, mx, my)
	sizes := calcSizes(grid)

	removeBorders(sizes, grid)
	c, size := getMax(sizes)
	fmt.Println("Coordinate", data[c], "has the highest area with a total size of", size)
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

func buildGrid(data map[int]coordinate, sx, sy int) [][]int {
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

func mapCoordinates(cs []coordinate) map[int]coordinate {
	r := make(map[int]coordinate)
	for i, _ := range cs {
		r[i+1] = cs[i]
	}
	return r
}

func getClosest(data map[int]coordinate, x int, y int) int {
	ref := coordinate{x: x, y: y}
	distance := 500
	current := 0
	for i, c := range data {
		d := manhattan(c, ref)
		if d < distance {
			distance = d
			current = i
		} else if d == distance {
			current = 0
		}
	}
	return current

}

func findMax(cs []coordinate) (int, int) {
	x, y := 0, 0
	for _, c := range cs {
		if c.x > x {
			x = c.x
		}
		if c.y > y {
			y = c.y
		}
	}
	return x, y
}

func manhattan(a, b coordinate) int {
	return util.Abs(a.x-b.x) + util.Abs(a.y-b.y)
}

func parseCoordinate(s string) coordinate {
	data := strings.Split(s, ", ")
	return coordinate{
		x: atoi(data[0]),
		y: atoi(data[1]),
	}
}

func atoi(s string) int {
	r, _ := strconv.Atoi(s)
	return r
}
