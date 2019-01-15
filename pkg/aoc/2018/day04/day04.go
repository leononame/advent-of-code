package day04

import (
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/logrusorgru/aurora"

	"github.com/sirupsen/logrus"

	"gitlab.com/leononame/advent-of-code-2018/pkg/aoc"
)

type date struct {
	year  int
	month int
	day   int
}

type hour struct {
	hour   int
	minute int
}

var logger *logrus.Logger

func Run(c *aoc.Config) (result aoc.Result) {
	logger = c.Logger

	t0 := time.Now()
	sort.Strings(c.Input)
	data := parseData(c.Input)
	result.ParseTime = time.Since(t0)

	t1 := time.Now()
	result.Solution1 = part1(data)
	result.Duration1 = time.Since(t1)

	t2 := time.Now()
	result.Solution2 = part2(data)
	result.Duration2 = time.Since(t2)
	return

}

type guardData map[date]*[60]int

func (gd *guardData) calcMinutesAsleep() int {
	sum := 0
	for _, day := range *gd {
		for _, isAsleep := range day {
			sum += isAsleep
		}
	}

	return sum
}

func (gd *guardData) calcBestMinute() (minute int, value int) {
	sumByMinute := [60]int{}
	for _, day := range *gd {
		for minute, isAsleep := range day {
			sumByMinute[minute] += isAsleep
		}
	}
	max := sumByMinute[0]
	maxIdx := 0
	for i, v := range sumByMinute {
		if v > max {
			max = v
			maxIdx = i
		}
	}
	return maxIdx, max
}

func parseData(input []string) map[int]guardData {
	// Data contains for each guard a map of minutes for each day
	data := make(map[int]guardData)
	guard, _ := regexp.Compile(`Guard #(\d+) begins shift`)
	// Current guard id
	id := 0
	for _, l := range input {
		d, h := parseDate(l)
		// If guard, parse guard id
		if strings.Contains(l, "Guard") {
			id, _ = strconv.Atoi(guard.FindStringSubmatch(l)[1])
			if data[id] == nil {
				data[id] = make(map[date]*[60]int)
			}
			continue
		}
		// Get the action. 1 == falls asleep, 0 == wakes up
		a := action(l)
		// If our data is currently nil, create new array
		if data[id][d] == nil {
			data[id][d] = &[60]int{}
		}
		// Current day
		curr := data[id][d]
		curr[h.minute] = a
		// If the action is wake up, go back and fill up with state asleep backwards
		if a == 0 {
			fillMinutesBackwards(curr, h.minute-1)
		}
	}
	return data
}

// Fills a 60 minute int array backwards with 1s until a 1 is found
// i.e.
// minutes = 0 0 0 0 0 1 0 0 0 0 0 0
// idx = 10
// Results in
// minutes = 0 0 0 0 0 1 1 1 1 1 0 0
func fillMinutesBackwards(minutes *[60]int, idx int) {
	for i := idx; i > 0; i-- {
		if minutes[i] == 1 {
			return
		}
		minutes[i] = 1
	}
}

func action(s string) int {
	if strings.Contains(s, "wakes up") {
		return 0
	}
	return 1
}

func parseDate(s string) (date, hour) {
	time := regexp.MustCompile(`(\d+)-(\d+)-(\d+) (\d+):(\d+)`)
	data := time.FindStringSubmatch(s)
	var vals []int
	for i := 1; i < len(data); i++ {
		v, _ := strconv.Atoi(data[i])
		vals = append(vals, v)
	}
	return date{vals[0], vals[1], vals[2]}, hour{vals[3], vals[4]}
}

func part1(data map[int]guardData) int {
	max := 0
	maxID := 0
	for g, d := range data {
		s := d.calcMinutesAsleep()
		if s > max {
			max = s
			maxID = g
		}
	}
	gd := data[maxID]
	bestMinute, bestValue := gd.calcBestMinute()
	logger.Debugf("Part 1: Best guard is id %d. They fell asleep for a total of %d minutes.", aurora.Green(maxID), max)
	logger.Debugf("The best minute to sneak through is %d. The guard was asleep on %d days during that minute.", aurora.Green(bestMinute), bestValue)
	logger.Debugf("The product of m and g is %d.\n", aurora.Green(bestMinute*maxID))
	return bestMinute * maxID
}

func part2(data map[int]guardData) int {
	max := 0
	maxID := 0
	maxMinute := 0
	for g, d := range data {
		bm, v := d.calcBestMinute()
		if v > max {
			max = v
			maxID = g
			maxMinute = bm
		}
	}

	logger.Debugf("Part 2: The guard that is most frequently asleep on the same minute is guard #%d.\n", aurora.Green(maxID))
	logger.Debugf("He spent minute %d asleep on %d days. The product of the minute and the guard ID is %d.\n", aurora.Green(maxMinute), max, aurora.Green(maxMinute*maxID))
	return maxMinute * maxID
}
