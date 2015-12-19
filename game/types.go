package game

import (
	"github.com/go-gl/mathgl/mgl64"
	"go.owls.io/fauxbox/engine"
)

const (
	MAINMENU engine.State = iota
	INGAME
)

type Direction uint

const (
	UP Direction = iota
	LUP
	RUP
	LEFT
	RIGHT
	LDOWN
	RDOWN
	DOWN
)

type Object interface {
	ID() uint
}

type Transformed interface {
	Object

	Pos() mgl64.Vec2
	Rot() float64
}

type Controllable interface {
	Transformed

	Move(mgl64.Vec2)
}

type RadiusCollidable interface {
	Transformed

	Bounds() (radius int)
}

type BoundsCollidable interface {
	Transformed

	Bounds() (width, height int)
}
