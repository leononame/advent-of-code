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

	var grid [300][300]int
	build(&grid)
	mx, my, max := max3x3(grid)

	fmt.Println("The biggest power cell has a power of ", max)
	fmt.Printf("It starts at %d,%d", mx+1, my+1)
}

func max3x3(grid [300][300]int) (maxX, maxY, max int) {
	for y := 0; y < len(grid)-3; y++ {
		for x := 0; x < len(grid[y])-3; x++ {
			s := calc3x3(&grid, x, y)
			if s > max {
				max = s
				maxX = x
				maxY = y
			}
		}
	}
	return
}

func calc3x3(grid *[300][300]int, x, y int) (sum int) {
	for dx := x; dx < x+3; dx++ {
		for dy := y; dy < y+3; dy++ {
			sum += grid[dy][dx]
		}
	}
	return sum
}

func build(grid *[300][300]int) {
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
