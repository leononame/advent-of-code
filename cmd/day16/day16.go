package main

import (
	"fmt"
	"strings"

	"gitlab.com/leononame/advent-of-code-2018/pkg/util"
	"gitlab.com/leononame/advent-of-code-2018/pkg/version"
)

type registers [4]int
type instruction struct {
	code   int
	params [3]int
}
type capture struct {
	before, after registers
	i             instruction
}
type operation func(rs registers, a, b, c int) registers

var oplist = [16]operation{addr, addi, mulr, muli, banr, bani, borr, bori, setr, seti, eqir, eqri, eqrr, gtir, gtri, gtrr}
var opcodes = [16]operation{}

func main() {
	fmt.Println("Advent of Code 2018, ", version.Str)
	fmt.Println("Challenge: 2018-16")
	input := util.GetInput("input")
	cs, is := parse(input)
	combinations := countAndCombine(cs)
	mapOpcodes(cs, combinations)
	fmt.Println("Part 2", run(is))
}

func run(is []*instruction) registers {
	var rs registers
	for _, i := range is {
		rs = opcodes[i.code](rs, i.params[0], i.params[1], i.params[2])
	}
	return rs
}

func mapOpcodes(cs []*capture, combinations [16][16]bool) {
	found := 0
	for found < 16 {
		for opcode := range combinations {
			idx, count := 0, 0
			for j, bad := range combinations[opcode] {
				if !bad {
					idx = j
					count++
				}
			}
			if count == 1 {
				opcodes[opcode] = oplist[idx]
				found++
				for i := 0; i < 16; i++ {
					combinations[i][idx] = true
				}
			}
		}
	}
}

func countAndCombine(cs []*capture) [16][16]bool {
	count := 0
	var combinations [16][16]bool
	for _, c := range cs {
		success := 0
		for i, op := range oplist {
			if op(c.before, c.i.params[0], c.i.params[1], c.i.params[2]) == c.after {
				success++
			} else {
				combinations[c.i.code][i] = true
			}
		}
		if success > 2 {
			count++
		}
	}
	fmt.Println("Part 1:", count)
	return combinations
}

func parse(input []string) ([]*capture, []*instruction) {
	parts := strings.Split(strings.Join(input, "\n"), "\n\n\n\n")
	p1 := strings.Split(parts[0], "\n")
	p2 := strings.Split(parts[1], "\n")
	_ = p2
	var cs []*capture
	var is []*instruction
	for i := 0; i < len(p1); i += 4 {
		var c capture
		fmt.Sscanf(p1[i], "Before: [%d, %d, %d, %d]", &c.before[0], &c.before[1], &c.before[2], &c.before[3])
		fmt.Sscanf(p1[i+1], "%d %d %d %d", &c.i.code, &c.i.params[0], &c.i.params[1], &c.i.params[2])
		fmt.Sscanf(p1[i+2], "After: [%d, %d, %d, %d]", &c.after[0], &c.after[1], &c.after[2], &c.after[3])
		cs = append(cs, &c)
	}
	for i := 0; i < len(p2); i++ {
		var in instruction
		fmt.Sscanf(p2[i], "%d %d %d %d", &in.code, &in.params[0], &in.params[1], &in.params[2])
		is = append(is, &in)
	}
	return cs, is
}

func addr(rs registers, a, b, c int) registers {
	rs[c] = rs[a] + rs[b]
	return rs
}

func addi(rs registers, a, b, c int) registers {
	rs[c] = rs[a] + b
	return rs
}

func mulr(rs registers, a, b, c int) registers {
	rs[c] = rs[a] * rs[b]
	return rs
}

func muli(rs registers, a, b, c int) registers {
	rs[c] = rs[a] * b
	return rs
}

func banr(rs registers, a, b, c int) registers {
	rs[c] = rs[a] & rs[b]
	return rs
}

func bani(rs registers, a, b, c int) registers {
	rs[c] = rs[a] & b
	return rs
}

func borr(rs registers, a, b, c int) registers {
	rs[c] = rs[a] | rs[b]
	return rs
}

func bori(rs registers, a, b, c int) registers {
	rs[c] = rs[a] | b
	return rs
}

func setr(rs registers, a, b, c int) registers {
	rs[c] = rs[a]
	return rs
}

func seti(rs registers, a, b, c int) registers {
	rs[c] = a
	return rs
}

func gtir(rs registers, a, b, c int) registers {
	if a > rs[b] {
		rs[c] = 1
	} else {
		rs[c] = 0
	}
	return rs
}

func gtri(rs registers, a, b, c int) registers {
	if rs[a] > b {
		rs[c] = 1
	} else {
		rs[c] = 0
	}
	return rs
}

func gtrr(rs registers, a, b, c int) registers {
	if rs[a] > rs[b] {
		rs[c] = 1
	} else {
		rs[c] = 0
	}
	return rs
}

func eqir(rs registers, a, b, c int) registers {
	if a == rs[b] {
		rs[c] = 1
	} else {
		rs[c] = 0
	}
	return rs
}

func eqri(rs registers, a, b, c int) registers {
	if b == rs[a] {
		rs[c] = 1
	} else {
		rs[c] = 0
	}
	return rs
}

func eqrr(rs registers, a, b, c int) registers {
	if rs[a] == rs[b] {
		rs[c] = 1
	} else {
		rs[c] = 0
	}
	return rs
}
