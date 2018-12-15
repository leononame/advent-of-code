package main

import (
	"fmt"
	"strconv"

	"gitlab.com/leononame/advent-of-code-2018/pkg/util"
	"gitlab.com/leononame/advent-of-code-2018/pkg/version"
)

func main() {
	fmt.Println("Advent of Code 2018, ", version.Str)
	fmt.Println("Challenge: 2018-14")
	count, _ := strconv.Atoi(util.GetInput("input/day14")[0])
	recipes := []int{3, 7}
	e1, e2 := 0, 1
	for len(recipes) < count+10 {
		ne1, ne2 := recipes[e1], recipes[e2]
		s := ne1 + ne2
		// if s is greater than 9, it's in the 10-18 range. Append 1
		if s > 9 {
			recipes = append(recipes, 1)
			s -= 10
		}
		// Append s
		recipes = append(recipes, s)
		// New indices
		e1 = (e1 + ne1 + 1) % len(recipes)
		e2 = (e2 + ne2 + 1) % len(recipes)
	}
	for _, v := range recipes[count : count+10] {
		fmt.Print(v)
	}
	fmt.Println()
}
