package geo

// Pointer3 is an interface for all 3d point implementations
type Pointer3 interface {
	GetX() int
	GetY() int
	GetZ() int
	Up() Pointer3
	Down() Pointer3
	Left() Pointer3
	Right() Pointer3
	Higher() Pointer3
	Lower() Pointer3
}
