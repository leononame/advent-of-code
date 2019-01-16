package day15

import (
	"sort"
	"time"

	"github.com/sirupsen/logrus"
	"gitlab.com/leononame/advent-of-code-2018/pkg/aoc"

	"gitlab.com/leononame/advent-of-code-2018/pkg/geo"
)

var logger *logrus.Logger

func Run(config *aoc.Config) (result aoc.Result) {
	logger = config.Logger

	t0 := time.Now()
	c := parse(config.Input)
	result.ParseTime = time.Since(t0)

	t1 := time.Now()
	result.Solution1 = part1(c)
	result.Duration1 = time.Since(t1)

	t2 := time.Now()
	result.Solution2 = part2(config.Input)
	result.Duration2 = time.Since(t2)
	return
}

func part2(input []string) int {
	// pow := 20
	for pow := 4; ; pow++ {
		logger.Debug("Trying attack power ", pow)
		c := parse(input)
		for _, e := range c.elves {
			e.pow = pow
		}
		nelves := len(c.elves)
		outcome := part1(c)
		if nelves-len(c.elves) == 0 {
			logger.Debug("Elves have not lost any unit with attack power ", pow)
			return outcome
		}
	}
}

func part1(c *cave) int {
	round := 0
outer:
	for {
		sort.Sort(c)
		for i := 0; i < len(c.all); {
			u := c.all[i]
			enemyKilled := u.tick()
			if enemyKilled {
				i = c.removeKilled(i)
				if c.fightOver() {
					break outer
				}
			} else {
				i++
			}
		}
		round++
		// c.PrintInfo()
	}

	c.PrintInfo()
	logger.Debugf("Round %d, Goblins: %d, Elves %d\n", round, len(c.goblins), len(c.elves))
	if len(c.elves) == 0 {
		logger.Debugf("Goblins won. Round %d, Hitpoints %d\n", round, hitpoints(c.goblins))
		logger.Debug("Outcome: ", round*hitpoints(c.goblins))
		return round * hitpoints(c.goblins)
	}

	logger.Debugf("Elves won. Round %d, Hitpoints %d\n", round, hitpoints(c.elves))
	logger.Debug("Outcome: ", round*hitpoints(c.elves))
	return round * hitpoints(c.elves)

}

func hitpoints(us []*unit) int {
	sum := 0
	for _, u := range us {
		sum += u.hp
	}
	return sum
}

// Checks whether a and b are in reading order
func readingOrder(a, b geo.Point) bool {
	return a.GetY() < b.GetY() || (a.GetY() == b.GetY() && a.GetX() < b.GetX())
}

func parse(input []string) *cave {
	c := &cave{}
	c.layout = make([][]rune, len(input))
	xlen := len(input[0])
	for i := range c.layout {
		c.layout[i] = make([]rune, xlen)
	}
	for i := range input {
		for j, el := range []rune(input[i]) {
			pos := geo.Point{j, i}
			if el == goblin || el == elf {
				c.addUnit(pos, el)
			}
			c.layout[i][j] = el
		}
	}
	return c
}
