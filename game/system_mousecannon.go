package game

// import (
// 	"github.com/go-gl/mathgl/mgl64"
// 	"github.com/veandco/go-sdl2/sdl"
// 	"go.owls.io/fauxbox/engine"
// )

// ////////////////////////////////////////////////////////////////////////////////
// // SYSTEM //
// ///////////

// type MouseCannon struct {
// 	armed            bool
// 	NextBulletPos    mgl64.Vec2
// 	NextBulletAimPos mgl64.Vec2
// }

// func init() {
// 	engine.Register(&MouseCannon)
// }

// ////////////////////////////////////////////////////////////////////////////////
// // STATES //
// ///////////

// var MouseCannonStates map[engine.State]bool = map[engine.State]bool{INGAME: true}

// func (mc *MouseCannon) States() map[engine.State]bool {
// 	return MouseCannonStates
// }

// ////////////////////////////////////////////////////////////////////////////////
// // CALLBACKS //
// //////////////

// func (mc *MouseCannon) StateEntry() {

// }

// func (mc *MouseCannon) ProcessEvent(event sdl.Event) {
// 	// Launcher
// 	switch e := event.(type) {
// 	case *sdl.MouseButtonEvent:
// 		switch e.Type {
// 		case sdl.MOUSEBUTTONDOWN:
// 			if e.Button == sdl.BUTTON_LEFT {
// 				mc.arm(mgl64.Vec2{World.CamXInt32(e.X), World.CamYInt32(e.Y)})
// 			} else if e.Button == sdl.BUTTON_RIGHT {
// 				mc.unarm()
// 			}

// 		case sdl.MOUSEBUTTONUP:
// 			if e.Button == sdl.BUTTON_LEFT {
// 				mc.fire(mgl64.Vec2{World.CamXInt32(e.X), World.CamYInt32(e.Y)})
// 			}
// 		}

// 	case *sdl.MouseMotionEvent:
// 		if mc.isArmed() {
// 			mc.aim(mgl64.Vec2{World.CamXInt32(e.X), World.CamYInt32(e.Y)})
// 		}
// 	}
// }

// func (mc *MouseCannon) Update(dt float64) {
// 	// Launcher
// 	if mc.armed {
// 		x, y, _ := sdl.GetMouseState()
// 		mc.aim(mgl64.Vec2{World.CamXInt(x), World.CamYInt(y)})
// 	}

// 	// Bullets
// 	for _, bullet := range States.CurrentState.(State).GetBullets() {
// 		if !bullet.active {
// 			continue
// 		}

// 		bullet.livedtime += dt
// 		if bullet.livedtime > bullet.lifetime {
// 			bullet.active = false
// 			continue
// 		}

// 		bullet.pos = bullet.pos.Add(bullet.vel.Mul(dt))
// 	}
// }

// func (mc *MouseCannon) Draw(dt float64) {
// 	// Launcher
// 	if mc.armed {
// 		engine.Renderer.SetDrawColor(255, 255, 255, 255)
// 		engine.Renderer.DrawLine(
// 			int(World.X(mc.NextBulletPos.X())), int(World.Y(mc.NextBulletPos.Y())),
// 			int(World.X(mc.NextBulletAimPos.X())), int(World.Y(mc.NextBulletAimPos.Y())))
// 	}

// 	// Bullets
// 	for _, bullet := range States.CurrentState.(State).GetBullets() {
// 		if !bullet.active {
// 			continue
// 		}

// 		engine.Renderer.SetDrawColor(255, 0, 255, 255)
// 		engine.Renderer.DrawRect(&sdl.Rect{
// 			World.XInt32(bullet.pos.X()) - 2,
// 			World.YInt32(bullet.pos.Y()) - 2,
// 			4, 4})
// 	}
// }

// ////////////////////////////////////////////////////////////////////////////////
// // FUNCTIONS //
// //////////////

// func (mc *MouseCannon) isArmed() bool {
// 	return mc.armed
// }

// func (mc *MouseCannon) arm(mousePos mgl64.Vec2) {
// 	mc.NextBulletPos = mousePos
// 	mc.NextBulletAimPos = mousePos
// 	mc.armed = true
// }

// func (mc *MouseCannon) unarm() {
// 	mc.armed = false
// }

// func (mc *MouseCannon) aim(mousePos mgl64.Vec2) {
// 	if mc.armed {
// 		mc.NextBulletAimPos = mousePos
// 	}
// }

// func (mc *MouseCannon) fire(mousePos mgl64.Vec2) {
// 	if mc.armed {
// 		for _, bullet := range States.CurrentState.(State).GetBullets() {
// 			if bullet.active {
// 				continue
// 			}

// 			bullet.Fire(mc.NextBulletPos, mc.NextBulletAimPos.Sub(mc.NextBulletPos), 2)
// 			break
// 		}
// 		mc.armed = false
// 	}
// }
