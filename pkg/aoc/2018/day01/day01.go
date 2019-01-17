package day01

import (
	"strconv"
	"time"

	"github.com/sirupsen/logrus"

	"gitlab.com/leononame/advent-of-code/pkg/aoc"
)

var logger *logrus.Logger

func Run(config *aoc.Config) aoc.Result {
	logger = config.Logger
	r := aoc.Result{}

	t0 := time.Now()
	vals := parse(config)
	r.ParseTime = time.Since(t0)

	r.ParseTime = time.Duration(0)
	t1 := time.Now()
	r.Solution1 = part1(vals)
	r.Duration1 = time.Since(t1)

	t2 := time.Now()
	r.Solution2 = part2(vals)
	r.Duration2 = time.Since(t2)
	return r
}

func parse(config *aoc.Config) []int {
	var vals []int
	for _, l := range config.Input {
		v, _ := strconv.Atoi(l)
		vals = append(vals, v)
	}
	return vals
}

func part1(vals []int) int {
	result := 0
	for _, v := range vals {
		result += v
	}
	return result
}

func part2(vals []int) int {
	seen := map[int]bool{0: true}
	result := 0
	for {
		logger.Debug(result)
		for _, v := range vals {
			result += v
			if seen[result] {
				return result
			}
			seen[result] = true
		}
	}
}
