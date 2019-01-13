package points

import "gitlab.com/leononame/advent-of-code-2018/pkg/geo"

func Manhattan(p1, p2 geo.Pointer) int {
	return abs(p1.GetX()-p2.GetX()) + abs(p1.GetY()-p2.GetY())
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
