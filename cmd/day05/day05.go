package main

import (
	"fmt"
	"strings"

	"gitlab.com/leononame/advent-of-code-2018/pkg/util"
)

func main() {
	fmt.Println("Challenge:\t2018-05")
	input := util.GetInput("input/day05")[0]

	length := react(input)
	fmt.Printf("Part 1 resulted in a deconstructed string of lenght: %d\n", length)

	min := len(input)
	badChar := "a"
	for i := 0x41; i < 0x5B; i++ {
		cleaned := removeChar(input, i)
		length := react(cleaned)
		if length < min {
			min = length
			badChar = string(i)
		}
	}
	fmt.Printf("For part 2, the best solution is removing the character %s. This results in a length of: %d\n", badChar, min)
}

func removeChar(s string, c int) string {
	tmp := strings.Replace(s, string(c), "", -1)
	return strings.Replace(tmp, string(c+0x20), "", -1)
}

func react(polymer string) int {
	s1 := polymer
	s2 := parse(polymer)
	for s2 != s1 {
		s1 = s2
		s2 = parse(s1)
	}
	return len(s1)
}

func parse(data string) string {
	for i := 0; i < len(data)-1; i++ {
		if abs(int8(data[i]-data[i+1])) == 0x20 {
			return data[:i] + data[i+2:]
		}
	}
	return data
}

func abs(v int8) int8 {
	if v < 0 {
		return -v
	}
	return v
}