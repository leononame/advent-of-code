package day21

import (
	"fmt"
	"reflect"
	"runtime"
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
	result.Solution1 = part1(*p)
	result.Duration1 = time.Since(t1)

	t2 := time.Now()
	result.Solution2 = part2(*p)
	result.Duration2 = time.Since(t2)
	return
}

type registers [6]int
type instruction struct {
	op     operation
	params [3]int
}
type operation func(rs registers, a, b, c int) registers

var oplist = map[string]operation{
	"addr": addr,
	"addi": addi,
	"mulr": mulr,
	"muli": muli,
	"banr": banr,
	"bani": bani,
	"borr": borr,
	"bori": bori,
	"setr": setr,
	"seti": seti,
	"eqir": eqir,
	"eqri": eqri,
	"eqrr": eqrr,
	"gtir": gtir,
	"gtri": gtri,
	"gtrr": gtrr,
}

type program struct {
	ipReg int
	is    []*instruction
}

func part1(p program) int {
	// Looking at the source code, register 0 is only used once. It's used to
	// compare the value with eqrr to another register. In my case, with R5.
	// For my input, this would be instruction 29. Let's just run the program
	// regularly until the comparison hits. Then we can extract that number
	res, i := run(p, registers{}, eqrr)
	reg := i.params[0]
	return res[reg]
}

func part2(p program) int {
	// Magic number, our input
	magic := p.is[7].params[0]
	// Algorithm
	return rev10(magic)
}

func run(p program, rs registers, breakOn operation) (registers, *instruction) {
	ip := 0
	for ip < len(p.is) && ip >= 0 {
		i := p.is[ip]
		rs = i.op(rs, i.params[0], i.params[1], i.params[2])
		logger.Debug(ip, "\t", runtime.FuncForPC(reflect.ValueOf(i.op).Pointer()).Name(), i.params, rs)
		rs[p.ipReg]++
		ip = rs[p.ipReg]

		if reflect.ValueOf(i.op).Pointer() == reflect.ValueOf(breakOn).Pointer() {
			return rs, i
		}
	}
	return rs, nil
}

func parse(input []string) *program {
	var p program
	fmt.Sscanf(input[0], "#ip %d", &p.ipReg)
	for i := 1; i < len(input); i++ {
		var in instruction
		fname := ""
		fmt.Sscanf(input[i], "%s %d %d %d", &fname, &in.params[0], &in.params[1], &in.params[2])
		in.op = oplist[fname]
		p.is = append(p.is, &in)
	}
	return &p
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

// Refactor the last program (rev09) into a function that looks for a repeat
func rev10(input int) int {
	var R3, R5, last int
	// List of values for R0 which would exit our program
	repeats := make(map[int]bool)
	for {
		R3 = R5 | 65536 // R3 = 65536
		for R5 = input; ; R3 /= 256 {
			R5 = ((((R5 + (R3 & 255)) & 16777215) * 65899) & 16777215)
			if R3 < 256 {
				break
			}
		}

		if repeats[R5] {
			return last
		}
		last = R5
		repeats[R5] = true
	}
}
