package main

import (
	"fmt"
	"strings"

	"gitlab.com/leononame/advent-of-code-2018/pkg/util"

	"gitlab.com/leononame/advent-of-code-2018/pkg/version"
)

type point complex64

func newPoint(x, y int) point {
	var c = point(complex(float32(x), float32(y)))
	return c
}

func (p *point) x() int {
	return int(real(complex64(*p)))
}

func (p *point) y() int {
	return int(imag(complex64(*p)))
}

func (p *point) next() point {
	return newPoint(p.x(), p.y()+1)
}

func (p *point) left() point {
	return newPoint(p.x()-1, p.y())
}

func (p *point) right() point {
	return newPoint(p.x()+1, p.y())
}

func (p *point) to(pp point) []point {
	var ps []point
	sx := min(p.x(), pp.x())
	sy := min(p.y(), pp.y())
	ex := max(p.x(), pp.x())
	ey := max(p.y(), pp.y())
	for x := sx; x <= ex; x++ {
		for y := sy; y <= ey; y++ {
			ps = append(ps, newPoint(x, y))
		}
	}
	return ps
}

type layout struct {
	data       map[point]byte
	minx, maxx int
	miny, maxy int
}

func (l *layout) setMax(minx, maxx, miny, maxy int) {
	l.minx = min(minx, l.minx)
	l.maxx = max(maxx, l.maxx)
	l.miny = min(miny, l.miny)
	l.maxy = max(maxy, l.maxy)
}

func (l *layout) fill(p point) {
	if p.y() > l.maxy {
		return
	} else if l.isBarrier(p) {
		return
	}
	n := p.next()

	if l.isBarrier(n) {
		pl := p // point left
		// While to the left there is a floor and no wall, continue to fill with '|'
		for l.isBarrier(pl.next()) && !l.isBarrier(pl) {
			l.data[pl] = '|'
			pl = pl.left()
		}
		pr := p.right()
		s1, s2 := string(l.data[pr]), string(l.data[pr.next()])
		_, _ = s1, s2
		// While to the right there is a floor and no wall, continue to fill with '|'
		for l.isBarrier(pr.next()) && !l.isBarrier(pr) {
			l.data[pr] = '|'
			pr = pr.right()
			s1, s2 = string(l.data[pr]), string(l.data[pr.next()])
		}
		// If there is no floor at any of the exit points, the water fill flow out of there. Try both
		if !l.isBarrier(pr.next()) {
			l.fill(pr)
			if l.isBarrier(pr.next()) {
				l.fill(p)
			}
		}
		if !l.isBarrier(pl.next()) {
			l.fill(pl)
			if l.isBarrier(pl.next()) {
				l.fill(p)
			}

		}
		if l.isBarrier(pr.next()) && l.isBarrier(pl.next()) {
			// Fill up if there is no barrier
			pl = pl.right()
			pr = pr.left()
			for _, pp := range pr.to(pl) {
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

func (l *layout) isBarrier(p point) bool {
	return l.data[p] != 0 && l.data[p] != '|'
}

func (l *layout) String() string {
	var sb strings.Builder
	for y := l.miny; y <= l.maxy; y++ {
		sb.WriteString(fmt.Sprintf("%04d ", y))
		// X bounds increased by one so water can flow on the side
		for x := l.minx - 1; x <= l.maxx+1; x++ {
			p := newPoint(x, y)
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

func main() {
	fmt.Println("Advent of Code 2018, ", version.Str)
	fmt.Println("Challenge: 2018-17")
	input := util.GetInput("input")
	l := parse(input)
	fmt.Println(l.minx, l.maxx, l.miny, l.maxy)
	l.fill(newPoint(500, 0))
	s := l.String()
	sand := strings.Count(s, "|")
	water := strings.Count(s, "~")
	fmt.Println(l)
	fmt.Println("Part 1:", sand+water)
	fmt.Println("Part 2:", water)
}

func parse(input []string) *layout {
	l := layout{make(map[point]byte), 2000, 0, 2000, 0}
	for _, line := range input {
		var x1, x2, y1, y2 int
		if line[0] == 'x' {
			fmt.Sscanf(line, "x=%d, y=%d..%d", &x1, &y1, &y2)
			l.setMax(x1, x1, y1, y2)
			for _, y := range n(y1, y2) {
				l.data[newPoint(x1, y)] = '#'
			}
		} else {
			fmt.Sscanf(line, "y=%d, x=%d..%d", &y1, &x1, &x2)
			l.setMax(x1, x2, y1, y1)
			for _, x := range n(x1, x2) {
				l.data[newPoint(x, y1)] = '#'
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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
