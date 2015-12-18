package game

import (
	"github.com/go-gl/mathgl/mgl64"
	"github.com/veandco/go-sdl2/sdl"
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

func (mc *MouseCannon) ProcessEvent(event sdl.Event) {
	switch e := event.(type) {
	case *sdl.MouseButtonEvent:
		switch e.Type {
		case sdl.MOUSEBUTTONDOWN:
			if e.Button == sdl.BUTTON_LEFT {
				mc.arm(mgl64.Vec2{World.CamXInt32(e.X), World.CamYInt32(e.Y)})
			} else if e.Button == sdl.BUTTON_RIGHT {
				mc.unarm()
			}

		case sdl.MOUSEBUTTONUP:
			if e.Button == sdl.BUTTON_LEFT {
				mc.fire(mc.Bullets, mgl64.Vec2{World.CamXInt32(e.X), World.CamYInt32(e.Y)})
			}
		}

	case *sdl.MouseMotionEvent:
		if mc.isArmed() {
			mc.aim(mgl64.Vec2{World.CamXInt32(e.X), World.CamYInt32(e.Y)})
		}
	}
}

func (mc *MouseCannon) Update(dt float64) {
	if mc.armed {
		x, y, _ := sdl.GetMouseState()
		mc.aim(mgl64.Vec2{World.CamXInt(x), World.CamYInt(y)})
	}
}

func (mc *MouseCannon) Draw(dt float64) {
	if mc.armed {
		engine.Renderer.SetDrawColor(255, 255, 255, 255)
		engine.Renderer.DrawLine(
			int(World.X(mc.NextBulletPos.X())), int(World.Y(mc.NextBulletPos.Y())),
			int(World.X(mc.NextBulletAimPos.X())), int(World.Y(mc.NextBulletAimPos.Y())))
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
