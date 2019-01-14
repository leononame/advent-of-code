package main

import (
	"fmt"

	"gitlab.com/einfachst/dgserver/pkg/version"
	"gitlab.com/leononame/advent-of-code-2018/pkg/aoc/2018/day25"
	"gitlab.com/leononame/advent-of-code-2018/pkg/util"
)

func main() {
	fmt.Println("advent of code 2018, ", version.Str)
	fmt.Println("challenge: 2018-25")
	input := util.GetInput("input")
	day25.Run(input)
}
