package geo

// Pointer is an interface for all point implementations
type Pointer interface {
	GetX() int
	GetY() int
	Up() Pointer
	Down() Pointer
	Left() Pointer
	Right() Pointer
}
