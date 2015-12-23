package game

import (
	"github.com/go-gl/mathgl/mgl64"
	"github.com/veandco/go-sdl2/sdl"
	"go.owls.io/fauxbox/engine"
)

////////////////////////////////////////////////////////////////////////////////
// SYSTEM //
///////////

type MouseCannon struct {
	bullets          []*Bullet
	armed            bool
	NextBulletPos    mgl64.Vec2
	NextBulletAimPos mgl64.Vec2
}

func init() {
	engine.Register(&MouseCannon{})
}

////////////////////////////////////////////////////////////////////////////////
// STATES //
///////////

var MouseCannonStates map[engine.State]bool = map[engine.State]bool{INGAME: true}

func (mc *MouseCannon) States() map[engine.State]bool {
	return MouseCannonStates
}

////////////////////////////////////////////////////////////////////////////////
// CALLBACKS //
//////////////

func (mc *MouseCannon) StateEntry() {
	for i := 0; i < 20; i++ {
		mc.bullets = append(mc.bullets,
			&Bullet{
				color:   sdl.Color{255, 0, 0, 255},
				posRect: &sdl.Rect{0, 0, 2, 2},
			})
		engine.Register(mc.bullets[i])
	}
}

func (mc *MouseCannon) ProcessEvent(event sdl.Event) {
	switch e := event.(type) {
	case *sdl.MouseButtonEvent:
		switch e.Type {
		case sdl.MOUSEBUTTONDOWN:
			if e.Button == sdl.BUTTON_LEFT {
				mc.arm(GameWorld.Ray(mgl64.Vec2{float64(e.X), float64(e.Y)}))
			} else if e.Button == sdl.BUTTON_RIGHT {
				mc.unarm()
			}

		case sdl.MOUSEBUTTONUP:
			if e.Button == sdl.BUTTON_LEFT {
				mc.fire(GameWorld.Ray(mgl64.Vec2{float64(e.X), float64(e.Y)}))
			}
		}

	case *sdl.MouseMotionEvent:
		if mc.isArmed() {
			mc.aim(GameWorld.Ray(mgl64.Vec2{float64(e.X), float64(e.Y)}))
		}
	}
}

func (mc *MouseCannon) Update(dt float64) {
	if mc.armed {
		mc.aim(GameWorld.Ray(mgl64.Vec2{float64(engine.Mouse.X), float64(engine.Mouse.Y)}))
	}
}

func (mc *MouseCannon) Draw(dt float64) {
	if mc.armed {
		engine.Renderer.SetDrawColor(255, 255, 255, 255)
		engine.Renderer.DrawLine(
			int(GameWorld.Pos(mc.NextBulletPos).X()), int(GameWorld.Pos(mc.NextBulletPos).Y()),
			int(GameWorld.Pos(mc.NextBulletAimPos).X()), int(GameWorld.Pos(mc.NextBulletAimPos).Y()))
	}
}

////////////////////////////////////////////////////////////////////////////////
// FUNCTIONS //
//////////////

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

func (mc *MouseCannon) fire(mousePos mgl64.Vec2) {
	if mc.armed {
		for _, bullet := range mc.bullets {
			if bullet.active {
				continue
			}

			bullet.Fire(mc.NextBulletPos,
				mc.NextBulletAimPos.Sub(mc.NextBulletPos),
				sdl.Color{255, 0, 0, 255},
				2)
			break
		}
		mc.armed = false
	}
}
