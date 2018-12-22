package main

import (
	"fmt"
	"regexp"
	"strconv"

	"gitlab.com/leononame/advent-of-code-2018/pkg/util"
)

type coordinate struct {
	x int
	y int
}

type rectangle struct {
	id int
	sx int
	sy int
	coordinate
}

func main() {
	fmt.Println("Challenge:\t2018-03")
	input := util.GetInput("input")

	// Run function
	part1(input)
	part2(input)
}

func parseInput(input []string) *[]rectangle {
	var rectangles []rectangle
	for _, l := range input {
		rectangles = append(rectangles, *parseLine(l))
	}
	return &rectangles
}

func parseLine(line string) *rectangle {
	r := regexp.MustCompile(`^#(\d+) @ (\d+),(\d+): (\d+)x(\d+)$`)
	m := r.FindStringSubmatch(line)
	var vals []int
	for i := 1; i < len(m); i++ {
		v, _ := strconv.Atoi(m[i])
		vals = append(vals, v)
	}
	return &rectangle{
		id: vals[0],
		sx: vals[3],
		sy: vals[4],
		coordinate: coordinate{
			x: vals[1],
			y: vals[2],
		},
	}
}

func calcCoordinates(r rectangle) []coordinate {
	var cs []coordinate
	for i := 0; i < r.sx; i++ {
		for j := 0; j < r.sy; j++ {
			cs = append(cs, coordinate{i + r.x, j + r.y})
		}
	}
	return cs
}

func calcCoordinateCounts(input []string) map[coordinate]int {
	counter := make(map[coordinate]int)
	rs := parseInput(input)
	for _, r := range *rs {
		cs := calcCoordinates(r)
		for _, c := range cs {
			counter[c]++
		}
	}
	return counter
}

func part1(input []string) {
	counts := calcCoordinateCounts(input)

	// count coordinates with at least double count
	c := 0
	for _, v := range counts {
		if v > 1 {
			c++
		}
	}

	fmt.Printf("Coordinates with overlap: %d\n", c)
}

func part2(input []string) {
	counts := calcCoordinateCounts(input)
	rs := parseInput(input)
	for _, r := range *rs {
		cs := calcCoordinates(r)
		if checkCoordinagtes(&cs, &counts) {
			fmt.Printf("Rectangle with ID %d has no overlaps!\n", r.id)
		}
	}
}

func checkCoordinagtes(cs *[]coordinate, data *map[coordinate]int) bool {
	for _, c := range *cs {
		if (*data)[c] != 1 {
			return false
		}
	}
	return true
}
