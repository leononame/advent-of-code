package main

import (
	"fmt"

	"gitlab.com/leononame/advent-of-code-2018/pkg/aoc/2018/day15"
	"gitlab.com/leononame/advent-of-code-2018/pkg/util"
	"gitlab.com/leononame/advent-of-code-2018/pkg/version"
)

func main() {
	fmt.Println("advent of code 2018, ", version.Str)
	fmt.Println("challenge: 2018-15")
	input := util.GetInput("input")
	day15.Run(input)
}
