package main

import (
	"fmt"
	"sort"
	"strings"

	"gitlab.com/leononame/advent-of-code-2018/pkg/util"
)

type sequence struct {
	requirements map[string][]string
	nodes        map[string]struct{}
	result       []string
}

func main() {
	fmt.Println("Challenge:\t2018-07")
	input := util.GetInput("input")
	s := parse(input)
	seq := part1(s)
	s = parse(input)
	part2(s, seq)
}

func part2(s sequence, order []string) {
	pool := 15
	tick := 0

	finished := make(map[string]int)
	for {
		for k, v := range finished {
			if v == tick {
				s.nextStep(k)
				pool++
			}
		}
		if len(s.nodes) == 0 {
			break
		}

		an := s.getAvailableNodes()
		sort.Strings(an)

		for i := 0; i < len(an) && pool > 0; i++ {
			if finished[an[i]] > 0 {
				continue
			}
			letter := an[i]
			duration := 60 + int(letter[0]) - 64
			finished[an[i]] = tick + duration
			pool--
		}
		tick++
	}

	fmt.Println("The sequence for part 2 is ", strings.Join(s.result, ""))
	fmt.Println("It would take", tick, "seconds")
}

func part1(s sequence) []string {
	for {
		an := s.getAvailableNodes()
		if len(an) == 0 {
			break
		}
		sort.Strings(an)
		s.nextStep(an[0])
	}
	fmt.Println("The correct sequence is", strings.Join(s.result, ""))
	return s.result
}

func parse(input []string) sequence {
	requirements := make(map[string][]string)
	nodes := make(map[string]struct{})
	for _, l := range input {
		ws := strings.Split(l, " ")
		before := ws[1]
		after := ws[7]
		nodes[before] = struct{}{}
		nodes[after] = struct{}{}
		requirements[after] = append(requirements[after], before)
	}
	return sequence{requirements, nodes, []string{}}
}

func (s *sequence) getAvailableNodes() []string {
	var available []string
	for node := range s.nodes {
		if s.isNodeAvailable(node) {
			available = append(available, node)
		}
	}
	return available
}

func (s *sequence) isNodeAvailable(node string) bool {
	for _, required := range s.requirements[node] {
		if !contains(s.result, required) {
			return false
		}
	}
	return true
}

func (s *sequence) nextStep(node string) {
	s.result = append(s.result, node)
	delete(s.nodes, node)
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
