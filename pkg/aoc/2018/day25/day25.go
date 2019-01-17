package day25

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"gitlab.com/leononame/advent-of-code/pkg/aoc"
)

var logger *logrus.Logger

func Run(c *aoc.Config) (result aoc.Result) {
	logger = c.Logger

	t0 := time.Now()
	ps := parse(c.Input)
	result.ParseTime = time.Since(t0)

	t1 := time.Now()
	var cs constellations
outer:
	for _, p := range ps {
		for _, c := range cs {
			if c.addIfBelongs(p) {
				for cs.merge() {
				}
				continue outer
			}
		}
		cs = append(cs, &constellation{p})
	}
	result.Solution1 = len(cs)
	result.Duration1 = time.Since(t1)
	return
}

type point struct {
	a, b, c, d int
}

func (p point) distance(p2 point) int {
	return abs(p.a-p2.a) + abs(p.b-p2.b) + abs(p.c-p2.c) + abs(p.d-p2.d)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

type constellation []point

func (c *constellation) belongs(p point) bool {
	for _, cp := range *c {
		if cp.distance(p) <= 3 {
			return true
		}
	}
	return false
}
func (c *constellation) addIfBelongs(p point) bool {
	if c.belongs(p) {
		*c = append(*c, p)
		return true
	}
	return false
}

func (c *constellation) mergeIfBelongs(other *constellation) bool {
	if c == other {
		return false
	}
	for _, p := range *other {
		if c.belongs(p) {
			*c = append(*c, *other...)
			return true
		}
	}
	return false
}

type constellations []*constellation

func (cs *constellations) merge() bool {
	for i := 0; i < len(*cs); i++ {
		for j := i + 1; j < len(*cs); j++ {
			c1 := (*cs)[i]
			c2 := (*cs)[j]
			if c1.mergeIfBelongs(c2) {
				// Remove second constellation
				*cs = append((*cs)[0:j], (*cs)[j+1:len(*cs)]...)
				return true
			}
		}
	}
	return false
}

func parse(input []string) []point {
	var ps []point
	for _, l := range input {
		var p point
		fmt.Sscanf(l, "%d,%d,%d,%d", &p.a, &p.b, &p.c, &p.d)
		ps = append(ps, p)
	}
	return ps
}
