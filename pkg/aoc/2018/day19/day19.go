package day19

import (
	"fmt"
	"math"
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
	rp1 := run(*p, registers{})
	result.Solution1 = rp1[0]
	result.Duration1 = time.Since(t1)

	t2 := time.Now()
	rp2 := part2(*p)
	result.Solution2 = rp2[0]
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

func part2(p program) registers {
	// These two operations set the jump of seti to 100, effectively halting the
	// program prematurely because it doesn't jump into the loop
	// Thus, we can run the program until the number that we need to sum the
	// divisors is calculated. Then, the program exits.
	p.is[26].params[0] = 100 // For part 1
	p.is[35].params[0] = 100 // For part 2
	// Run the program to calculate our number
	resp := run(p, registers{1})
	// Instruction 33 writes the number into our target register. Extract that
	// register
	reg := p.is[33].params[2]
	num := resp[reg]
	// Calculate sum of divisors of said number
	return registers{divSum(num)}
}

func divSum(n int) int {
	res := 1
	for i := 2; i < int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			res += (i + n/i)
		}
	}
	return res + n
}

func run(p program, rs registers) registers {
	ip := 0
	for ip < len(p.is) && ip >= 0 {
		i := p.is[ip]
		rs = i.op(rs, i.params[0], i.params[1], i.params[2])
		rs[p.ipReg]++
		ip = rs[p.ipReg]
		// fmt.Println(ip, "\t", runtime.FuncForPC(reflect.ValueOf(i.op).Pointer()).Name(), i.params, rs)
	}
	return rs
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
