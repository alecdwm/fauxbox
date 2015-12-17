package game

import (
	"github.com/dradtke/go-allegro/allegro"
	"github.com/dradtke/go-allegro/allegro/primitives"
	"github.com/go-gl/mathgl/mgl64"
	"go.owls.io/fauxbox/engine"
)

////////////////////////////////////////////////////////////////////////////////
// SYSTEM //////////////////////////////////////////////////////////////////////
////////////

type MouseCannon struct {
	Bullets *MouseCannonBullets

	armed            bool
	NextBulletPos    mgl64.Vec2
	NextBulletAimPos mgl64.Vec2
}

func init() {
	mouseCannon := &MouseCannon{
		Bullets: &MouseCannonBullets{},
	}
	engine.Register(mouseCannon)
	engine.Register(mouseCannon.Bullets)
}

////////////////////////////////////////////////////////////////////////////////
// CALLBACKS ///////////////////////////////////////////////////////////////////
///////////////

func (mc *MouseCannon) ProcessEvent(event interface{}) {
	switch e := event.(type) {
	case allegro.MouseButtonDownEvent:
		if e.Button() == 1 {
			mc.arm(mgl64.Vec2{float64(e.X()), float64(e.Y())})
		} else if e.Button() == 2 {
			mc.unarm()
		}

	case allegro.MouseAxesEvent:
		if mc.isArmed() {
			mc.aim(mgl64.Vec2{float64(e.X()), float64(e.Y())})
		}

	case allegro.MouseButtonUpEvent:
		if e.Button() == 1 {
			mc.fire(mc.Bullets, mgl64.Vec2{float64(e.X()), float64(e.Y())})
		}
	}
}

func (mc *MouseCannon) Draw(dt float64) {
	if mc.armed {
		primitives.DrawLine(
			primitives.Point{float32(mc.NextBulletPos.X()), float32(mc.NextBulletPos.Y())},
			primitives.Point{float32(mc.NextBulletAimPos.X()), float32(mc.NextBulletAimPos.Y())},
			allegro.MapRGB(255, 255, 255),
			1.0,
		)
	}
}

////////////////////////////////////////////////////////////////////////////////
// FUNCTIONS ///////////////////////////////////////////////////////////////////
///////////////

func (mc *MouseCannon) isArmed() bool {
	return mc.armed
}

func (mc *MouseCannon) arm(mousePos mgl64.Vec2) {
	mc.NextBulletPos = mousePos
	mc.NextBulletAimPos = mousePos
	mc.armed = true
}

func (mc *MouseCannon) unarm() {
	mc.armed = false
}

func (mc *MouseCannon) aim(mousePos mgl64.Vec2) {
	if mc.armed {
		mc.NextBulletAimPos = mousePos
	}
}

func (mc *MouseCannon) fire(bulletPool *MouseCannonBullets, mousePos mgl64.Vec2) {
	if mc.armed {
		bulletPool.fire(mc.NextBulletPos, mc.NextBulletAimPos.Sub(mc.NextBulletPos), 2)
		mc.armed = false
	}
}
