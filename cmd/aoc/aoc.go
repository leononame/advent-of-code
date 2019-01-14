package main

import (
	"fmt"
	"os"

	"github.com/logrusorgru/aurora"

	"gitlab.com/leononame/advent-of-code-2018/pkg/aoc"
	"gitlab.com/leononame/advent-of-code-2018/pkg/aoc/2018/day01"
	"gitlab.com/leononame/advent-of-code-2018/pkg/aoc/2018/day02"
)

func main() {
	// Setup challenge
	config := aoc.Setup()

	if runners[config.Year] == nil {
		config.Logger.Errorf("Year %d not found", config.Year)
		os.Exit(1)
	}
	if runners[config.Year][config.Day] == nil {
		config.Logger.Errorf("Year %d, day %d not found", config.Year, config.Day)
		os.Exit(1)
	}

	r := runners[config.Year][config.Day](config)
	fmt.Printf("Advent of Code: Year %d, Day %02d\n", config.Year, config.Day)
	fmt.Println(aurora.Gray("Parse:"), aurora.Black(r.ParseTime))
	fmt.Println(aurora.Gray("Part1:"), aurora.Green(r.Solution1), aurora.Black(r.Duration1))
	fmt.Println(aurora.Gray("Part2:"), aurora.Green(r.Solution2), aurora.Black(r.Duration2))

}

var runners = map[int]map[int]func(config *aoc.Config) aoc.Result{
	2018: {
		1: day01.Run,
		2: day02.Run,
	},
}
