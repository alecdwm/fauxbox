package game

import (
	"github.com/go-gl/mathgl/mgl64"
	"github.com/veandco/go-sdl2/sdl"
	"go.owls.io/fauxbox/engine"
)

////////////////////////////////////////////////////////////////////////////////
// ENT //
////////

type Bullet struct {
	pos mgl64.Vec2
	vel mgl64.Vec2

	color sdl.Color

	active    bool
	lifetime  float64
	livedtime float64

	posRect *sdl.Rect
}

func (b *Bullet) ID() uint {
	return 1
}

////////////////////////////////////////////////////////////////////////////////
// CALLBACKS //
//////////////

func (b *Bullet) Update(dt float64) {
	if !b.active {
		return
	}

	b.livedtime += dt
	if b.livedtime > b.lifetime {
		b.active = false
		return
	}

	if b.vel.X() != 0 || b.vel.Y() != 0 {
		b.pos = b.pos.Add(b.vel.Mul(dt))
	}

	b.posRect.X = int32(GameWorld.Pos(b.pos).X()) - b.posRect.W/2
	b.posRect.Y = int32(GameWorld.Pos(b.pos).Y()) - b.posRect.H/2
}

func (b *Bullet) Draw(dt float64) {
	if !b.active {
		return
	}

	engine.Renderer.SetDrawColor(b.color.R, b.color.G, b.color.B, b.color.A)
	engine.Renderer.DrawRect(b.posRect)
}

////////////////////////////////////////////////////////////////////////////////
// FUNCTIONS //
//////////////

func (b *Bullet) Fire(pos, vel mgl64.Vec2, color sdl.Color, lifetime float64) {
	if b.active {
		return
	}

	b.pos = pos
	b.vel = vel
	b.color = color
	b.active = true
	b.lifetime = lifetime
	b.livedtime = 0
}
