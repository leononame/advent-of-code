package points3

import "gitlab.com/leononame/advent-of-code-2018/pkg/geo3"

func Equal(p1, p2 geo3.Pointer) bool {
	x, y, z := p1.GetX(), p1.GetY(), p1.GetZ()
	dx, dy, dz := p2.GetX(), p2.GetY(), p2.GetZ()
	return x == dx && y == dy && z == dz
}
