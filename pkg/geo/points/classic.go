package points

import "gitlab.com/leononame/advent-of-code-2018/pkg/geo"

// Classic is a classic implementation of Pointer using X and Y
type Classic struct {
	X, Y int
}

func NewClassicPointer(x, y int) geo.Pointer {
	return NewClassic(x, y)
}
func NewClassic(x, y int) Classic {
	return Classic{x, y}
}

func (c Classic) GetX() int {
	return c.X
}

func (c Classic) GetY() int {
	return c.Y
}

func (c Classic) Up() geo.Pointer {
	return &Classic{c.X, c.Y - 1}
}

func (c Classic) Down() geo.Pointer {
	return &Classic{c.X, c.Y + 1}
}

func (c Classic) Left() geo.Pointer {
	return &Classic{c.X - 1, c.Y}
}

func (c Classic) Right() geo.Pointer {
	return &Classic{c.X + 1, c.Y}
}
