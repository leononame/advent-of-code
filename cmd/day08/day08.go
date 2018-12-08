package main

import (
	"fmt"
	"strconv"
	"strings"

	"gitlab.com/leononame/advent-of-code-2018/pkg/util"
)

type node struct {
	meta     []int
	children []*node
}

type tree struct {
	root *node
}

func main() {
	fmt.Println("Challenge:\t2018-08")
	input := util.GetInput("input/day08")[0]
	l := strings.Split(input, " ")
	var data []int
	for _, entry := range l {
		num, _ := strconv.Atoi(entry)
		data = append(data, num)
	}

	t := parse(data)
	part1(t)
	part2(t)
}

func part1(t tree) {
	n := calcMeta(t.root)
	fmt.Println("Tree has a meta sum of", n)
}

func part2(t tree) {
	v := calcValue(t.root)
	fmt.Println("The root node has the value", v)
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

func parse(data []int) tree {
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
