package day15

import (
	"fmt"
	"sort"

	"gitlab.com/leononame/advent-of-code-2018/pkg/geo"
	"gitlab.com/leononame/advent-of-code-2018/pkg/geo/points"
)

func Run(input []string) {
	// Part 1: 250648
	// Part 2: 42392 25
	c := parse(input)
	part1(c)
	part2(input)
}

func part2(input []string) {
	// pow := 20
	for pow := 4; ; pow++ {
		fmt.Println("Trying attack power", pow)
		c := parse(input)
		for _, e := range c.elves {
			e.pow = pow
		}
		nelves := len(c.elves)
		part1(c)
		if nelves-len(c.elves) == 0 {
			fmt.Println("Elves have not lost any unit with attack power", pow)
			return
		}
	}
}

func part1(c *cave) {
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
	fmt.Printf("Round %d, Goblins: %d, Elves %d\n", round, len(c.goblins), len(c.elves))
	if len(c.elves) == 0 {
		fmt.Printf("Goblins won. Round %d, Hitpoints %d\n", round, hitpoints(c.goblins))
		fmt.Println("Outcome:", round*hitpoints(c.goblins))
		return
	}
	if len(c.goblins) == 0 {
		fmt.Printf("Elves won. Round %d, Hitpoints %d\n", round, hitpoints(c.elves))
		fmt.Println("Outcome:", round*hitpoints(c.elves))
		return
	}
}

func hitpoints(us []*unit) int {
	sum := 0
	for _, u := range us {
		sum += u.hp
	}
	return sum
}

// Checks whether a and b are in reading order
func readingOrder(a, b geo.Pointer) bool {
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
			pos := points.NewClassicPointer(j, i)
			if el == goblin || el == elf {
				c.addUnit(pos, el)
			}
			c.layout[i][j] = el
		}
	}
	return c
}
