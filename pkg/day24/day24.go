package day24

import (
	"fmt"
	"regexp"
	"strconv"
)

func Run(input []string) {
	b := parse(input)
	part1(b)
	part2(input)
}

func part2(input []string) {
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
		fmt.Printf("\rBoost: %d, Radius: %d, Winner: %t\n", boost, radius, won)

		if !won {
			boost += radius
			continue
		}

		if radius == 1 {
			fmt.Printf("Immune System won with boost %d. Remaining units: %d\n",
				boost,
				b.imm.countUnits())
			return
		}
		// We have reached a winning point, half the boost and look backwards
		boost -= radius
	}
}

func part1(b *battle) {
	// expected: 26937
	winner := b.fight()
	name := "Immune System"
	if winner == &b.inf {
		name = "Infection"
	}
	fmt.Printf("%s won with %d remaining units\n", name, winner.countUnits())
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
