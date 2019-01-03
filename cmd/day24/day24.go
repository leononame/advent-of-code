package main // package main

import (
	"fmt"

	"gitlab.com/einfachst/dgserver/pkg/version"
	"gitlab.com/leononame/advent-of-code-2018/pkg/day24"
	"gitlab.com/leononame/advent-of-code-2018/pkg/util"
)

func main() {
	fmt.Println("advent of code 2018, ", version.Str)
	fmt.Println("challenge: 2018-24")
	input := util.GetInput("input")
	day24.Run(input)
}
