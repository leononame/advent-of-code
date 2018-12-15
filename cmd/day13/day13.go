package main

import (
	"fmt"
	"sort"

	"github.com/juju/errors"

	"gitlab.com/leononame/advent-of-code-2018/pkg/util"
	"gitlab.com/leononame/advent-of-code-2018/pkg/version"
)

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

func (p *plan) Tick() error {
	// Sort myself
	sort.Sort(p)

	for _, c := range p.carts {
		delete(p.positions, c.position)
		c.position += c.direction
		if p.positions[c.position] != nil {
			return errors.New(
				fmt.Sprintf("Cart has crashed at position %f, %f", real(c.position), imag(c.position)))
		}
		p.positions[c.position] = c
		c.checkDir(p.track[c.position])
	}
	return nil
}

func main() {
	fmt.Println("Advent of Code 2018, ", version.Str)
	fmt.Println("Challenge: 2018-13")
	input := util.GetInput("input/day13")
	p := parse(input)
	for {
		if err := p.Tick(); err != nil {
			fmt.Println(err.Error())
			return
		}
	}
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
