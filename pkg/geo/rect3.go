package geo

import (
	"gitlab.com/leononame/advent-of-code/pkg/mmath"
)

type Rectangle3 struct {
	Min, Max Point3
}

func (r Rectangle3) LongestSide() int {
	dx := r.Max.GetX() - r.Min.GetX()
	dy := r.Max.GetY() - r.Min.GetY()
	dz := r.Max.GetZ() - r.Min.GetZ()
	return mmath.Max(dx, mmath.Max(dy, dz))
}

func FromPointCloud(ps []Point3) Rectangle3 {
	r := Rectangle3{}
	var minX, minY, minZ, maxX, maxY, maxZ int
	for _, p := range ps {
		minX, minY, minZ = mmath.Min(minX, p.X), mmath.Min(minY, p.Y), mmath.Min(minZ, p.Z)
		maxX, maxY, maxZ = mmath.Max(maxX, p.X), mmath.Max(maxY, p.Y), mmath.Max(maxZ, p.Z)
	}
	r.Min = Point3{minX, minY, minZ}
	r.Max = Point3{maxX, maxY, maxZ}
	return r
}
