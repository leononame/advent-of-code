package day11

import (
	"fmt"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
	"gitlab.com/leononame/advent-of-code-2018/pkg/aoc"
)

var logger *logrus.Logger

func Run(c *aoc.Config) (result aoc.Result) {
	logger = c.Logger

	t0 := time.Now()
	id, _ := strconv.Atoi(c.Input[0])
	grid := build(id)
	result.ParseTime = time.Since(t0)

	t1 := time.Now()
	result.Solution1 = part1(grid, 3)
	result.Duration1 = time.Since(t1)

	t2 := time.Now()
	result.Solution2 = part2(grid)
	result.Duration2 = time.Since(t2)
	return
}

func part2(grid [][]int) string {
	var power, maxX, maxY, size int
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
	logger.Debug("For any size, the most power is", power)
	logger.Debugf("It had the identifier %d,%d,%d", maxX+1, maxY+1, size)
	return fmt.Sprintf("%d,%d,%d", maxX+1, maxY+1, size)
}

func part1(grid [][]int, size int) string {
	var maxX, maxY, max int
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
	logger.Debug("The biggest power cell has a power of ", max)
	logger.Debugf("It starts at %d,%d\n", maxX+1, maxY+1)
	return fmt.Sprintf("%d,%d", maxX+1, maxY+1)
}

func calc(grid [][]int, x, y, size int) (sum int) {
	for dx := x; dx < x+size; dx++ {
		for dy := y; dy < y+size; dy++ {
			sum += grid[dy][dx]
		}
	}
	return
}

func build(id int) (grid [][]int) {
	grid = make([][]int, 300)
	for y := range grid {
		grid[y] = make([]int, 300)
		for x := range grid[y] {
			grid[y][x] = calcPower(x, y, id)
		}
	}
	return grid
}

func calcPower(x, y, id int) int {
	rid := (x + 1) + 10 // Our x is 0 based, their x is 1 based
	p := rid * (y + 1)
	p += id
	p *= rid
	p = (p / 100) % 10
	return p - 5
}
