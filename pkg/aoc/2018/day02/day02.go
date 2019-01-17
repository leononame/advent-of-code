package day02

import (
	"time"

	"github.com/sirupsen/logrus"

	"gitlab.com/leononame/advent-of-code-2018/pkg/aoc"
)

var logger *logrus.Logger

func Run(config *aoc.Config) aoc.Result {
	logger = config.Logger
	logger.Info(`Usage:
aoc -y 2018 -d 02 [SUBCOMMAND]
Available Subcommands:
  opt	Optimized version that looks for strings with Hamming Distance 1
`)
	r := aoc.Result{}

	r.ParseTime = 0

	t1 := time.Now()
	r.Solution1 = part1(config.Input)
	r.Duration1 = time.Since(t1)

	p2 := part2
	if config.SubCommand == "opt" {
		p2 = part2opt
	}

	t2 := time.Now()
	r.Solution2 = p2(config.Input)
	r.Duration2 = time.Since(t2)
	return r
}

func part1(input []string) int {
	n2, n3 := 0, 0
	// How often a number appeared
	for i, line := range input {
		// Count occurences for each letters
		counter := make(map[rune]int)
		for _, r := range line {
			counter[r]++
		}
		// Count how often a letter appeared 0, 1, 2, ... times
		var counts [26]int
		for _, v := range counter {
			counts[v]++
		}
		// If at least one letter appeared twice, increment n2
		if counts[2] > 0 {
			n2++
		}
		// If at least one letter appeared three times, increment n3
		if counts[3] > 0 {
			n3++
		}
		logger.Debugf("%3d. %d letters appear twice, %d three times. n2: %d, ne: %d\n",
			i,
			counts[2],
			counts[3],
			n2,
			n3)
	}
	// Hash is n2*n3
	return n2 * n3
}

func part2(input []string) string {
	a, b := findMinPair(input)
	return findSharedLetters(input[a], input[b])
}

func part2opt(input []string) string {
	a, b := findMinPairOptimized(input)
	return findSharedLetters(input[a], input[b])
}

// findSharedLetters compares two strings and returns a string that contains the shared letters,
// i.e. all the letters that appear in both strings at the same position
func findSharedLetters(s1, s2 string) string {
	logger.Debugf("s1: %s", s1)
	logger.Debugf("s2: %s", s2)
	b1 := []byte(s1)
	b2 := []byte(s2)
	s := ""
	for i, v := range b1 {
		if b2[i] == v {
			s += string(v)
		}
	}
	return s
}

// findMinPair searches for the pair of strings which have lowest Hamming Distance
func findMinPair(input []string) (int, int) {
	min := len(input[0])
	p1, p2 := 0, 0
	for i, l := range input {
		for j := i + 1; j < len(input); j++ {
			dist := hamming(l, input[j])
			logger.Debugf("Min: %2d. %3d, %3d have distance %2d.", min, i, j, dist)
			if dist < min {
				logger.Debugf("Selecting new strings at %3d, %3d", i, j)
				min = dist
				p1, p2 = i, j
			}
		}
	}
	return p1, p2
}

// findMinPairOptimized searches for the pair of strings which have a Hamming Distance of 1
func findMinPairOptimized(input []string) (int, int) {
	for i, l := range input {
		for j := i + 1; j < len(input); j++ {
			dist := hamming(l, input[j])
			logger.Debugf("%3d, %3d have distance %2d.", i, j, dist)
			if dist == 1 {
				logger.Debug("FOUND!")
				return i, j
			}
		}
	}
	return 0, 0
}

// hamming calculates the Hamming Distance of two strings.
// This is done on a byte-level, not bit-level. That is 2 differing bytes increase the
// Hamming Distance by 1.
func hamming(s1, s2 string) int {
	if len(s1) != len(s2) {
		// I don't care if I chose the shorter string by accident
		return len(s1)
	}
	b1 := []byte(s1)
	b2 := []byte(s2)
	h := 0
	for i, v := range b1 {
		if b2[i] != v {
			h++
		}
	}
	return h
}
