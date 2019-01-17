package day10

import (
	"fmt"
	"math"
	"strings"
	"time"

	"gitlab.com/leononame/advent-of-code-2018/pkg/geo"
	"gitlab.com/leononame/advent-of-code-2018/pkg/mmath"

	"github.com/sirupsen/logrus"
	"gitlab.com/leononame/advent-of-code-2018/pkg/aoc"
)

var logger *logrus.Logger

func Run(c *aoc.Config) (result aoc.Result) {
	logger = c.Logger

	t0 := time.Now()
	im := parse(c.Input)
	result.ParseTime = time.Since(t0)

	t1 := time.Now()
	im.tick()
	for {
		old := im.h
		im.tick()
		// Once the image is growing instead of shrinking, we surpassed our point
		if im.h > old {
			im.untick(1)
			break
		}
	}
	d := time.Since(t1)
	result.Duration2 = d
	result.Duration1 = d
	result.Solution2 = im.age
	result.Solution1 = "See output above"
	logger.Info("Result:\n", im)
	return
}

type position struct {
	geo.Point
	vx, vy int
}

type sky []*position

type image struct {
	sky
	maxX, minX, maxY, minY int
	h, w                   int
	age                    int
}

func (im *image) tick() {
	im.age++
	im.maxX, im.minX, im.maxY, im.minY = math.MinInt64, math.MaxInt64, math.MinInt64, math.MaxInt64
	for _, p := range im.sky {
		p.X += p.vx
		p.Y += p.vy
		im.maxX = mmath.Max(p.X, im.maxX)
		im.minX = mmath.Min(p.X, im.minX)
		im.maxY = mmath.Max(p.Y, im.maxY)
		im.minY = mmath.Min(p.Y, im.minY)
	}
	im.h = im.maxY - im.minY
	im.w = im.maxX - im.minX
}

func (im *image) untick(count int) {
	im.age -= count
	im.maxX, im.minX, im.maxY, im.minY = math.MinInt64, math.MaxInt64, math.MinInt64, math.MaxInt64
	for _, p := range im.sky {
		p.X -= (count) * (p.vx)
		p.Y -= (count) * (p.vy)
		im.maxX = mmath.Max(p.X, im.maxX)
		im.minX = mmath.Min(p.X, im.minX)
		im.maxY = mmath.Max(p.Y, im.maxY)
		im.minY = mmath.Min(p.Y, im.minY)
	}
	im.h = im.maxX - im.minX
	im.w = im.maxY - im.minY
}

func (im *image) String() string {
	data := make([][]string, im.w+1)
	for i := range data {
		data[i] = make([]string, im.h+1)
	}
	for i := range data {
		for j := range data[i] {
			data[i][j] = " "
		}
	}
	for _, p := range im.sky {
		data[p.Y-im.minY][p.X-im.minX] = "#"
	}

	var sb strings.Builder
	for _, row := range data {
		sb.WriteString(strings.Join(row, ""))
		sb.WriteByte('\n')
	}
	return sb.String()
}

func parse(input []string) *image {
	var im image
	for _, l := range input {
		var p position
		fmt.Sscanf(l, "position=<%d, %d> velocity=<%d, %d>", &p.X, &p.Y, &p.vx, &p.vy)
		im.sky = append(im.sky, &p)
	}
	return &im
}
