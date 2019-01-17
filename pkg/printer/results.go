package printer

import (
	"fmt"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
	"gitlab.com/leononame/advent-of-code-2018/pkg/aoc"
)

type Printer struct {
	table   *tablewriter.Table
	counter int
}

func (p *Printer) AppendResult(r aoc.Result, year, day int) {
	y := strconv.Itoa(year)
	d := fmt.Sprintf("%02d", day)

	// Separator if needed
	if p.counter > 0 {
		p.table.Append([]string{"", "", "", "", ""})
	}
	// Write data
	p.table.Append([]string{y, d, "Parse", fmt.Sprint(r.ParseTime), ""})
	p.table.Append([]string{"", "", "Part 1", fmt.Sprint(r.Duration1), fmt.Sprint(r.Solution1)})
	p.table.Append([]string{"", "", "Part 2", fmt.Sprint(r.Duration2), fmt.Sprint(r.Solution2)})
	p.counter++
}

func (p *Printer) Flush() {
	p.table.Render()
}

func New() *Printer {
	table := tablewriter.NewWriter(os.Stdout)
	// Headers
	table.SetHeader([]string{"Year", "Day", "Step", "Execution Time", "Solution"})
	// No border
	table.SetBorder(false)
	// Alignments
	table.SetColumnAlignment([]int{tablewriter.ALIGN_RIGHT,
		tablewriter.ALIGN_RIGHT,
		tablewriter.ALIGN_LEFT,
		tablewriter.ALIGN_RIGHT,
		tablewriter.ALIGN_RIGHT})
	// Colors
	table.SetHeaderColor(tablewriter.Colors{tablewriter.Bold, tablewriter.FgMagentaColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgMagentaColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgMagentaColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgMagentaColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgMagentaColor})
	table.SetColumnColor(tablewriter.Colors{tablewriter.Normal, tablewriter.FgYellowColor},
		tablewriter.Colors{tablewriter.Normal, tablewriter.FgYellowColor},
		tablewriter.Colors{tablewriter.Normal, tablewriter.FgCyanColor},
		tablewriter.Colors{tablewriter.Normal, tablewriter.FgBlackColor},
		tablewriter.Colors{tablewriter.Normal, tablewriter.FgGreenColor})
	return &Printer{table: table}
}
