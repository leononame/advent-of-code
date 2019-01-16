package day08

import (
	"strconv"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"gitlab.com/leononame/advent-of-code-2018/pkg/aoc"
)

var logger *logrus.Logger

func Run(c *aoc.Config) (result aoc.Result) {
	logger = c.Logger

	t0 := time.Now()
	t := parse(c.Input[0])
	result.ParseTime = time.Since(t0)

	t1 := time.Now()
	result.Solution1 = part1(t)
	result.Duration1 = time.Since(t1)

	t2 := time.Now()
	result.Solution2 = part2(t)
	result.Duration2 = time.Since(t2)
	return
}

type node struct {
	meta     []int
	children []*node
}

type tree struct {
	root *node
}

func part1(t tree) int {
	n := calcMeta(t.root)
	logger.Debug("Tree has a meta sum of", n)
	return n
}

func part2(t tree) int {
	v := calcValue(t.root)
	logger.Debug("The root node has the value", v)
	return v
}

func calcValue(n *node) int {
	if n.children == nil {
		return calcMeta(n)
	}
	sum := 0
	for _, m := range n.meta {
		if m == 0 || m > len(n.children) {
			continue
		}
		sum += calcValue(n.children[m-1])
	}
	return sum
}

func calcMeta(n *node) int {
	s := sum(n.meta)
	for _, node := range n.children {
		s += calcMeta(node)
	}
	return s
}

func parse(input string) tree {
	l := strings.Split(input, " ")
	var data []int
	for _, entry := range l {
		num, _ := strconv.Atoi(entry)
		data = append(data, num)
	}

	t := tree{}
	n, _ := parseNode(0, &data)
	t.root = n
	return t
}

func parseNode(start int, data *[]int) (n *node, end int) {
	n = &node{}

	children := (*data)[start]
	metas := (*data)[start+1]
	idx := start + 2

	for i := 0; i < children; i++ {
		nn, ii := parseNode(idx, data)
		idx = ii
		n.children = append(n.children, nn)
	}

	for i := 0; i < metas; i++ {
		n.meta = append(n.meta, (*data)[idx])
		idx++
	}

	return n, idx
}

func sum(nums []int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return sum
}
