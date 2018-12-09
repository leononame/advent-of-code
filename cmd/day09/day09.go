package main

import (
	"fmt"
	"strconv"
)

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

func main() {
	n := 465
	players := make(map[int]uint32)
	max := 71498

	c := makeCircle()

	// Part 1
	val := 1
	for val <= max {
		p := val % n // 0 is player 465
		players[p] += c.insertMarble(val)
		val++
	}
	// Find largest points
	fmt.Println("The winning score is ", maxScore(players))

	// Part 2
	max = 100 * max
	for val <= max {
		p := val % n // 0 is player 465
		players[p] += c.insertMarble(val)
		val++
	}
	// Find largest points
	fmt.Println("The winning score for part 2 is ", maxScore(players))

}

func maxScore(m map[int]uint32) uint32 {
	var max uint32
	for _, v := range m {
		if v > max {
			max = v
		}
	}
	return max
}

func (c *circle) insertMarble(v int) uint32 {
	if v%23 == 0 {
		m7 := c.current.ccw.ccw.ccw.ccw.ccw.ccw.ccw
		m7.cw.ccw = m7.ccw
		m7.ccw.cw = m7.cw
		c.current = m7.cw
		return uint32(v + m7.val)
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
