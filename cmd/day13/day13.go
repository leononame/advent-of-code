package main

import (
	"errors"
	"fmt"
	"sort"

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
	var err error
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
			// Return position error, but only if first error
			if err == nil {
				err = errors.New(
					fmt.Sprintf("Cart has crashed at position %d,%d", int(real(c.position)), int(imag(c.position))))
			}
		} else {
			p.positions[c.position] = c
		}
		c.checkDir(p.track[c.position])
	}
	p.carts = removeNil(p.carts)
	return err
}

func main() {
	fmt.Println("Advent of Code 2018, ", version.Str)
	fmt.Println("Challenge: 2018-13")
	input := util.GetInput("input")
	p := parse(input)
	for {
		if err := p.Tick(); err != nil {
			fmt.Println(err.Error())
			break
		}
	}
	// part 2
	for {
		err := p.Tick()
		if err != nil && len(p.carts) == 1 {
			fmt.Printf("The position of the last remaining cart is: %d,%d\n",
				int(real(p.carts[0].position)),
				int(imag(p.carts[0].position)),
			)
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
