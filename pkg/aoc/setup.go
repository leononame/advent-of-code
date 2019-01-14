package aoc

import (
	"bufio"
	"os"
	"time"

	"github.com/pborman/getopt/v2"
	"github.com/sirupsen/logrus"
)

type Result struct {
	ParseTime time.Duration
	Duration1 time.Duration
	Duration2 time.Duration
	Solution1 string
	Solution2 string
}

type Config struct {
	Input      []string
	Logger     *logrus.Logger
	SubCommand string
	Year, Day  int
}

func Setup() *Config {
	verbose := getopt.BoolLong("verbose", 'v', "Log more information")
	help := getopt.BoolLong("help", 'h', "Show help")
	fname := getopt.StringLong("input", 'i', "input", "Path to the input file, defaults to 'input'", "path")
	year := getopt.IntLong("year", 'y', 2018, "Select the year, defaults to the latest challenge", "2018")
	day := getopt.IntLong("day", 'd', 25, "Select the day, defaults to the latest challenge", "25")
	// part := getopt.IntLong("part", 'p', 0, "Select only one part of the challenge", "2")

	getopt.Parse()

	if *help {
		getopt.PrintUsage(os.Stdout)
		os.Exit(0)
	}

	// if *part != 0 && *part != 1 && *part != 2 {
	// 	logrus.Errorf("Part %d does not exist. Only 1 and 2 are allowed.", *part)
	// 	os.Exit(1)
	// }

	o := Config{}
	o.Input = readFile(*fname)
	o.Logger = logrus.New()
	o.Logger.SetLevel(logrus.InfoLevel)
	if *verbose {
		o.Logger.SetLevel(logrus.DebugLevel)
	}
	o.Year, o.Day = *year, *day

	if getopt.NArgs() > 0 {
		o.SubCommand = getopt.Args()[0]
	}
	return &o
}

func readFile(path string) []string {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	s := bufio.NewScanner(f)
	var lines []string
	for s.Scan() {
		lines = append(lines, s.Text())
	}
	return lines
}
