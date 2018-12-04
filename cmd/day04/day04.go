package main

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"gitlab.com/leononame/advent-of-code-2018/pkg/util"
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

func main() {
	fmt.Println("Challenge:\t2018-04")
	input := util.GetInput("input/day04")
	sort.Strings(input)
	part1(input)
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

func (gd *guardData) calcBestMinute() int {
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
	return maxIdx
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

func part1(input []string) {
	data := parseData(input)
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
	bestMinute := gd.calcBestMinute()
	fmt.Printf("Part 1: Best guard is id %d.\nThey fell asleep for a total of %d minutes.\n", maxID, max)
	fmt.Printf("The best minute to sneak through is %d.\nThe product of m and g is %d.\n", bestMinute, bestMinute*maxID)
}
