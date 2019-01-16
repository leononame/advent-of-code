package day14

import (
	"bytes"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
	"gitlab.com/leononame/advent-of-code-2018/pkg/aoc"
)

var logger *logrus.Logger

func Run(c *aoc.Config) (result aoc.Result) {
	logger = c.Logger

	t0 := time.Now()
	count, _ := strconv.Atoi(c.Input[0])
	bs := []byte(c.Input[0])
	result.ParseTime = time.Since(t0)

	t1 := time.Now()
	recipes := []byte{'3', '7'}
	var e1, e2 = 0, 1
	var p2 = 0
	for i := 1; ; i++ {
		for len(recipes) < i*(count+10) {
			ne1, ne2 := recipes[e1]-'0', recipes[e2]-'0'
			s := strconv.Itoa(int(ne1 + ne2))
			recipes = append(recipes, s...)
			// New indices
			e1 = (e1 + int(ne1) + 1) % len(recipes)
			e2 = (e2 + int(ne2) + 1) % len(recipes)
		}
		if i == 1 {
			result.Duration1 = time.Since(t1)
		}
		p2 = bytes.Index(recipes, bs)
		if p2 != -1 {
			break
		}
	}
	result.Duration2 = time.Since(t1)
	result.Solution1 = string(recipes[count : count+10])
	result.Solution2 = p2
	return
}
