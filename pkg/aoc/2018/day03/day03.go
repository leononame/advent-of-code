package day03

import (
	"fmt"
	"time"

	"gitlab.com/leononame/advent-of-code-2018/pkg/geo/points"

	"gitlab.com/leononame/advent-of-code-2018/pkg/aoc"

	"gitlab.com/leononame/advent-of-code-2018/pkg/geo/rect"
)

func Run(c *aoc.Config) (result aoc.Result) {
	t0 := time.Now()
	rs := parse(c.Input)
	result.ParseTime = time.Since(t0)

	t1 := time.Now()
	result.Solution1 = part1(rs)
	result.Duration1 = time.Since(t1)

	t2 := time.Now()
	result.Solution2 = part2(rs)
	result.Duration2 = time.Since(t2)
	return
}

type rectangle struct {
	id int
	rect.Rectangle
	points []points.Classic
}

// calcPoints fills the rectangle's points slice with all the points the rectangle covers
func (r *rectangle) calcPoints() {
	size := (r.Max.X - r.Min.X + 1) * (r.Max.Y - r.Min.Y + 1)
	var cs = make([]points.Classic, size)
	idx := 0
	for i := r.Min.X; i <= r.Max.X; i++ {
		for j := r.Min.Y; j <= r.Max.Y; j++ {
			cs[idx] = points.NewClassic(i, j)
			idx++
		}
	}
	r.points = cs
}

// check returns true if for the given list of counts the rectangle doesn't overlap
func (r *rectangle) check(counts [][]int) bool {
	for _, p := range r.points {
		// no overlap if count is exactly 1 for each point
		if counts[p.X][p.Y] != 1 {
			return false
		}
	}
	return true
}

func parse(input []string) []rectangle {
	var rectangles []rectangle
	for _, l := range input {
		r := rectangle{}
		fmt.Sscanf(l, "#%d @ %d,%d: %dx%d", &r.id, &r.Min.X, &r.Min.Y, &r.Max.X, &r.Max.Y)
		r.Max.X += r.Min.X - 1
		r.Max.Y += r.Min.Y - 1
		r.calcPoints()
		rectangles = append(rectangles, r)
	}
	return rectangles
}

func part1(rs []rectangle) int {
	counts := count(rs)
	c := 0
	for _, x := range counts {
		for _, y := range x {
			if y > 1 {
				c++
			}
		}
	}
	return c
}

func part2(rs []rectangle) int {
	counts := count(rs)
	for _, r := range rs {
		if r.check(counts) {
			return r.id
		}
	}
	return 0
}

// count counts how many rectangles are at a given point.
func count(rs []rectangle) [][]int {
	// Calculate size of slice so no extra allocations are needed during counting
	mx, my := 0, 0
	for _, r := range rs {
		if r.Max.X > mx {
			mx = r.Max.X
		}
		if r.Max.Y > my {
			my = r.Max.Y
		}
	}

	counter := make([][]int, mx+1)
	for _, r := range rs {
		for _, p := range r.points {
			if counter[p.X] == nil {
				counter[p.X] = make([]int, my+1)
			}
			counter[p.X][p.Y]++
		}
	}
	return counter
}
