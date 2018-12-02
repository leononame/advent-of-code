package main

import (
	"fmt"
	"gitlab.com/leononame/advent-of-code-2018/pkg/util"
	"os"
	"strings"
)

var alphabet [26]string
var parts = map[string]func(*[]string){
	"1": part1,
}

func main() {
	util.CheckArgs()
	println("Challenge:\t2018-02")

	input := util.GetInput()

	// Build the alphabet lower case
	for i := 0; i < 26; i++ {
		alphabet[i] = string(i+97)
	}


	// Run function
	parts[os.Args[1]](input)
}

func part1(input *[]string) {
	// How often a number appeared
	counter := make(map[int]int)
	for _, line := range *input {
		m := countLetters(line)
		for key := range m {
			counter[key]++
		}
	}

	hash := 1
	for _, v := range counter {
		hash *= v
	}
	fmt.Println("Result: ")
	fmt.Println(counter)
	fmt.Printf("Hash: %d", hash)
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