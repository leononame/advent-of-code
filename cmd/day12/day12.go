package main

import (
	"fmt"
	"strings"

	"gitlab.com/leononame/advent-of-code-2018/pkg/util"
	"gitlab.com/leononame/advent-of-code-2018/pkg/version"
)

type generation struct {
	index int
	s     string
	rules map[string]string
}

func (g *generation) sum() int {
	s := 0
	for i, v := range g.s {
		if v == '#' {
			s += i + g.index
		}
	}
	return s
}

func main() {
	fmt.Println("Advent of Code 2018, ", version.Str)
	fmt.Println("Challenge: 2018-12")
	input := util.GetInput("input/day12")
	g := parseInput(input)
	generations := []generation{g}
	for i := 0; i < 20; i++ {
		generations = append(generations, next(generations[i]))
	}
	fmt.Println("Generation 20: ", generations[20].s)
	fmt.Println("Total count: ", generations[20].sum())
}

func parseInput(input []string) generation {
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

// ###.......###.........#...###.....###.#....#.###..#...#....#...#.#.......#.#..............................###
// ###.......###.........#...###.....###.#....#.###..#...#....#...#.#.......#.#..............................###
