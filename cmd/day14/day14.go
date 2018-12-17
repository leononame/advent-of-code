package main

import (
	"bytes"
	"fmt"
	"strconv"

	"gitlab.com/leononame/advent-of-code-2018/pkg/util"
	"gitlab.com/leononame/advent-of-code-2018/pkg/version"
)

func main() {
	fmt.Println("Advent of Code 2018, ", version.Str)
	fmt.Println("Challenge: 2018-14")
	input := util.GetInput("input/day14")[0]
	count, _ := strconv.Atoi(input)

	recipes := []byte{'3', '7'}
	var e1, e2 = 0, 1
	for i := 1; ; i++ {
		for len(recipes) < i*(count+10) {
			ne1, ne2 := recipes[e1]-'0', recipes[e2]-'0'
			s := strconv.Itoa(int(ne1 + ne2))
			recipes = append(recipes, s...)
			// New indices
			e1 = (e1 + int(ne1) + 1) % len(recipes)
			e2 = (e2 + int(ne2) + 1) % len(recipes)
		}
		p2 := bytes.Index(recipes, []byte(input))
		if p2 != -1 {
			fmt.Println("Part 2:", p2)
			break
		}
	}
	fmt.Println("Part 1:", string(recipes[count:count+10]))
}
