package main

import (
	"fmt"
	"os"
	"strconv"

	"gitlab.com/leononame/advent-of-code-2018/pkg/util"
)

func main() {
	fmt.Println("Challenge:\t2018-01")

	input := util.GetInput("input")
	part1(input)
	part2(input)
}

func part1(input []string) {
	result := 0
	for _, l := range input {
		v := calcValue(l)
		result += v
	}
	fmt.Printf("Result part 1: %d\n", result)
}

func part2(lines []string) {
	var freqs []int
	result := 0

	for {
		for _, l := range lines {
			v := calcValue(l)
			result += v
			if exists(freqs, result) {
				fmt.Printf("Result part 2: %d!\n", result)
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
