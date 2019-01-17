package main

import (
	"fmt"
	"os"

	"gitlab.com/leononame/advent-of-code-2018/pkg/aoc/2018/day18"
	"gitlab.com/leononame/advent-of-code-2018/pkg/aoc/2018/day19"
	"gitlab.com/leononame/advent-of-code-2018/pkg/aoc/2018/day20"
	"gitlab.com/leononame/advent-of-code-2018/pkg/aoc/2018/day21"
	"gitlab.com/leononame/advent-of-code-2018/pkg/aoc/2018/day22"
	"gitlab.com/leononame/advent-of-code-2018/pkg/aoc/2018/day23"
	"gitlab.com/leononame/advent-of-code-2018/pkg/aoc/2018/day24"
	"gitlab.com/leononame/advent-of-code-2018/pkg/aoc/2018/day25"

	"github.com/olekukonko/tablewriter"
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

	writeResultToTable(r)
}

func writeResultToTable(r aoc.Result) {
	table := tablewriter.NewWriter(os.Stdout)
	// Headers
	table.SetHeader([]string{"Step", "Execution Time", "Solution"})
	// No border
	table.SetBorder(false)
	// Alignments
	table.SetColumnAlignment([]int{tablewriter.ALIGN_LEFT, tablewriter.ALIGN_RIGHT, tablewriter.ALIGN_RIGHT})
	// Colors
	table.SetHeaderColor(tablewriter.Colors{tablewriter.Bold, tablewriter.FgMagentaColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgMagentaColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgMagentaColor})
	table.SetColumnColor(tablewriter.Colors{tablewriter.Normal, tablewriter.FgCyanColor},
		tablewriter.Colors{tablewriter.Normal, tablewriter.FgBlackColor},
		tablewriter.Colors{tablewriter.Normal, tablewriter.FgGreenColor})
	// Write data
	table.Append([]string{"Parse", fmt.Sprint(r.ParseTime), ""})
	table.Append([]string{"Part 1", fmt.Sprint(r.Duration1), fmt.Sprint(r.Solution1)})
	table.Append([]string{"Part 2", fmt.Sprint(r.Duration2), fmt.Sprint(r.Solution2)})
	table.Render()
}

var runners = map[int]map[int]func(config *aoc.Config) aoc.Result{
	2018: {
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
