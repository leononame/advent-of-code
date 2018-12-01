package main

import (
	"bufio"
	"fmt"
	"gitlab.com/leononame/advent-of-code-2018/pkg/version"
	"os"
	"strconv"
)

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	println("Verison:\t" + version.Str)
	println("Challenge:\t2018-01")

	if len(os.Args) < 3 {
		println("Usage: ./bin part input\nPart is 1 or 2\ninput is the path to the input file")
		os.Exit(1)
	}
	part, err := strconv.Atoi(os.Args[1])
	checkErr(err)

	path := os.Args[2]
	f, err := os.Open(path)
	checkErr(err)

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
	checkErr(err)
	if line[0] == '-' {
		v = -v
	}
	return v
}
