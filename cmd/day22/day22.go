package main

import (
	"fmt"

	"gitlab.com/leononame/advent-of-code-2018/pkg/day22"
	"gitlab.com/leononame/advent-of-code-2018/pkg/geo/points"
	"gitlab.com/leononame/advent-of-code-2018/pkg/util"
	"gitlab.com/leononame/advent-of-code-2018/pkg/version"
)

func main() {
	fmt.Println("advent of code 2018, ", version.Str)
	fmt.Println("challenge: 2018-22")
	input := util.GetInput("input")
	day22.Run(input, points.NewClassicPointer)
}