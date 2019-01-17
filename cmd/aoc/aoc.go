package main

import (
	"fmt"
	"os"

	"gitlab.com/leononame/advent-of-code-2018/pkg/printer"

	"gitlab.com/leononame/advent-of-code-2018/pkg/aoc"
	"gitlab.com/leononame/advent-of-code-2018/pkg/aoc/2018/day01"
	"gitlab.com/leononame/advent-of-code-2018/pkg/aoc/2018/day02"
	"gitlab.com/leononame/advent-of-code-2018/pkg/aoc/2018/day03"
	"gitlab.com/leononame/advent-of-code-2018/pkg/aoc/2018/day04"
	"gitlab.com/leononame/advent-of-code-2018/pkg/aoc/2018/day05"
	"gitlab.com/leononame/advent-of-code-2018/pkg/aoc/2018/day06"
	"gitlab.com/leononame/advent-of-code-2018/pkg/aoc/2018/day07"
	"gitlab.com/leononame/advent-of-code-2018/pkg/aoc/2018/day08"
	"gitlab.com/leononame/advent-of-code-2018/pkg/aoc/2018/day09"
	"gitlab.com/leononame/advent-of-code-2018/pkg/aoc/2018/day10"
	"gitlab.com/leononame/advent-of-code-2018/pkg/aoc/2018/day11"
	"gitlab.com/leononame/advent-of-code-2018/pkg/aoc/2018/day12"
	"gitlab.com/leononame/advent-of-code-2018/pkg/aoc/2018/day13"
	"gitlab.com/leononame/advent-of-code-2018/pkg/aoc/2018/day14"
	"gitlab.com/leononame/advent-of-code-2018/pkg/aoc/2018/day15"
	"gitlab.com/leononame/advent-of-code-2018/pkg/aoc/2018/day16"
	"gitlab.com/leononame/advent-of-code-2018/pkg/aoc/2018/day17"
	"gitlab.com/leononame/advent-of-code-2018/pkg/aoc/2018/day18"
	"gitlab.com/leononame/advent-of-code-2018/pkg/aoc/2018/day19"
	"gitlab.com/leononame/advent-of-code-2018/pkg/aoc/2018/day20"
	"gitlab.com/leononame/advent-of-code-2018/pkg/aoc/2018/day21"
	"gitlab.com/leononame/advent-of-code-2018/pkg/aoc/2018/day22"
	"gitlab.com/leononame/advent-of-code-2018/pkg/aoc/2018/day23"
	"gitlab.com/leononame/advent-of-code-2018/pkg/aoc/2018/day24"
	"gitlab.com/leononame/advent-of-code-2018/pkg/aoc/2018/day25"
)

func main() {
	// Setup challenge
	config := aoc.Setup()
	if config.All {
		RunAll(config)
		os.Exit(0)
	}

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

	p := printer.New()
	p.AppendResult(r, config.Year, config.Day)
	p.Flush()
}

func RunAll(config *aoc.Config) {
	p := printer.New()
	challenges := runners[config.Year]
	for i, ch := range challenges {
		if ch == nil {
			continue
		}
		config.Day = i
		config.SetToDefaultFilePath()
		config.ReadFile()
		result := ch(config)
		p.AppendResult(result, config.Year, i)
	}
	p.Flush()
}

var runners = map[int][]func(config *aoc.Config) aoc.Result{
	2018: {
		0:  nil,
		1:  day01.Run,
		2:  day02.Run,
		3:  day03.Run,
		4:  day04.Run,
		5:  day05.Run,
		6:  day06.Run,
		7:  day07.Run,
		8:  day08.Run,
		9:  day09.Run,
		10: day10.Run,
		11: day11.Run,
		12: day12.Run,
		13: day13.Run,
		14: day14.Run,
		15: day15.Run,
		16: day16.Run,
		17: day17.Run,
		18: day18.Run,
		19: day19.Run,
		20: day20.Run,
		21: day21.Run,
		22: day22.Run,
		23: day23.Run,
		24: day24.Run,
		25: day25.Run,
	},
}
