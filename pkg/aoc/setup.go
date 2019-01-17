package aoc

import (
	"os"
	"time"

	"github.com/pborman/getopt/v2"
	"github.com/sirupsen/logrus"
)

type Result struct {
	ParseTime time.Duration
	Duration1 time.Duration
	Duration2 time.Duration
	Solution1 interface{}
	Solution2 interface{}
}

func Setup() *Config {
	verbose := getopt.BoolLong("verbose", 'v', "Log more information")
	help := getopt.BoolLong("help", 'h', "Show help")
	fname := getopt.StringLong("input", 'i', "", "Path to the input file, defaults to 'input/%year/day%day'", "input/2018/day03")
	year := getopt.IntLong("year", 'y', CurrentYear, "Select the year, defaults to the latest challenge", "2018")
	day := getopt.IntLong("day", 'd', CurrentChallenge, "Select the day, defaults to the latest challenge", "25")
	all := getopt.BoolLong("all", 'a', "Run all challenges of a given year", "true")
	getopt.Parse()

	logger := logrus.New()
	logger.SetLevel(logrus.InfoLevel)
	if *verbose {
		logger.SetLevel(logrus.DebugLevel)
	}
	if *all {
		logger.SetLevel(logrus.ErrorLevel)
	}

	if *help {
		getopt.PrintUsage(os.Stdout)
		os.Exit(0)
	}
	if *all && getopt.IsSet("day") {
		logger.Error("Option -a/--all can't be combined with option -d/--day.")
		os.Exit(1)
	}

	o := Config{Logger: logger, File: *fname, Year: *year, Day: *day, All: *all}
	o.ReadFile()
	if getopt.NArgs() > 0 {
		o.SubCommand = getopt.Args()[0]
	}
	return &o
}
