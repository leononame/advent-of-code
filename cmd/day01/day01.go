package main

import (
	"bufio"
	"fmt"
	"gitlab.com/leononame/advent-of-code-2018/pkg/version"
	"os"
	"strconv"
)

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	result := 0
	println("Verison:\t" + version.Str)
	println("Challenge:\t2018-01-01")

	if len(os.Args) == 1 {
		println("Pass the input file as parameter")
		os.Exit(1)
	}

	path := os.Args[1]
	f, err := os.Open(path)
	checkErr(err)

	s := bufio.NewScanner(f)
	for s.Scan() {
		v := calcValue(s.Text())
		result += v
		fmt.Printf("Sum: %d\t, Value; %d\n", result, v)
	}
	fmt.Printf("Result: %d\n", result)
}

func calcValue(line string) int {
	v, err := strconv.Atoi(line[1:])
	checkErr(err)
	if line[0] == '-' {
		v = -v
	}
	return v
}
