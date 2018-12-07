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
	input := util.GetInput("input/day07")
	s := parse(input)
	part1(s)
}

func part1(s sequence) {
	for {
		an := s.getAvailableNodes()
		if len(an) == 0 {
			break
		}
		sort.Strings(an)
		s.nextStep(an[0])
	}
	fmt.Println("The correct sequence is", strings.Join(s.result, ""))
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
