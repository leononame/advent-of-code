package rect3

import (
	"gitlab.com/leononame/advent-of-code-2018/pkg/geo"
	"gitlab.com/leononame/advent-of-code-2018/pkg/geo/points3"
)

type Rectangle struct {
	Min, Max geo.Pointer3
}

func (r Rectangle) LongestSide() int {
	dx := r.Max.GetX() - r.Min.GetX()
	dy := r.Max.GetY() - r.Min.GetY()
	dz := r.Max.GetZ() - r.Min.GetZ()
	return max(dx, max(dy, dz))
}

func FromPointCloud(ps []geo.Pointer3) Rectangle {
	r := Rectangle{}
	var minX, minY, minZ, maxX, maxY, maxZ int
	for _, p := range ps {
		x, y, z := p.GetX(), p.GetY(), p.GetZ()
		minX, minY, minZ = min(minX, x), min(minY, y), min(minZ, z)
		maxX, maxY, maxZ = max(maxX, x), max(maxY, y), max(maxZ, z)
	}
	r.Min = points3.NewClassic(minX, minY, minZ)
	r.Max = points3.NewClassic(maxX, maxY, maxZ)
	return r
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
