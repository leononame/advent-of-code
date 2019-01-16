package day12

import (
	"fmt"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"gitlab.com/leononame/advent-of-code-2018/pkg/aoc"
)

var logger *logrus.Logger

func Run(c *aoc.Config) (result aoc.Result) {
	logger = c.Logger

	t0 := time.Now()
	g := parse(c.Input)
	result.ParseTime = time.Since(t0)

	t1 := time.Now()
	result.Solution1 = part1(g)
	result.Duration1 = time.Since(t1)

	t2 := time.Now()
	result.Solution2 = part2(g)
	result.Duration2 = time.Since(t2)
	return
}

type generation struct {
	index int
	s     string
	rules map[string]string
}

func (g *generation) sum() int {
	var s int
	for i, v := range g.s {
		if v == '#' {
			s += i + g.index
		}
	}
	return s
}

func part1(g generation) int {
	generations := []generation{g}
	for i := 0; i < 20; i++ {
		generations = append(generations, next(generations[i]))
	}

	sum := generations[20].sum()
	logger.Debug("Generation 20: ", generations[20].s)
	logger.Debug("Total count: ", sum)
	return sum
}

func parse(input []string) generation {
	var g generation
	fmt.Sscanf(input[0], "initial state: %s", &g.s)
	g.index = 0
	g.rules = make(map[string]string)
	for _, s := range input[2:] {
		rd := strings.Split(s, " => ")
		g.rules[rd[0]] = rd[1]
	}
	return g
}

func next(g generation) generation {
	var next strings.Builder
	// Extend string with 4 dots so it can grow out of bounds
	var extended = "...." + g.s + "...."

	// Iterate over everything
	for i := 0; i < len(g.s)+4; i++ {
		pattern := extended[i : i+5]
		next.WriteString(g.rules[pattern])
	}

	// get new index offset
	idx := g.index - 2
	n := next.String()
	for n[0] == '.' {
		n = n[1:]
		idx++
	}
	n = strings.Trim(n, ".")
	return generation{idx, n, g.rules}
}

func part2(g generation) int {
	// Increment becomes constant at some point, just watch for it
	diff, sum, count := 0, g.sum(), 0
	for i := 1; ; i++ {
		g = next(g)
		s := g.sum()
		d := s - sum
		sum = s
		if d == diff {
			count++
		} else {
			count = 0
			diff = d
		}

		// If for 20 iterations the diff between sums hasn't changed, it's steady
		if count == 20 {
			return s + (5e10-i)*d
		}
	}
}
