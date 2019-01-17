package geo

import "gitlab.com/leononame/advent-of-code/pkg/mmath"

// Point3 is a classic implementation of Pointer3 using X, Y and Z
type Point3 struct {
	X, Y, Z int
}

func (c Point3) GetX() int {
	return c.X
}

func (c Point3) GetY() int {
	return c.Y
}

func (c Point3) GetZ() int {
	return c.Z
}

func (c Point3) Left() Point3 {
	return Point3{c.X - 1, c.Y, c.Z}
}

func (c Point3) Right() Point3 {
	return Point3{c.X + 1, c.Y, c.Z}
}

func (c Point3) Up() Point3 {
	return Point3{c.X, c.Y - 1, c.Z}
}

func (c Point3) Down() Point3 {
	return Point3{c.X, c.Y + 1, c.Z}
}

func (c Point3) Higher() Point3 {
	return Point3{c.X, c.Y, c.Z - 1}
}

func (c Point3) Lower() Point3 {
	return Point3{c.X, c.Y, c.Z + 1}
}

func (p Point3) Manhattan(to Point3) int {
	return mmath.Abs(p.X-to.X) + mmath.Abs(p.Y-to.Y) + mmath.Abs(p.Z-to.Z)
}

func (p Point3) Neighbours() []Point3 {
	return []Point3{p.Down(), p.Left(), p.Up(), p.Right(), p.Lower(), p.Higher()}
}
