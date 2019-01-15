package day05

import (
	"strings"
	"time"

	"github.com/logrusorgru/aurora"

	"github.com/sirupsen/logrus"

	"gitlab.com/leononame/advent-of-code-2018/pkg/aoc"
)

var logger *logrus.Logger

func Run(c *aoc.Config) (result aoc.Result) {
	logger = c.Logger
	result.ParseTime = 0

	t1 := time.Now()
	result.Solution1 = part1(c.Input[0])
	result.Duration1 = time.Since(t1)

	t2 := time.Now()
	result.Solution2 = part2(c.Input[0])
	result.Duration2 = time.Since(t2)
	return
}

func part1(polymer string) int {
	l := react(polymer)
	logger.Debugf("Reaction finished. Polymer of length %d is now of length %d.", aurora.Red(len(polymer)), aurora.Green(l))
	return l
}

func part2(polymer string) int {
	min := len(polymer)
	badChar := "a"
	for i := 0x41; i < 0x5B; i++ {
		cleaned := removeChar(polymer, i)
		length := react(cleaned)
		if length < min {
			min = length
			badChar = string(i)
		}
		logger.Debugf("Removing char %c results in a polymer of length %d.", aurora.Red(i), aurora.Green(length))
	}
	logger.Debugf("For part 2, the best solution is removing the character %s. This results in a length of: %d\n", badChar, min)
	return min
}

func removeChar(s string, c int) string {
	tmp := strings.Replace(s, string(c), "", -1)
	return strings.Replace(tmp, string(c+0x20), "", -1)
}

func react(polymer string) int {
	data := []byte(polymer)
	l := len(data)
	for {
		for i := 1; i < len(data); i++ {
			if data[i-1]-data[i] == 0x20 || data[i]-data[i-1] == 0x20 {
				data = append(data[:i-1], data[i+1:]...)
			}
		}
		if l == len(data) {
			break
		}
		l = len(data)
	}
	return l
}
