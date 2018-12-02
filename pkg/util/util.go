package util

import (
	"gitlab.com/leononame/advent-of-code-2018/pkg/version"
	"os"
)

// I know, util package is an antipattern. This code here is so small I don't care

func CheckErr(e error) {
	if e != nil {
		panic(e)
	}
}

func CheckArgs() {
	println("Verison:\t" + version.Str)
	if len(os.Args) < 3 {
		println("Usage: ./bin part input\nPart is 1 or 2\ninput is the path to the input file")
		os.Exit(1)
	}
}