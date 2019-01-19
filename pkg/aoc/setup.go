package aoc

import (
	"fmt"
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

const usage = `Usage: aoc [-ahv] [-d 25] [-y 2018] [-i input/2018/day03] [subcommand]
 -a, --all        Run all challenges of a given year. Incompatible with options -d and -v.
 -d, --day=25     Select the day, defaults to the latest challenge
 -y, --year=2018  Select the year, defaults to the latest challenge
 -i, --input=input/2018/day03
                  Specify path to the input file, defaults to 'input/%year/day%day'
                  e.g. input/2018/day25
 -h, --help       Show help
 -v, --verbose    Log more information`

func printUsage() {
	fmt.Println(usage)
}

func Setup() *Config {
	verbose := getopt.BoolLong("verbose", 'v', "Log more information")
	help := getopt.BoolLong("help", 'h', "Show help")
	fname := getopt.StringLong("input", 'i', "", "Path to the input file, defaults to 'input/%year/day%day'", "input/2018/day03")
	year := getopt.IntLong("year", 'y', CurrentYear, "Select the year, defaults to the latest challenge", "2018")
	day := getopt.IntLong("day", 'd', CurrentChallenge, "Select the day, defaults to the latest challenge", "25")
	all := getopt.BoolLong("all", 'a', "Run all challenges of a given year", "true")
	getopt.SetUsage(printUsage)
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
		printUsage()
		os.Exit(0)
	}
	if *all && getopt.IsSet("day") {
		logger.Error("Option -a/--all can't be combined with option -d/--day.")
		os.Exit(1)
	}
	if *all && *verbose {
		logger.Error("Option -a/--all can't be combined with option -v/--verbose.")
		os.Exit(1)
	}

	o := Config{Logger: logger, File: *fname, Year: *year, Day: *day, All: *all}
	o.ReadFile()
	if getopt.NArgs() > 0 {
		o.SubCommand = getopt.Args()[0]
	}
	return &o
}
