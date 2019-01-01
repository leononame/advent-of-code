package points

import "gitlab.com/leononame/advent-of-code-2018/pkg/geo"

func Equal(p1, p2 geo.Pointer) bool {
	x, y := p1.GetX(), p1.GetY()
	dx, dy := p2.GetX(), p2.GetY()
	return x == dx && y == dy
}
