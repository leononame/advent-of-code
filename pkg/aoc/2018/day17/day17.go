package day17

import (
	"fmt"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"gitlab.com/leononame/advent-of-code/pkg/aoc"

	"gitlab.com/leononame/advent-of-code/pkg/mmath"

	"gitlab.com/leononame/advent-of-code/pkg/geo"
)

var logger *logrus.Logger

func Run(c *aoc.Config) (result aoc.Result) {
	logger = c.Logger

	t0 := time.Now()
	l := parse(c.Input)
	result.ParseTime = time.Since(t0)

	t1 := time.Now()

	l.fill(geo.Point{500, 0})
	s := l.String()
	sand := strings.Count(s, "|")
	water := strings.Count(s, "~")

	result.Solution1 = sand + water
	result.Solution2 = water
	result.Duration1 = time.Since(t1)
	result.Duration2 = result.Duration1
	return
}

type layout struct {
	data       map[geo.Point]byte
	minx, maxx int
	miny, maxy int
}

func (l *layout) setMax(minx, maxx, miny, maxy int) {
	l.minx = mmath.Min(minx, l.minx)
	l.maxx = mmath.Max(maxx, l.maxx)
	l.miny = mmath.Min(miny, l.miny)
	l.maxy = mmath.Max(maxy, l.maxy)
}

func (l *layout) fill(p geo.Point) {
	if p.Y > l.maxy {
		return
	} else if l.isBarrier(p) {
		return
	}
	n := p.Down()

	if l.isBarrier(n) {
		pl := p // geo.Point left
		// While to the left there is a floor and no wall, continue to fill with '|'
		for l.isBarrier(pl.Down()) && !l.isBarrier(pl) {
			l.data[pl] = '|'
			pl = pl.Left()
		}
		pr := p.Right()
		s1, s2 := string(l.data[pr]), string(l.data[pr.Down()])
		_, _ = s1, s2
		// While to the right there is a floor and no wall, continue to fill with '|'
		for l.isBarrier(pr.Down()) && !l.isBarrier(pr) {
			l.data[pr] = '|'
			pr = pr.Right()
			s1, s2 = string(l.data[pr]), string(l.data[pr.Down()])
		}
		// If there is no floor at any of the exit geo.Points, the water fill flow out of there. Try both
		if !l.isBarrier(pr.Down()) {
			l.fill(pr)
			if l.isBarrier(pr.Down()) {
				l.fill(p)
			}
		}
		if !l.isBarrier(pl.Down()) {
			l.fill(pl)
			if l.isBarrier(pl.Down()) {
				l.fill(p)
			}

		}
		if l.isBarrier(pr.Down()) && l.isBarrier(pl.Down()) {
			// Fill up if there is no barrier
			pl = pl.Right()
			pr = pr.Left()
			for _, pp := range to(pr, pl) {
				l.data[pp] = '~'
			}
		}
	} else if l.data[p] == 0 {
		l.data[p] = '|'
		l.fill(n)
		if l.isBarrier(n) {
			l.fill(p)
		}
	}
}

func (l *layout) isBarrier(p geo.Point) bool {
	return l.data[p] != 0 && l.data[p] != '|'
}

func (l *layout) String() string {
	var sb strings.Builder
	for y := l.miny; y <= l.maxy; y++ {
		sb.WriteString(fmt.Sprintf("%04d ", y))
		// X bounds increased by one so water can flow on the side
		for x := l.minx - 1; x <= l.maxx+1; x++ {
			p := geo.Point{x, y}
			s := l.data[p]
			if s == 0 {
				sb.WriteByte(' ')
			} else {
				sb.WriteByte(s)
			}
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

func parse(input []string) *layout {
	l := layout{make(map[geo.Point]byte), 2000, 0, 2000, 0}
	for _, line := range input {
		var x1, x2, y1, y2 int
		if line[0] == 'x' {
			fmt.Sscanf(line, "x=%d, y=%d..%d", &x1, &y1, &y2)
			l.setMax(x1, x1, y1, y2)
			for _, y := range n(y1, y2) {
				l.data[geo.Point{x1, y}] = '#'
			}
		} else {
			fmt.Sscanf(line, "y=%d, x=%d..%d", &y1, &x1, &x2)
			l.setMax(x1, x2, y1, y1)
			for _, x := range n(x1, x2) {
				l.data[geo.Point{x, y1}] = '#'
			}
		}
	}
	return &l
}

func n(start, end int) []int {
	var r []int
	for i := start; i <= end; i++ {
		r = append(r, i)
	}
	return r
}

func to(p1, p2 geo.Point) []geo.Point {
	var ps []geo.Point
	sx := mmath.Min(p1.X, p2.X)
	sy := mmath.Min(p1.Y, p2.Y)
	ex := mmath.Max(p1.X, p2.X)
	ey := mmath.Max(p1.Y, p2.Y)
	for x := sx; x <= ex; x++ {
		for y := sy; y <= ey; y++ {
			ps = append(ps, geo.Point{x, y})
		}
	}
	return ps
}
