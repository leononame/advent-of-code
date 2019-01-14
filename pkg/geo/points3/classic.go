package points3

import "gitlab.com/leononame/advent-of-code-2018/pkg/geo"

// Classic is a classic implementation of Pointer3 using X and Y
type Classic struct {
	X, Y, Z int
}

func NewClassicPointer(x, y, z int) geo.Pointer3 {
	return NewClassic(x, y, z)
}
func NewClassic(x, y, z int) Classic {
	return Classic{x, y, z}
}

func (c Classic) GetX() int {
	return c.X
}

func (c Classic) GetY() int {
	return c.Y
}

func (c Classic) GetZ() int {
	return c.Z
}

func (c Classic) Left() geo.Pointer3 {
	return Classic{c.X - 1, c.Y, c.Z}
}

func (c Classic) Right() geo.Pointer3 {
	return Classic{c.X + 1, c.Y, c.Z}
}

func (c Classic) Up() geo.Pointer3 {
	return Classic{c.X, c.Y - 1, c.Z}
}

func (c Classic) Down() geo.Pointer3 {
	return Classic{c.X, c.Y + 1, c.Z}
}

func (c Classic) Higher() geo.Pointer3 {
	return Classic{c.X, c.Y, c.Z - 1}
}

func (c Classic) Lower() geo.Pointer3 {
	return Classic{c.X, c.Y, c.Z + 1}
}
