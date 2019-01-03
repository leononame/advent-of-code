package geo3

// Pointer is an interface for all 3d point implementations
type Pointer interface {
	GetX() int
	GetY() int
	GetZ() int
	Up() Pointer
	Down() Pointer
	Left() Pointer
	Right() Pointer
	Higher() Pointer
	Lower() Pointer
}
