package geo

import "gitlab.com/leononame/advent-of-code-2018/pkg/mmath"

type Point struct {
	X, Y int
}

func (p Point) GetX() int {
	return p.X
}

func (p Point) GetY() int {
	return p.Y
}

func (p Point) Up() Point {
	return Point{p.X, p.Y - 1}
}

func (p Point) Down() Point {
	return Point{p.X, p.Y + 1}
}

func (p Point) Left() Point {
	return Point{p.X - 1, p.Y}
}

func (p Point) Right() Point {
	return Point{p.X + 1, p.Y}
}

func (p Point) Manhattan(to Point) int {
	return mmath.Abs(p.X-to.X) + mmath.Abs(p.Y-to.Y)
}

func (p Point) Neighbours() []Point {
	return []Point{p.Down(), p.Left(), p.Up(), p.Right()}
}
