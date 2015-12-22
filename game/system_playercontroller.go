package game

import (
	"github.com/go-gl/mathgl/mgl64"
	"github.com/veandco/go-sdl2/sdl"
	"go.owls.io/fauxbox/engine"
)

////////////////////////////////////////////////////////////////////////////////
// SYSTEM //
///////////

type PlayerController struct {
	target Controllable
	vel    mgl64.Vec2
}

func init() {
	engine.Register(&PlayerController{})
}

////////////////////////////////////////////////////////////////////////////////
// CALLBACKS //
//////////////

func (pc *PlayerController) ProcessEvent(event sdl.Event) {
	if pc.target == nil {
		for i := range GameWorld.Objects {
			if player, ok := GameWorld.Objects[i].(Controllable); ok {
				if !player.IsPlayer() {
					continue
				}
				if player.IsNetworked() {
					continue
				}
				pc.target = player
			}
		}
		return
	}

	switch e := event.(type) {
	case *sdl.KeyDownEvent, *sdl.KeyUpEvent:
		pc.vel = mgl64.Vec2{0, 0}

		if engine.Keyboard[sdl.SCANCODE_W] == 1 {
			pc.vel = pc.vel.Add(mgl64.Vec2{0, -10})
		}
		if engine.Keyboard[sdl.SCANCODE_A] == 1 {
			pc.vel = pc.vel.Add(mgl64.Vec2{-10, 0})
		}
		if engine.Keyboard[sdl.SCANCODE_S] == 1 {
			pc.vel = pc.vel.Add(mgl64.Vec2{0, 10})
		}
		if engine.Keyboard[sdl.SCANCODE_D] == 1 {
			pc.vel = pc.vel.Add(mgl64.Vec2{10, 0})
		}

		if pc.vel.X() != 0 || pc.vel.Y() != 0 {
			pc.vel = pc.vel.Normalize()
			pc.target.SetVelocity(pc.vel.Mul(pc.target.Speed() / pc.vel.Len()))
		} else {
			pc.target.SetVelocity(mgl64.Vec2{0, 0})
		}

	case *sdl.MouseMotionEvent:
		pc.target.SetTarget(
			GameWorld.Ray(mgl64.Vec2{float64(e.X), float64(e.Y)}))
	}
}

func (pc *PlayerController) Update(dt float64) {
	pc.target.SetTarget(
		GameWorld.Ray(
			mgl64.Vec2{float64(engine.Mouse.X), float64(engine.Mouse.Y)}))
}

////////////////////////////////////////////////////////////////////////////////
// FUNCTIONS //
//////////////

func (pc *PlayerController) SetPlayerTarget() {

}
