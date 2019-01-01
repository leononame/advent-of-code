package points

import "gitlab.com/leononame/advent-of-code-2018/pkg/geo"

// Complex implements points based on complex numbers
type Complex complex64

func NewComplexPointer(x, y int) geo.Pointer {
	return NewComplex(x, y)
}
func NewComplex(x, y int) Complex {
	c := Complex(complex(float32(x), float32(y)))
	return c
}

func (c Complex) GetX() int {
	return int(real(complex64(c)))
}

func (c Complex) GetY() int {
	return int(real(complex64(c)))
}

func (c Complex) Up() geo.Pointer {
	c2 := NewComplex(c.GetX(), c.GetY()-1)
	return c2
}

func (c Complex) Down() geo.Pointer {
	c2 := NewComplex(c.GetX(), c.GetY()+1)
	return c2
}

func (c Complex) Left() geo.Pointer {
	c2 := NewComplex(c.GetX()-1, c.GetY())
	return c2
}

func (c Complex) Right() geo.Pointer {
	c2 := NewComplex(c.GetX()+1, c.GetY())
	return c2
}
