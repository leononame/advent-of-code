package main

import (
	"bufio"
	"fmt"
	"gitlab.com/leononame/advent-of-code-2018/pkg/util"
	"os"
	"strconv"
)

func main() {
	util.CheckArgs()
	println("Challenge:\t2018-01")

	part, err := strconv.Atoi(os.Args[1])
	util.CheckErr(err)

	path := os.Args[2]
	f, err := os.Open(path)
	util.CheckErr(err)

	s := bufio.NewScanner(f)
	if part == 1 {
		part1(s)
	} else {
		part2(s)
	}
}

func part1(s *bufio.Scanner) {
	result := 0
	for s.Scan() {
		v := calcValue(s.Text())
		result += v
		fmt.Printf("Sum: %d,\tValue: %d\n", result, v)
	}
	fmt.Printf("Result: %d\n", result)
}

func part2(s *bufio.Scanner) {
	var lines []string
	var freqs []int
	result := 0

	for s.Scan() {
		lines = append(lines, s.Text())
	}

	for {
		for _, l := range lines {
			v := calcValue(l)
			result += v
			fmt.Printf("Current sum: %d\n", result)
			if exists(freqs, result) {
				fmt.Printf("Sum %d already found!\n", result)
				os.Exit(0)
			}
			freqs = append(freqs, result)
		}
	}
}

func exists(slice []int, val int) bool {
	for _, v := range slice {
		if val == v {
			return true
		}
	}
	return false
}

func calcValue(line string) int {
	v, err := strconv.Atoi(line[1:])
	util.CheckErr(err)
	if line[0] == '-' {
		v = -v
	}
	return v
}
