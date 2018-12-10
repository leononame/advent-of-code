package main

import (
	"fmt"
	"math"
	"strings"

	"gitlab.com/leononame/advent-of-code-2018/pkg/util"
	"gitlab.com/leononame/advent-of-code-2018/pkg/version"
)

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

type point struct {
	x, y   int64
	vx, vy int
}

type sky []*point

type image struct {
	sky
	maxX, minX, maxY, minY int64
	h, w                   int64
	age                    int
}

func (im *image) tick() {
	im.age++
	im.maxX, im.minX, im.maxY, im.minY = math.MinInt64, math.MaxInt64, math.MinInt64, math.MaxInt64
	for _, p := range im.sky {
		p.x += int64(p.vx)
		p.y += int64(p.vy)
		im.maxX = max(p.x, im.maxX)
		im.minX = min(p.x, im.minX)
		im.maxY = max(p.y, im.maxY)
		im.minY = min(p.y, im.minY)
	}
	im.h = im.maxY - im.minY
	im.w = im.maxX - im.minX
}

func (im *image) untick(count int) {
	im.age -= count
	im.maxX, im.minX, im.maxY, im.minY = math.MinInt64, math.MaxInt64, math.MinInt64, math.MaxInt64
	for _, p := range im.sky {
		p.x -= int64(count) * int64(p.vx)
		p.y -= int64(count) * int64(p.vy)
		im.maxX = max(p.x, im.maxX)
		im.minX = min(p.x, im.minX)
		im.maxY = max(p.y, im.maxY)
		im.minY = min(p.y, im.minY)
	}
	im.h = im.maxX - im.minX
	im.w = im.maxY - im.minY
}

func (im *image) print() {
	data := make([][]string, im.w+1)
	for i, _ := range data {
		data[i] = make([]string, im.h+1)
	}
	for i, _ := range data {
		for j, _ := range data[i] {
			data[i][j] = " "
		}
	}
	for _, p := range im.sky {
		data[p.y-im.minY][p.x-im.minX] = "#"
	}
	for _, row := range data {
		fmt.Println(strings.Join(row, ""))
	}
	fmt.Println("Age: ", im.age)
	return
}

func main() {
	fmt.Println("Advent of Code 2018, ", version.Str)
	fmt.Println("Challenge: 2018-10")
	input := util.GetInput("input/day10")
	im := parse(input)
	im.tick()

	for {
		old := im.h
		im.tick()

		if im.h > old {
			im.untick(1)
			im.print()
			return
		}
	}
}

func parse(input []string) *image {
	var im image
	for _, l := range input {
		var p point
		fmt.Sscanf(l, "position=<%d, %d> velocity=<%d, %d>", &p.x, &p.y, &p.vx, &p.vy)
		im.sky = append(im.sky, &p)
	}
	return &im
}
