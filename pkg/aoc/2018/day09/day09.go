package day09

import (
	"fmt"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
	"gitlab.com/leononame/advent-of-code/pkg/aoc"
)

var logger *logrus.Logger

func Run(c *aoc.Config) (result aoc.Result) {
	logger = c.Logger

	t0 := time.Now()
	nplayers, value := parse(c.Input[0])
	result.ParseTime = time.Since(t0)

	t1 := time.Now()
	result.Solution1 = part1(nplayers, value)
	result.Duration1 = time.Since(t1)

	t2 := time.Now()
	result.Solution2 = part2(nplayers, value)
	result.Duration2 = time.Since(t2)
	return

	return
}

func parse(line string) (players, value int) {
	fmt.Sscanf(line, "%d players; last marble is worth %d points", &players, &value)
	return
}

type circle struct {
	current *marble
}
type marble struct {
	cw, ccw *marble
	val     int
}

func (m *marble) String() string {
	return strconv.Itoa(m.val)
}

func part1(n, max int) int {
	players := make(map[int]int)
	c := makeCircle()
	val := 1
	for val <= max {
		p := val % n // 0 is player 465
		players[p] += c.insertMarble(val)
		val++
	}
	score := maxScore(players)
	logger.Debug("The winning score is ", score)
	return score
}

func part2(n, max int) int {
	players := make(map[int]int)
	c := makeCircle()
	max *= 100
	val := 1
	for val <= max {
		p := val % n // 0 is player 465
		players[p] += c.insertMarble(val)
		val++
	}
	score := maxScore(players)
	logger.Debug("The winning score for part 2 is ", score)
	return score
}

func maxScore(m map[int]int) int {
	var max int
	for _, v := range m {
		if v > max {
			max = v
		}
	}
	return max
}

func (c *circle) insertMarble(v int) int {
	if v%23 == 0 {
		m7 := c.current.ccw.ccw.ccw.ccw.ccw.ccw.ccw
		m7.cw.ccw = m7.ccw
		m7.ccw.cw = m7.cw
		c.current = m7.cw
		return v + m7.val
	}
	m := &marble{val: v}
	m.cw = c.current.cw.cw
	m.ccw = c.current.cw
	c.current.cw.cw.ccw = m
	c.current.cw.cw = m
	c.current = m
	return 0
}

func (c *circle) slice() []*marble {
	f := c.current
	var r []*marble
	r = append(r, c.current)
	m := c.current
	for {
		m = m.cw
		if m == f {
			return r
		}
		r = append(r, m)
	}
	return nil
}

func makeCircle() *circle {
	c := circle{}
	c.current = &marble{val: 0}
	c.current.cw = c.current
	c.current.ccw = c.current
	return &c
}
