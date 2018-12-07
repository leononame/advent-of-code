package util

import (
	"bufio"
	"os"
)

// I know, util package is an antipattern. This code here is so small I don't care

func CheckErr(e error) {
	if e != nil {
		panic(e)
	}
}

func GetInput(path string) []string {
	f, err := os.Open(path)
	CheckErr(err)

	s := bufio.NewScanner(f)
	var lines []string
	for s.Scan() {
		lines = append(lines, s.Text())
	}
	return lines
}

func Abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}
