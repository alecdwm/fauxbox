package game

import (
	"github.com/go-gl/mathgl/mgl64"
	"github.com/veandco/go-sdl2/sdl"
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
	IsPlayer() bool
	IsNetworked() bool

	SetPosition(newPos mgl64.Vec2)
	SetVelocity(newVel mgl64.Vec2)
	SetTarget(newTarget mgl64.Vec2)
	Speed() float64
	SetSpeed(speed float64)
	SetColor(color sdl.Color)

	// Move(mgl64.Vec2)
}

type RadiusCollidable interface {
	Transformed

	Bounds() (radius int)
}

type BoundsCollidable interface {
	Transformed

	Bounds() (width, height int)
}
