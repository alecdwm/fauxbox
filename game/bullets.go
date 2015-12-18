package game

import (
	"github.com/Sirupsen/logrus"
	"github.com/dradtke/go-allegro/allegro"
	"github.com/dradtke/go-allegro/allegro/primitives"
	"github.com/go-gl/mathgl/mgl64"
)

////////////////////////////////////////////////////////////////////////////////
// SYSTEM //////////////////////////////////////////////////////////////////////
////////////

type MouseCannonBullets struct {
	bullets [40]Bullet
}

func init() {
	// see mousecannon.go
	// engine.Register(&MouseCannonBullets{})
}

////////////////////////////////////////////////////////////////////////////////
// CALLBACKS ///////////////////////////////////////////////////////////////////
///////////////

func (mcb *MouseCannonBullets) Update(dt float64) {
	for i := range mcb.bullets {
		if !mcb.bullets[i].active {
			continue
		}

		mcb.bullets[i].livedtime += dt
		if mcb.bullets[i].livedtime > mcb.bullets[i].lifetime {
			mcb.bullets[i].active = false
			continue
		}

		// if mcb.bullets[i].pos.X() < 0 {
		// 	mcb.bullets[i].pos = mgl64.Vec2{float64(engine.Width), mcb.bullets[i].pos.Y()}
		// }
		// if mcb.bullets[i].pos.Y() < 0 {
		// 	mcb.bullets[i].pos = mgl64.Vec2{mcb.bullets[i].pos.X(), float64(engine.Height)}
		// }
		// if mcb.bullets[i].pos.X() > float64(engine.Width) {
		// 	mcb.bullets[i].pos = mgl64.Vec2{0.0, mcb.bullets[i].pos.Y()}
		// }
		// if mcb.bullets[i].pos.Y() > float64(engine.Height) {
		// 	mcb.bullets[i].pos = mgl64.Vec2{mcb.bullets[i].pos.X(), 0.0}
		// }

		mcb.bullets[i].pos = mcb.bullets[i].pos.Add(mcb.bullets[i].vel.Mul(dt))
	}
}

func (mcb *MouseCannonBullets) Draw(dt float64) {
	for i := range mcb.bullets {
		if !mcb.bullets[i].active {
			continue
		}

		primitives.DrawFilledCircle(
			primitives.Point{World.X(mcb.bullets[i].pos.X()), World.Y(mcb.bullets[i].pos.Y())},
			3.0,
			allegro.MapRGB(255, 0, 0),
		)
	}
}

////////////////////////////////////////////////////////////////////////////////
// FUNCTIONS ///////////////////////////////////////////////////////////////////
///////////////

func (mcb *MouseCannonBullets) fire(pos, vel mgl64.Vec2, lifetime float64) {
	for i := range mcb.bullets {
		if !mcb.bullets[i].active {
			mcb.bullets[i] = Bullet{
				active:    true,
				livedtime: 0,
				lifetime:  lifetime,

				pos: pos,
				vel: vel,
			}
			logrus.WithFields(logrus.Fields{
				"i":         i,
				"active":    mcb.bullets[i].active,
				"livedtime": mcb.bullets[i].livedtime,
				"lifetime":  mcb.bullets[i].lifetime,
				"pos":       mcb.bullets[i].pos,
				"vel":       mcb.bullets[i].vel,
			}).Info("set active")
			break
		}
	}
}

////////////////////////////////////////////////////////////////////////////////
// ENTS: BULLETS ///////////////////////////////////////////////////////////////
///////////////////

type Bullet struct {
	active    bool
	livedtime float64
	lifetime  float64

	pos mgl64.Vec2
	vel mgl64.Vec2
}
