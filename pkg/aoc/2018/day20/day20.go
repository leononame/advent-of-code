package day20

import (
	"os"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"gitlab.com/leononame/advent-of-code/pkg/aoc"
)

var logger *logrus.Logger

func Run(c *aoc.Config) (result aoc.Result) {
	logger = c.Logger

	t0 := time.Now()
	n := parse(c.Input[0])
	result.ParseTime = time.Since(t0)

	t1 := time.Now()
	result.Solution1 = n.furthest()
	result.Duration1 = time.Since(t1)

	t2 := time.Now()
	result.Solution2 = n.countRooms(1000)
	result.Duration2 = time.Since(t2)
	return
}

const left point = -1
const up point = -1i
const right point = 1
const down point = 1i

var dirs = map[byte]point{
	'N': up,
	'W': left,
	'S': down,
	'E': right,
}
var doors = map[byte]byte{
	'N': '-',
	'S': '-',
	'W': '|',
	'E': '|',
}

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

// North Pole
type np struct {
	data                   map[point]byte
	rooms                  map[point]int
	minx, miny, maxx, maxy int
}

func (n *np) setMax(p point) {
	x := p.x()
	y := p.y()
	n.minx = min(x, n.minx)
	n.maxx = max(x, n.maxx)
	n.miny = min(y, n.miny)
	n.maxy = max(y, n.maxy)
}

func (n *np) furthest() int {
	d := 0
	for _, v := range n.rooms {
		if v > d {
			d = v
		}
	}
	return d
}

func (n *np) countRooms(minDistance int) int {
	c := 0
	for _, v := range n.rooms {
		if v >= minDistance {
			c++
		}
	}
	return c
}

func (n *np) String() string {
	var sb strings.Builder
	for y := n.miny - 1; y < n.maxy+2; y++ {
		for x := n.minx - 1; x < n.maxx+2; x++ {
			p := newPoint(x, y)
			el := n.data[p]
			if el == 0 {
				sb.WriteString("#")
			} else {
				sb.WriteByte(el)
			}
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

func (n *np) fill(pos point, path *[]byte, dist, idx *int) {
	current := pos
	dd := *dist
	for ; *idx < len(*path); *idx++ {
		el := (*path)[*idx]
		s := string(el)
		_ = s
		switch el {
		case '(':
			*idx++
			n.fill(current, path, dist, idx)
		case ')':
			return
		case '|':
			current = pos
			*dist = dd
		default:
			current += 2 * dirs[el]
			if _, ok := n.data[current]; ok {
				*dist = n.rooms[current]
			} else {
				n.setMax(current)
				n.data[current] = '.'
				n.data[current-dirs[el]] = doors[el]
				*dist++
				n.rooms[current] = *dist
			}
		}
	}
}

func parse(input string) *np {
	r := []byte(input[1 : len(input)-1])
	n := np{data: make(map[point]byte), rooms: make(map[point]int)}
	n.data[0] = 'X'
	var distance, idx int
	n.fill(0, &r, &distance, &idx)
	return &n
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

func test(input string, val int) {
	n := parse(input)
	if n.furthest() != val {
		os.Exit(1)
	}
}
