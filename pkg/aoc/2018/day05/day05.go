package day05

import (
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/logrusorgru/aurora"

	"github.com/sirupsen/logrus"

	"gitlab.com/leononame/advent-of-code/pkg/aoc"
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
	l, _ := reactParallel([]byte(polymer))
	logger.Debugf("Reaction finished. Polymer of length %d is now of length %d.", aurora.Red(len(polymer)), aurora.Green(l))
	return l
}

func part2(polymer string) int {
	min := len(polymer)
	badChar := "a"
	for i := 0x41; i < 0x5B; i++ {
		cleaned := removeChar(polymer, i)
		length, _ := reactParallel([]byte(cleaned))
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

func reactParallel(data []byte) (int, []byte) {
	// Turn into NumCPU (or NumCPU+1) batches for parallel processing
	batchSize := len(data) / runtime.NumCPU()
	var batches [][]byte
	for batchSize < len(data) {
		data, batches = data[batchSize:], append(batches, data[0:batchSize:batchSize])
	}
	batches = append(batches, data)
	// waitgroup
	var wg sync.WaitGroup
	wg.Add(len(batches))
	// React each batch in parallel
	for i := range batches {
		go func(i int) {
			defer wg.Done()
			for {
				l := len(batches[i])
				_, batches[i] = react(batches[i])
				if len(batches[i]) == l {
					return
				}
			}
		}(i)
	}
	wg.Wait()
	// Merge data
	data = data[:0]
	for _, batch := range batches {
		data = append(data, batch...)
	}
	// React the merged polymer
	_, data = react(data)
	return len(data), data
}

func react(data []byte) (int, []byte) {
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
	return l, data
}
