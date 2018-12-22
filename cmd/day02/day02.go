package main

import (
	"fmt"
	"strings"

	"gitlab.com/leononame/advent-of-code-2018/pkg/util"
)

var alphabet [26]string

func main() {
	fmt.Println("Challenge:\t2018-02")

	input := util.GetInput("input")

	// Build the alphabet lower case
	for i := 0; i < 26; i++ {
		alphabet[i] = string(i + 97)
	}

	// Run function
	part1(input)
	part2(input)
}

// Part 1
func part1(input []string) {
	// How often a number appeared
	counter := make(map[int]int)
	for _, line := range input {
		m := countLetters(line)
		for key := range m {
			counter[key]++
		}
	}

	hash := 1
	for _, v := range counter {
		hash *= v
	}
	fmt.Printf("Resulting counts:: ")
	fmt.Println(counter)
	fmt.Printf("Result part 1: %d\n", hash)
}

func countLetters(line string) map[int]struct{} {
	// Counter holds the information if letters appear more than one timer
	counter := make(map[int]struct{})
	for _, letter := range alphabet {
		c := strings.Count(line, letter)
		if c > 1 {
			counter[c] = struct{}{}
		}
	}
	return counter
}

// Part 2
func part2(input []string) {
	a, b := findMinPair(input)
	s1 := input[a]
	s2 := input[b]
	result := findSharedLetters(input[a], input[b])
	fmt.Printf("The two most similar IDs are:\n")
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Printf("They share the following letter sequence: %s\n", result)
}

func findSharedLetters(s1, s2 string) string {
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

func findMinPair(input []string) (int, int) {
	for i, l := range input {
		for j := i + 1; j < len(input); j++ {
			if hamming(l, input[j]) == 1 {
				return i, j
			}
		}
	}
	return 0, 0
}

func hamming(s1, s2 string) int {
	if len(s1) != len(s2) {
		// I really don't care here
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
