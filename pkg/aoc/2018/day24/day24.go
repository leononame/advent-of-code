package day24

import (
	"regexp"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
	"gitlab.com/leononame/advent-of-code/pkg/aoc"
)

var logger *logrus.Logger

func Run(c *aoc.Config) (result aoc.Result) {
	logger = c.Logger

	t0 := time.Now()
	b := parse(c.Input)
	result.ParseTime = time.Since(t0)

	t1 := time.Now()
	result.Solution1 = part1(b)
	result.Duration1 = time.Since(t1)

	t2 := time.Now()
	result.Solution2 = part2(c.Input)
	result.Duration2 = time.Since(t2)
	return
}

func part2(input []string) int {
	// Binary search the radius
	radius := 1 << 8
	boost := radius
	for {
		b := parse(input)
		for _, g := range b.imm {
			g.str += boost
		}

		won := b.fight() == &b.imm
		// Always half the search radius, but it shouldn't go below 1
		// This is needed because this binary search doesn't know immediately whether
		// the correct value was found. It still needs to go back and forth one step
		radius /= 2
		if radius == 0 {
			radius = 1
		}
		logger.Debugf("Boost: %d, Radius: %d, Winner: %t", boost, radius, won)

		if !won {
			boost += radius
			continue
		}

		if radius == 1 {
			units := b.imm.countUnits()

			logger.Debugf("Immune System won with boost %d. Remaining units: %d\n",
				boost,
				units)
			return units
		}
		// We have reached a winning point, half the boost and look backwards
		boost -= radius
	}
}

func part1(b *battle) int {
	// expected: 26937
	winner := b.fight()
	name := "Immune System"
	if winner == &b.inf {
		name = "Infection"
	}
	units := winner.countUnits()
	logger.Debugf("%s won with %d remaining units\n", name, units)
	return units
}

func atoi(in string) int {
	out, err := strconv.Atoi(in)
	if err != nil {
		panic(err)
	}
	return out
}

func parse(input []string) *battle {
	r := regexp.MustCompile(`(\d+) units each with (\d+) hit points (.+)? ?with an attack that does (\d+) (\w+) damage at initiative (\d+)`)
	rw := regexp.MustCompile(`weak to ([\w, ]+)`)
	ri := regexp.MustCompile(`immune to ([\w, ]+)`)
	parseArmy := func(input []string) (int, army) {
		a := army{}
		for i, l := range input {
			if l == "" {
				return i, a
			}
			g := &group{}
			m := r.FindStringSubmatch(l)[1:]
			g.n, g.hp, g.str, g.init = atoi(m[0]), atoi(m[1]), atoi(m[3]), atoi(m[5])
			g.id = i + 1
			g.dtype = m[4]
			if m[2] != "" {
				if w := rw.FindStringSubmatch(m[2]); w != nil {
					g.weak = w[1]
					// g.weak = strings.Split(w[1], ", ")
				}
				if i := ri.FindStringSubmatch(m[2]); i != nil {
					g.immune = i[1]
					// g.immune = strings.Split(i[1], ", ")
				}
			}
			a = append(a, g)
		}
		return 0, a
	}
	idx, imm := parseArmy(input[1:])
	_, inf := parseArmy(input[idx+3:])
	all := append(inf, imm...)
	return &battle{imm, inf, all}
}
