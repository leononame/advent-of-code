package main

import (
	"fmt"
	"strconv"

	"gitlab.com/leononame/advent-of-code-2018/pkg/util"
	"gitlab.com/leononame/advent-of-code-2018/pkg/version"
)

var id int

func main() {
	fmt.Println("Advent of Code 2018, ", version.Str)
	fmt.Println("Challenge: 2018-11")
	id, _ = strconv.Atoi(util.GetInput("input/day11")[0])

	var grid [][]int
	grid = make([][]int, 300)
	for i := range grid {
		grid[i] = make([]int, 300)
	}
	build(grid)
	mx, my, max := max(grid, 3)

	fmt.Println("The biggest power cell has a power of ", max)
	fmt.Printf("It starts at %d,%d\n", mx+1, my+1)

	p, x, y, s := iterSizes(grid)
	fmt.Println("For any size, the most power is", p)
	fmt.Printf("It had the identifier %d,%d,%d\n", x+1, y+1, s)
}

func iterSizes(grid [][]int) (power, maxX, maxY, size int) {
	for s := 1; s < 14; s++ {
		for y := 0; y < len(grid)-s; y++ {
			for x := 0; x < len(grid[y])-s; x++ {
				p := calc(grid, x, y, s)
				if p > power {
					power = p
					maxX = x
					maxY = y
					size = s
				}
			}
		}
	}
	return
}

func max(grid [][]int, size int) (maxX, maxY, max int) {
	for y := 0; y < len(grid)-size; y++ {
		for x := 0; x < len(grid[y])-size; x++ {
			s := calc(grid, x, y, size)
			if s > max {
				max = s
				maxX = x
				maxY = y
			}
		}
	}
	return
}

func calc(grid [][]int, x, y, size int) (sum int) {
	for dx := x; dx < x+size; dx++ {
		for dy := y; dy < y+size; dy++ {
			sum += grid[dy][dx]
		}
	}
	return
}

func build(grid [][]int) {
	for y := range grid {
		for x := range grid[y] {
			grid[y][x] = calcPower(x, y)
		}
	}
}

func calcPower(x, y int) int {
	rid := (x + 1) + 10 // Our x is 0 based, their x is 1 based
	p := rid * (y + 1)
	p += id
	p *= rid
	p = (p / 100) % 10
	return p - 5
}
