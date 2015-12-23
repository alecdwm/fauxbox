package game

import (
	"github.com/go-gl/mathgl/mgl64"
	"go.owls.io/fauxbox/engine"
)

////////////////////////////////////////////////////////////////////////////////
// SYSTEM //
///////////

type Camera struct {
	Pos       mgl64.Vec2
	TargetPos mgl64.Vec2
}

func init() {
	engine.Register(&GameWorld.Camera)
}

////////////////////////////////////////////////////////////////////////////////
// STATES //
///////////

var CameraStates map[engine.State]bool = map[engine.State]bool{INGAME: true}

func (c *Camera) States() map[engine.State]bool {
	return CameraStates
}

////////////////////////////////////////////////////////////////////////////////
// CALLBACKS //
//////////////

func (c *Camera) Update(dt float64) {
	c.Pos = c.Pos.Add(c.TargetPos.Sub(c.Pos).Mul(dt))
}

////////////////////////////////////////////////////////////////////////////////
// FUNCTIONS //
//////////////

func (c *Camera) SetTarget(targetPos mgl64.Vec2) {
	c.TargetPos = targetPos
}
