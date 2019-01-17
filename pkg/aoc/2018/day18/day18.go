package day18

import (
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"gitlab.com/leononame/advent-of-code/pkg/aoc"
)

var logger *logrus.Logger

func Run(c *aoc.Config) (result aoc.Result) {
	logger = c.Logger

	t0 := time.Now()
	a := parse(c.Input)
	result.ParseTime = time.Since(t0)

	t1 := time.Now()
	for a.tick() < 10 {
	}
	result.Solution1 = a.score()
	result.Duration1 = time.Since(t1)

	result.Solution2 = a.scoreFor(1000000000)
	result.Duration2 = time.Since(t1)
	return
}

const tree = '|'
const lumberyard = '#'
const ground = '.'

type area struct {
	data [50][50]byte
	cp   [50][50]byte
	age  int
}

func (a *area) String() string {
	var sb strings.Builder
	for y := range a.data {
		for x := range a.data[y] {
			sb.WriteByte(a.data[y][x])
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

func (a *area) tick() int {
	a.cp = a.data
	for y := range a.data {
		for x := range a.data[y] {
			el := a.cp[y][x]
			if el == '.' && a.count(x, y, tree) > 2 {
				a.data[y][x] = '|'
			} else if el == '|' && a.count(x, y, lumberyard) > 2 {
				a.data[y][x] = '#'
			} else if el == '#' && (a.count(x, y, tree) == 0 || a.count(x, y, lumberyard) == 0) {
				a.data[y][x] = '.'
			}
		}
	}
	a.age++
	return a.age
}

func (a *area) scoreFor(n int) int {
	c := a.age
	if n < c {
		return c
	}
	const threshold = 1000
	// If less than 100 ticks, don't look for pattern
	if n < threshold {
		for a.tick() < threshold {
		}
		return a.score()
	}

	// Otherwise, find pattern
	idx := 0
	var scores []int
	for i := 0; ; i++ {
		a.tick()
		score := a.score()
		scores = append(scores, score)
		if a.age == threshold {
			idx = a.age - c - 1
		} else if a.age > threshold && score == scores[idx] {
			// fount pattern
			curr := len(scores) - 1
			repeat := curr - idx
			offset := (n - a.age) % repeat
			idx = curr - (repeat - offset)
			break
		}
	}
	return scores[idx]
}

func (a *area) score() int {
	trees := 0
	yards := 0
	for y := range a.data {
		for x := range a.data[y] {
			if a.data[y][x] == tree {
				trees++
			} else if a.data[y][x] == lumberyard {
				yards++
			}
		}
	}
	return trees * yards
}

func (a *area) count(x, y int, acre byte) int {
	c := 0
	for dx := x - 1; dx <= x+1; dx++ {
		for dy := y - 1; dy <= y+1; dy++ {
			if dx < 0 || dy < 0 || dx == len(a.data) || dy == len(a.data) || (dx == x && dy == y) {
				continue
			}
			if a.cp[dy][dx] == acre {
				c++
			}
		}
	}
	return c
}

func parse(input []string) *area {
	var a area

	for y, l := range input {
		for x := range l {
			a.data[y][x] = l[x]
		}
	}
	return &a
}
