package engine

type RadiusCollidable interface {
	// Object
	Radius() (x, y, radius int)
}

type BoundsCollidable interface {
	// Object
	Bounds() (x, y, width, height int)
}
