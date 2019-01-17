package day15

import (
	"strings"

	"gitlab.com/leononame/advent-of-code-2018/pkg/geo"
)

const (
	goblin = 'G'
	elf    = 'E'
	floor  = '.'
	wall   = '#'
)

type cave struct {
	elves, goblins []*unit
	all            []*unit
	layout         [][]rune
}

func (c *cave) String() string {
	var sb strings.Builder
	for _, l := range c.layout {
		for _, r := range l {
			sb.WriteRune(r)
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

func (c *cave) PrintInfo() {
	logger.Debug("Cave:\n", c)
}

func (c *cave) fightOver() bool {
	return len(c.goblins) == 0 || len(c.elves) == 0
}

// removeKilled removes a maximum of 1 enemy of each unit slice
// not beautiful code, I admit it, but I don't like this challenge and don't care
func (c *cave) removeKilled(idx int) int {
	for i := 0; i < len(c.elves); {
		if c.elves[i].hp == 0 {
			c.elves = append(c.elves[:i], c.elves[i+1:]...)
			break
		}
		i++
	}

	for i := 0; i < len(c.goblins); {
		if c.goblins[i].hp == 0 {
			c.goblins = append(c.goblins[:i], c.goblins[i+1:]...)
			break
		}
		i++
	}

	for i := 0; i < len(c.all); {
		if c.all[i].hp == 0 {
			c.layout[c.all[i].pos.GetY()][c.all[i].pos.GetX()] = floor
			c.all = append(c.all[:i], c.all[i+1:]...)
			if i > idx {
				return idx + 1
			}
			return idx
		}
		i++
	}
	return idx
}

func (c *cave) neighbours(p geo.Point) []geo.Point {
	var ret []geo.Point
	for _, n := range p.Neighbours() {
		y, x := n.GetY(), n.GetX()
		if x < 0 || y < 0 || y >= len(c.layout) || x >= len(c.layout[0]) {
			continue
		}
		if c.layout[y][x] == floor {
			ret = append(ret, n)
		}
	}
	return ret
}

var count = 0

func (c *cave) addUnit(pos geo.Point, t rune) {
	u := &unit{hp: 200, pos: pos, t: t, cave: c, id: count, pow: 3}
	count++
	c.all = append(c.all, u)
	if t == goblin {
		c.goblins = append(c.goblins, u)
		u.enemies = &c.elves
	} else {
		c.elves = append(c.elves, u)
		u.enemies = &c.goblins
	}
}

func (c *cave) Len() int {
	return len(c.all)
}

func (c *cave) Less(i, j int) bool {
	a, b := c.all[i].pos, c.all[j].pos
	return readingOrder(a, b)
}

func (c *cave) Swap(i, j int) {
	c.all[i], c.all[j] = c.all[j], c.all[i]
}
