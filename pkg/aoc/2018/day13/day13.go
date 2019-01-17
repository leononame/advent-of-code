package day13

import (
	"fmt"
	"sort"
	"time"

	"github.com/sirupsen/logrus"
	"gitlab.com/leononame/advent-of-code/pkg/aoc"
)

var logger *logrus.Logger

func Run(c *aoc.Config) (result aoc.Result) {
	logger = c.Logger

	t0 := time.Now()
	p := parse(c.Input)
	result.ParseTime = time.Since(t0)

	t1 := time.Now()
	result.Solution1 = part1(p)
	result.Duration1 = time.Since(t1)

	t2 := time.Now()
	result.Solution2 = part2(p)
	result.Duration2 = time.Since(t2)
	return
}

func part1(p *plan) string {
	for crash := p.tick(); ; crash = p.tick() {
		if crash != "" {
			return crash
		}
	}
}

func part2(p *plan) string {
	for crash := p.tick(); ; crash = p.tick() {
		if crash != "" && len(p.carts) == 1 {
			pos := fmt.Sprintf("%d,%d",
				int(real(p.carts[0].position)),
				int(imag(p.carts[0].position)))
			logger.Debugf("The position of the last remaining cart is: %s", pos)
			return pos
		}
	}
}

const left complex64 = -1
const up complex64 = -1i
const right complex64 = 1
const down complex64 = 1i

type cart struct {
	position  complex64
	direction complex64
	turn      complex64
}

func (c *cart) checkDir(track rune) {
	if track == '+' {
		c.crossing()
	} else if track == '/' {
		c.direction = complex(-imag(c.direction), -real(c.direction))
	} else if track == '\\' {
		c.direction = complex(imag(c.direction), real(c.direction))
	}
}

func (c *cart) crossing() {
	c.direction *= c.turn
	if c.turn == 1i {
		c.turn = -1i
	} else if c.turn == -1i {
		c.turn = 1
	} else {
		c.turn = 1i
	}
}

type plan struct {
	carts     []*cart
	track     map[complex64]rune
	positions map[complex64]*cart
}

func (p *plan) Len() int {
	return len(p.carts)
}

func (p *plan) Less(i, j int) bool {
	c1 := p.carts[i]
	c2 := p.carts[j]
	return imag(c1.position) < imag(c2.position) ||
		(imag(c1.position) == imag(c2.position) && real(c1.position) < real(c2.position))
}

func (p *plan) Swap(i, j int) {
	p.carts[i], p.carts[j] = p.carts[j], p.carts[i]
}

func (p *plan) tick() string {
	// Sort myself
	sort.Sort(p)
	var pos string
	for i, c := range p.carts {
		if c == nil {
			continue
		}
		// c.position == complex(112.0, 9.0)
		delete(p.positions, c.position)
		c.position += c.direction
		if p.positions[c.position] != nil {
			// Remove carts
			p.carts[i] = nil
			p.carts[find(p.carts, p.positions[c.position])] = nil
			delete(p.positions, c.position)
			// Return position crash, but only if first time
			if pos == "" {
				pos = fmt.Sprintf("%d,%d", int(real(c.position)), int(imag(c.position)))
			}
		} else {
			p.positions[c.position] = c
		}
		c.checkDir(p.track[c.position])
	}
	p.carts = removeNil(p.carts)
	return pos
}

var directions = map[rune]complex64{
	'<': left,
	'>': right,
	'^': up,
	'v': down,
}

func parse(input []string) *plan {
	track := make(map[complex64]rune)
	var carts []*cart
	positions := make(map[complex64]*cart)
	for i := range input {
		for j, el := range []rune(input[i]) {
			pos := complex(float32(j), float32(i))
			track[pos] = el
			if directions[el] != 0 {
				c := &cart{pos, directions[el], -1i}
				carts = append(carts, c)
				positions[c.position] = c
			}
		}
	}
	return &plan{carts, track, positions}
}

func find(carts []*cart, c *cart) int {
	for i, v := range carts {
		if v == c {
			return i
		}
	}
	return -1
}

func removeNil(carts []*cart) []*cart {
	ret := carts
	for i := find(ret, nil); i != -1; i = find(ret, nil) {
		ret = append(ret[:i], ret[i+1:]...)
	}
	return ret
}
