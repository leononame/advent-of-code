package main

import (
	"fmt"
	"os"

	"gitlab.com/leononame/advent-of-code-2018/pkg/aoc/2018/day05"

	"gitlab.com/leononame/advent-of-code-2018/pkg/aoc/2018/day04"

	"github.com/olekukonko/tablewriter"

	"github.com/logrusorgru/aurora"

	"gitlab.com/leononame/advent-of-code-2018/pkg/aoc"
	"gitlab.com/leononame/advent-of-code-2018/pkg/aoc/2018/day01"
	"gitlab.com/leononame/advent-of-code-2018/pkg/aoc/2018/day02"
	"gitlab.com/leononame/advent-of-code-2018/pkg/aoc/2018/day03"
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

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Step", "Execution Time", "Solution"})
	table.SetBorder(false)
	table.SetAlignment(tablewriter.ALIGN_RIGHT)
	writeResult(table, r)
	table.Render()
}

func writeResult(table *tablewriter.Table, r aoc.Result) {
	table.Append([]string{
		fmt.Sprint(aurora.Gray("Parse")),
		fmt.Sprint(aurora.Black(r.ParseTime)),
		fmt.Sprint(aurora.Green("")),
	})
	table.Append([]string{
		fmt.Sprint(aurora.Gray("Part1")),
		fmt.Sprint(aurora.Black(r.Duration1)),
		fmt.Sprint(aurora.Green(r.Solution1)),
	})
	table.Append([]string{
		fmt.Sprint(aurora.Gray("Part2")),
		fmt.Sprint(aurora.Black(r.Duration2)),
		fmt.Sprint(aurora.Green(r.Solution2)),
	})
}

var runners = map[int]map[int]func(config *aoc.Config) aoc.Result{
	2018: {
		1: day01.Run,
		2: day02.Run,
		3: day03.Run,
		4: day04.Run,
		5: day05.Run,
	},
}
